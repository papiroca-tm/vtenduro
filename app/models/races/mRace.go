package races

import (
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/lib/pq" //
	"github.com/revel/revel"
	"strconv"
	"strings"
	//"time"
)

/*
MRace -
*/
type MRace struct {
	Races       Races
	RaceSimple  RaceSimple
	RacesSimple RacesSimple
	Race        Race
	DB          *sql.DB
}

/*
openDB - методт подключения к БД
*/
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

/*
closeDB - методт отключения от БД
*/
func (m *MRace) closeDB() (err error) {
	err = m.DB.Close()
	if err != nil {
		revel.ERROR.Println("DB close Error", err)
		return err
	}
	return nil
}

/*
GetRaceList - методот возвращает массив гонок с базовыми характеристиками
	dt - дата в формате 2006-01-02
	city - имя города, ищет по принципу LIKE %city% (опционально)
	name - имя гонки, ищет по принципу LIKE %name% (опционально)
*/
func (m *MRace) GetRaceList(dt, city, name string) (result RacesSimple, err error) {
	err = m.openDB()
	defer m.closeDB()
	if err != nil {
		revel.ERROR.Println(err)
		return result, err
	}
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
	rows, err := m.DB.Query("SELECT array_to_json(ARRAY_AGG(row_to_json(row))) FROM (" + query + ") row")
	defer rows.Close()
	if err != nil {
		revel.ERROR.Println(err)
		return result, err
	}
	var data string
	var row sql.NullString
	for rows.Next() {
		err = rows.Scan(&row)
		if err != nil {
			revel.ERROR.Println(err)
			return result, err
		}
	}
	data = row.String
	json.Unmarshal([]byte(data), &m.RacesSimple.RacesArr)
	result = m.RacesSimple
	return result, nil
}

/*
GetRace - метод возвращает структуру гонки заполненную основными свойствами
	raceUID - уникальный индификатор гонки
*/
func (m *MRace) GetRace(raceUID string) (result Race, err error) {
	err = m.openDB()
	defer m.closeDB()
	if err != nil {
		revel.ERROR.Println(err)
		return result, err
	}
	raceStruct := m.Race
	query := `SELECT "Race"."race_UID", "Race".date, "Race".name, "Race".start_type, "Race".gps, "Race".city As cityid, "RefCitys"."city_ID", "RefCitys".name AS city
				FROM "Race", "RefCitys"
				WHERE ("Race"."race_UID" = '` + raceUID + `' AND "Race".city = "RefCitys"."city_ID")`

	rows, err := m.DB.Query("SELECT row_to_json(row) FROM (" + query + ") row")
	defer rows.Close()
	if err != nil {
		revel.ERROR.Println(err)
		return result, err
	}

	var data string
	var row sql.NullString
	for rows.Next() {
		err = rows.Scan(&row)
		if err != nil {
			revel.ERROR.Println(err)
			return result, err
		}
	}
	data = row.String
	json.Unmarshal([]byte(data), &raceStruct)
	return raceStruct, err
}

/*
GetRaceMarshalsArr - метод возвращает массив маршалов гонки
	raceUID - уникальный индификатор гонки
*/
func (m *MRace) GetRaceMarshalsArr(raceUID string) (result []Marshal, err error) {
	err = m.openDB()
	defer m.closeDB()
	if err != nil {
		revel.ERROR.Println(err)
		return result, err
	}
	query := `SELECT m_number, "race_UID", name, phone, gps, "marshal_ID"
				FROM "Marshals"
				WHERE ("race_UID" = '` + raceUID + `')`
	rows, err := m.DB.Query("SELECT array_to_json(ARRAY_AGG(row_to_json(row))) FROM (" + query + ") row")
	defer rows.Close()
	if err != nil {
		revel.ERROR.Println(err)
		return result, err
	}
	var data string
	var row sql.NullString
	for rows.Next() {
		err = rows.Scan(&row)
		if err != nil {
			revel.ERROR.Println(err)
			return result, err
		}
	}
	data = row.String
	json.Unmarshal([]byte(data), &result)
	return result, err
}

