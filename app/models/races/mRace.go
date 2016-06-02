package races

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
	"strconv"
	"strings"
	//"time"
)

type MRace struct {
	Races       Races
	RacesSimple RacesSimple
	Race        Race
	DB          *sql.DB
}

//var db *sql.DB

type Marshal struct {
	Number int    `json:"m_number"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Gps    string `json:"gps"`
}

type Checkpoint struct {
	ID     int `json:"checkpoint_ID"`
	Number int `json:"number"`
	Gps    int `json:"gps"`
}

type RaceClass struct {
	UID                 string `json:"class_UID"`
	Name                string `json:"name"`
	Laps                int    `json:"laps"`
	DateTimeStart       string `json:"date_start"`
	DateTimeFinish      string `json:"date_finish"`
	CheckpointsArr_todo []Checkpoint
	Checkpoints         string
}

type Race struct {
	UID         string `json:"race_UID"`
	Date        string `json:"date"`
	Name        string `json:"name"`
	StartType   int    `json:"start_type"`
	Gps         string `json:"gps"`
	City        string `json:"city"`
	ClassesArr  []RaceClass
	MarshalsArr []Marshal
}

type Races struct {
	RacesArr []Race
}

type RaceSimple struct {
	UID       string `json:"race_UID"`
	Date      string `json:"date"`
	Name      string `json:"name"`
	StartType int    `json:"start_type"`
	Gps       string `json:"gps"`
	City      string `json:"city"`
}

type RacesSimple struct {
	RacesArr []RaceSimple
}

func (m *MRace) GetRaceList(dt, city, name string) (res string) {
	err := m.openDB()
	defer m.closeDB()
	if err != nil {
		revel.ERROR.Println(err)
	}
	revel.INFO.Println("openDB OK")
	query := `SELECT "Race"."race_UID", "Race".date, "Race".name, "Race".start_type, "Race".gps,  "Race".city As cityid, "RefCitys"."city_ID", "RefCitys".name AS city
					FROM "Race","RefCitys"
					WHERE (
						"Race".date = '` + dt + `' 
						AND "Race".city = "RefCitys"."city_ID" `
	if city != "" {
		query = query + ` AND lower("RefCitys".name) LIKE '%` + strings.ToLower(city) + `%'`
	}
	if name != "" {
		query = query + ` AND lower("Race".name) LIKE '%` + strings.ToLower(name) + `%'`
	}
	query = query + `)`
	revel.WARN.Println("query ", query)

	rows, err := m.DB.Query("SELECT array_to_json(ARRAY_AGG(row_to_json(row))) FROM (" + query + ") row")
	defer rows.Close()
	if err != nil {
		revel.ERROR.Println(err)
	}

	var data string
	var row sql.NullString
	for rows.Next() {
		err = rows.Scan(&row)
		if err != nil {
			revel.ERROR.Println(err)
			break
		}
	}
	data = row.String
	json.Unmarshal([]byte(data), &m.RacesSimple.RacesArr)

	resByte, _ := json.Marshal(m.RacesSimple)
	res = string(resByte[:])
	return res
}

func (m *MRace) GetRaceInfo(raceUID string) (res string) {
	err := m.openDB()
	defer m.closeDB()
	if err != nil {
		revel.ERROR.Println(err)
	}
	revel.INFO.Println("openDB OK")
	query := `SELECT "Race"."race_UID", "Race".date, "Race".name, "Race".start_type, "Race".gps, "Race".city As cityid, "RefCitys"."city_ID", "RefCitys".name AS city
				FROM "Race", "RefCitys"
				WHERE ("Race"."race_UID" = '` + raceUID + `' AND "Race".city = "RefCitys"."city_ID")`

	rows, err := m.DB.Query("SELECT row_to_json(row) FROM (" + query + ") row")
	defer rows.Close()
	if err != nil {
		revel.ERROR.Println(err)
	}

	var data string
	var row sql.NullString
	for rows.Next() {
		err = rows.Scan(&row)
		if err != nil {
			revel.ERROR.Println(err)
			break
		}
	}
	data = row.String
	revel.WARN.Println("data", data)
	json.Unmarshal([]byte(data), &m.Race)

	data = ""
	query = `SELECT m_number, "race_UID", name, phone, gps, "marshal_ID"
				FROM "Marshals"
				WHERE ("race_UID" = '` + raceUID + `')`
	rows, err = m.DB.Query("SELECT array_to_json(ARRAY_AGG(row_to_json(row))) FROM (" + query + ") row")
	defer rows.Close()
	if err != nil {
		revel.ERROR.Println(err)
	}

	for rows.Next() {
		err = rows.Scan(&row)
		if err != nil {
			revel.ERROR.Println(err)
			break
		}
	}
	data = row.String
	revel.WARN.Println("data", data)
	json.Unmarshal([]byte(data), &m.Race.MarshalsArr)

	// classes
	data = ""
	query = `SELECT "class_UID", "race_UID", name, laps, date_start, date_finish
				FROM "RaceClasses"
				WHERE ("race_UID" = '` + raceUID + `')`
	rows, err = m.DB.Query("SELECT array_to_json(ARRAY_AGG(row_to_json(row))) FROM (" + query + ") row")
	defer rows.Close()
	if err != nil {
		revel.ERROR.Println(err)
	}

	for rows.Next() {
		err = rows.Scan(&row)
		if err != nil {
			revel.ERROR.Println(err)
			break
		}
	}
	data = row.String
	revel.WARN.Println("data", data)
	json.Unmarshal([]byte(data), &m.Race.ClassesArr)

	for n, value := range m.Race.ClassesArr {
		classUID := value.UID

		data = ""
		query = `SELECT "checkpoint_ID", "race_UID", "class_UID", "number", gps, m_number
					FROM "Checkpoints"
					WHERE ("race_UID" = '` + raceUID + `' AND "class_UID" = '` + classUID + `')
					ORDER BY m_number`
		rows, err = m.DB.Query("SELECT array_to_json(ARRAY_AGG(row_to_json(row))) FROM (" + query + ") row")
		defer rows.Close()
		if err != nil {
			revel.ERROR.Println(err)
		}

		for rows.Next() {
			err = rows.Scan(&row)
			if err != nil {
				revel.ERROR.Println(err)
				break
			}
		}
		data = row.String
		revel.WARN.Println("data", data)
		json.Unmarshal([]byte(data), &m.Race.ClassesArr[n].CheckpointsArr_todo)
		for _, value := range m.Race.ClassesArr[n].CheckpointsArr_todo {
			chekpointNum := value.Number
			m.Race.ClassesArr[n].Checkpoints = m.Race.ClassesArr[n].Checkpoints + strconv.Itoa(chekpointNum) + ","
		}
		str := m.Race.ClassesArr[n].Checkpoints
		if len(str) > 0 {
			if strings.HasSuffix(str, ",") {
				str = str[:len(str)-len(",")]
				m.Race.ClassesArr[n].Checkpoints = str
			}
		}
	}

	resByte, _ := json.Marshal(m.Race)
	res = string(resByte[:])
	return res
}

func (m *MRace) openDB() (err error) {
	driver := "postgres"
	connectString := "postgres" + "://"
	connectString += "postgres" + ":"
	connectString += "vtenduro" + "@"
	connectString += "localhost" + ":"
	connectString += "5432" + "/"
	connectString += "postgres"
	connectString += "?sslmode=" + "disable"
	m.DB, err = sql.Open(driver, connectString)
	if err != nil {
		revel.ERROR.Println("DB open Error", err)
		return err
	}
	return nil
}

// closeDB ...
func (m *MRace) closeDB() (err error) {
	err = m.DB.Close()
	if err != nil {
		revel.ERROR.Println("DB close Error", err)
		return err
	}
	revel.INFO.Println("closeDB OK")
	return nil
}
