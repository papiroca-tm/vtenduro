package races

/*
Marshal - структура маршал
*/
type Marshal struct {
	Number int    `json:"m_number"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Gps    string `json:"gps"`
}

/*
Checkpoint - структура контрольной точки
*/
type Checkpoint struct {
	ID     int `json:"checkpoint_ID"`
	Number int `json:"number"`
	Gps    int `json:"gps"`
}

/*
RaceClass - структура классов гонки
*/
type RaceClass struct {
	UID                 string `json:"class_UID"`
	Name                string `json:"name"`
	Laps                int    `json:"laps"`
	DateTimeStart       string `json:"date_start"`
	DateTimeFinish      string `json:"date_finish"`
	CheckpointsArr_todo []Checkpoint
	Checkpoints         string
}

/*
Race - полная структура гонки
*/
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

/*
Races - структура с массивами полных структур гонок
*/
type Races struct {
	RacesArr []Race
}

/*
RaceSimple - упращенная структура гонки, без подчиненных массивов
*/
type RaceSimple struct {
	UID       string `json:"race_UID"`
	Date      string `json:"date"`
	Name      string `json:"name"`
	StartType int    `json:"start_type"`
	Gps       string `json:"gps"`
	City      string `json:"city"`
}

/*
RacesSimple - структура с массивами упращенных структур гонок
*/
type RacesSimple struct {
	RacesArr []RaceSimple
}