/*
GetRaceMarshalInfo - метод возвращает данные по маршалу гонки
	raceUID - уникальный индификатор гонки
	mNumber - номер маршала
*/
func (m *MRace) GetRaceMarshalInfo(raceUID string, mNumber int) (result Marshal, err error) {
	err = m.openDB()
	defer m.closeDB()
	if err != nil {
		revel.ERROR.Println(err)
		return result, err
	}
	query := `SELECT m_number, "race_UID", name, phone, gps, "marshal_ID"
				FROM "Marshals"
				WHERE ("race_UID" = '` + raceUID + `' AND m_number ='` + strconv.Itoa(mNumber) + `')`
	rows, err := m.DB.Query("SELECT row_to_json(row) FROM (" + query + ") row")
	defer rows.Close()
	if err != nil {
		revel.ERROR.Println(err)
		return result, err
	}
	var data string
	var row sql.NullString
	for rows.Next() {
		err = rows.Scan(&row)
		if err != nil {
			revel.ERROR.Println(err)
			return result, err
		}
	}
	data = row.String
	json.Unmarshal([]byte(data), &result)
	return result, err
}

/*
GetRaceClassInfo - метод возвращает данные по классу гонки
	raceUID - уникальный индификатор гонки
	classUID - уникальный индификатор класса гонки
*/
func (m *MRace) GetRaceClassInfo(raceUID, classUID string) (result RaceClass, err error) {
	err = m.openDB()
	defer m.closeDB()
	if err != nil {
		revel.ERROR.Println(err)
		return result, err
	}
	query := `SELECT "class_UID", "race_UID", name, laps, date_start, date_finish
				FROM "RaceClasses"
				WHERE ("race_UID" = '` + raceUID + `' AND "class_UID" ='` + classUID + `')`
	rows, err := m.DB.Query("SELECT row_to_json(row) FROM (" + query + ") row")
	defer rows.Close()
	if err != nil {
		revel.ERROR.Println(err)
		return result, err
	}
	var data string
	var row sql.NullString
	for rows.Next() {
		err = rows.Scan(&row)
		if err != nil {
			revel.ERROR.Println(err)
			return result, err
		}
	}
	data = row.String
	json.Unmarshal([]byte(data), &result)
	return result, err
}

/*
GetRaceClassesArr - метод возвращает массив классов гонки
	raceUID - уникальный индификатор гонки
*/
func (m *MRace) GetRaceClassesArr(raceUID string) (result []RaceClass, err error) {
	err = m.openDB()
	defer m.closeDB()
	if err != nil {
		revel.ERROR.Println(err)
		return result, err
	}
	query := `SELECT "class_UID", "race_UID", name, laps, date_start, date_finish
				FROM "RaceClasses"
				WHERE ("race_UID" = '` + raceUID + `')`
	rows, err := m.DB.Query("SELECT array_to_json(ARRAY_AGG(row_to_json(row))) FROM (" + query + ") row")
	defer rows.Close()
	if err != nil {
		revel.ERROR.Println(err)
		return result, err
	}
	var data string
	var row sql.NullString
	for rows.Next() {
		err = rows.Scan(&row)
		if err != nil {
			revel.ERROR.Println(err)
			return result, err
		}
	}
	data = row.String
	//revel.WARN.Println("data", data)
	json.Unmarshal([]byte(data), &result)
	return result, err
}

