# REST API под проект эндуро гонок и не только

##todo list
- список чекпоинтов через запятую в checkpoints, массив перенести в checkpointsTodo
- метод /getRaceList(dt,city,name)
- метод /getRaceInfo(raceUID)
- метод /getClassList(raceUID)
- метод /getClassInfo(raceUID,classUID)
- метод /getCheckpointList(raceUID)
- метод /getCheckpointInfo(raceUID,classUID,number)
- метод /getMarshalList(raceUID)
- метод /getMarshalInfo(raceUID,m_number)
- доработка запросов на получение списка гонок с учетом параметров name и city
- доработка реляционной модели с учетом замечаний по чекпоинтам, схема после доработок
- создание моделей и перенос кода работы с БД из контроллера в модель
- разработка реляционной модели оставшихся таблиц
- методы записи данных в БД по существующему функционалу get методов
- методы изменения данных в БД по существующему функционалу get методов
- методы удаления данных в БД по существующему функционалу get методов
- Переходим к участникам, результатам и тд.... 
