package controllers

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
	"strconv"
	"strings"
	//"time"
)

type App struct {
	*revel.Controller
}

var db *sql.DB

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

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) GetRaces(dt, city, name string) revel.Result {
	revel.INFO.Println("dt", dt)
	revel.INFO.Println("city", city)
	revel.INFO.Println("name", name)

	var racesStruct Races

	err := openDB()
	defer closeDB()
	if err != nil {
		revel.ERROR.Println(err)
	}
	revel.INFO.Println("openDB OK")

	query := `SELECT "Race"."race_UID", "Race".date, "Race".name, "Race".start_type, "Race".gps, "RefCitys".name As city
				FROM "Race","RefCitys"
				WHERE ("Race".date = '` + dt + `' AND "Race".city = "RefCitys"."city_ID")`
	rows, err := db.Query("SELECT array_to_json(ARRAY_AGG(row_to_json(row))) FROM (" + query + ") row")
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
	json.Unmarshal([]byte(data), &racesStruct.RacesArr)

	for i, value := range racesStruct.RacesArr {
		raceUID := value.UID

		query := `SELECT m_number, "race_UID", name, phone, gps, "marshal_ID"
  					FROM "Marshals"
			  		WHERE ("race_UID" = '` + raceUID + `')`
		rows, err := db.Query("SELECT array_to_json(ARRAY_AGG(row_to_json(row))) FROM (" + query + ") row")
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
		json.Unmarshal([]byte(data), &racesStruct.RacesArr[i].MarshalsArr)

		// classes
		query = `SELECT "class_UID", "race_UID", name, laps, date_start, date_finish
  					FROM "RaceClasses"
			  		WHERE ("race_UID" = '` + raceUID + `')`
		rows, err = db.Query("SELECT array_to_json(ARRAY_AGG(row_to_json(row))) FROM (" + query + ") row")
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
		json.Unmarshal([]byte(data), &racesStruct.RacesArr[i].ClassesArr)

		for n, value := range racesStruct.RacesArr[i].ClassesArr {
			classUID := value.UID

			query = `SELECT "checkpoint_ID", "race_UID", "class_UID", "number", gps, m_number
						FROM "Checkpoints"
						WHERE ("race_UID" = '` + raceUID + `' AND "class_UID" = '` + classUID + `')
						ORDER BY m_number`
			rows, err = db.Query("SELECT array_to_json(ARRAY_AGG(row_to_json(row))) FROM (" + query + ") row")
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
			json.Unmarshal([]byte(data), &racesStruct.RacesArr[i].ClassesArr[n].CheckpointsArr_todo)
			for _, value := range racesStruct.RacesArr[i].ClassesArr[n].CheckpointsArr_todo {
				chekpointNum := value.Number
				revel.WARN.Println("chekpointNum", chekpointNum)
				racesStruct.RacesArr[i].ClassesArr[n].Checkpoints = racesStruct.RacesArr[i].ClassesArr[n].Checkpoints + strconv.Itoa(chekpointNum) + ","
			}
			str := racesStruct.RacesArr[i].ClassesArr[n].Checkpoints
			if len(str) > 0 {
				if strings.HasSuffix(str, ",") {
					str = str[:len(str)-len(",")]
					racesStruct.RacesArr[i].ClassesArr[n].Checkpoints = str
				}
			}
		}
	}

	res, _ := json.Marshal(racesStruct)
	return c.RenderJson(string(res[:]))
}

func openDB() (err error) {
	driver := "postgres"
	connectString := "postgres" + "://"
	connectString += "postgres" + ":"
	connectString += "vtenduro" + "@"
	connectString += "localhost" + ":"
	connectString += "5432" + "/"
	connectString += "postgres"
	connectString += "?sslmode=" + "disable"
	db, err = sql.Open(driver, connectString)
	if err != nil {
		revel.ERROR.Println("DB open Error", err)
		return err
	}
	return nil
}

// closeDB ...
func closeDB() (err error) {
	err = db.Close()
	if err != nil {
		revel.ERROR.Println("DB close Error", err)
		return err
	}
	revel.INFO.Println("closeDB OK")
	return nil
}