/*
GetRaceCheckpointsArr - метод заполняет массив классов гонки массивом чекпоинтов
	и todo строкой чекпоинтов через запятую
	raceUID - уникальный индификатор гонки
	raceStruct - объект гонки
*/
func (m *MRace) GetRaceCheckpointsArr(raceUID string, raceStruct *Race) (err error) {
	err = m.openDB()
	defer m.closeDB()
	if err != nil {
		revel.ERROR.Println(err)
		return err
	}
	for n, value := range raceStruct.ClassesArr {
		classUID := value.UID
		query := `SELECT "checkpoint_ID", "race_UID", "class_UID", "number", gps, m_number
					FROM "Checkpoints"
					WHERE ("race_UID" = '` + raceUID + `' AND "class_UID" = '` + classUID + `')
					ORDER BY m_number`
		rows, err := m.DB.Query("SELECT array_to_json(ARRAY_AGG(row_to_json(row))) FROM (" + query + ") row")
		defer rows.Close()
		if err != nil {
			revel.ERROR.Println(err)
			return err
		}
		var data string
		var row sql.NullString
		for rows.Next() {
			err = rows.Scan(&row)
			if err != nil {
				revel.ERROR.Println(err)
				return err
			}
		}
		data = row.String
		json.Unmarshal([]byte(data), &raceStruct.ClassesArr[n].CheckpointsArr_todo)

		// todo бред, надо остовлять только массив
		for _, value := range raceStruct.ClassesArr[n].CheckpointsArr_todo {
			chekpointNum := value.Number
			raceStruct.ClassesArr[n].Checkpoints = raceStruct.ClassesArr[n].Checkpoints + strconv.Itoa(chekpointNum) + ","
		}
		str := raceStruct.ClassesArr[n].Checkpoints
		if len(str) > 0 {
			if strings.HasSuffix(str, ",") {
				str = str[:len(str)-len(",")]
				raceStruct.ClassesArr[n].Checkpoints = str
			}
		}

	}
	return err
}

/*
GetRaceInfo - метод возвращает полные данные по гонке
	raceUID - уникальный индификатор гонки
*/
func (m *MRace) GetRaceInfo(raceUID string) (result Race, err error) {
	if raceUID == "" {
		err := errors.New("raceUID пустой")
		return result, err
	}
	raceStruct, err := m.GetRace(raceUID)
	raceStruct.MarshalsArr, err = m.GetRaceMarshalsArr(raceUID)
	raceStruct.ClassesArr, err = m.GetRaceClassesArr(raceUID)
	err = m.GetRaceCheckpointsArr(raceUID, &raceStruct)
	if err != nil {
		return result, err
	}
	result = raceStruct
	return result, nil
}

/*
GetCheckpointsArr - метод возвращает массив чекпоинтов по классу гонки
	raceUID - уникальный индификатор гонки
	classUID - уникальный индификатор класса гонки
*/
func (m *MRace) GetCheckpointsArr(raceUID, classUID string) (result []Checkpoint, err error) {
	err = m.openDB()
	defer m.closeDB()
	if err != nil {
		revel.ERROR.Println(err)
		return result, err
	}
	query := `SELECT "checkpoint_ID", "race_UID", "class_UID", "number", gps, m_number
				FROM "Checkpoints"
				WHERE ("race_UID" = '` + raceUID + `' AND "class_UID" = '` + classUID + `')
				ORDER BY m_number`
	rows, err := m.DB.Query("SELECT array_to_json(ARRAY_AGG(row_to_json(row))) FROM (" + query + ") row")
	defer rows.Close()
	if err != nil {
		revel.ERROR.Println(err)
		return result, err
	}
	var data string
	var row sql.NullString
	for rows.Next() {
		err = rows.Scan(&row)
		if err != nil {
			revel.ERROR.Println(err)
			return result, err
		}
	}
	data = row.String
	json.Unmarshal([]byte(data), &result)
	return result, err
}

/*
GetCheckpointInfo - метод возвращает данные по контрольной точке класса гонки
	raceUID - уникальный индификатор гонки
	classUID - уникальный индификатор класса гонки
	number - номер контрольной точки
*/
func (m *MRace) GetCheckpointInfo(raceUID, classUID string, number int) (result Checkpoint, err error) {
	err = m.openDB()
	defer m.closeDB()
	if err != nil {
		revel.ERROR.Println(err)
		return result, err
	}
	query := `SELECT "checkpoint_ID", "race_UID", "class_UID", "number", gps, m_number
				FROM "Checkpoints"
				WHERE ("race_UID" = '` + raceUID + `' 
						AND "class_UID" = '` + classUID + `'
						AND m_number = ` + strconv.Itoa(number) + `)`
	rows, err := m.DB.Query("SELECT row_to_json(row) FROM (" + query + ") row")
	defer rows.Close()
	if err != nil {
		revel.ERROR.Println(err)
		return result, err
	}
	var data string
	var row sql.NullString
	for rows.Next() {
		err = rows.Scan(&row)
		if err != nil {
			revel.ERROR.Println(err)
			return result, err
		}
	}
	data = row.String
	json.Unmarshal([]byte(data), &result)
	return result, err
}
