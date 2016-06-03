#REST API под проект эндуро гонок и не только
##todo list
- убрать отдельное подключение к БД из каждого метода модели mRace -- (03.06.2016)
- создание документа с описанием рест методов в директории docs -- (02.06.2016)
- метод /api/getClassInfo(raceUID,classUID) -- 31.05.16
- метод /api/getCheckpointList(raceUID) -- 31.05.16
- метод /api/getCheckpointInfo(raceUID,classUID,number) -- 31.05.16

- метод /api/getMarshalInfo(raceUID,m_number) -- 31.05.16
- разработка реляционной модели оставшихся таблиц -- 31.05.16
- методы записи данных в БД по существующему функционалу get методов -- 31.05.16
- методы изменения данных в БД по существующему функционалу get методов -- 31.05.16
- методы удаления данных в БД по существующему функционалу get методов -- 31.05.16
- Переходим к участникам, результатам и тд.... 

##Функциональные возможности
|наименование метода|параметры|описание|
|---|---|---|
|/api/getRaceList|(dt,city,name)|методот возвращает массив гонок с базовыми характеристиками|
|/api/getRaceInfo|(raceUID)|метод возвращает полные данные по гонке|
|/api/getClassList|(raceUID)|метод возвращает массив классов гонки|
|/api/getMarshalList|(raceUID)|метод возвращает массив маршалов гонки|


##work done:
####03.06.2016
- рефакторинг модели mRace (разделить getRaceInfo на функции) -- (02.06.2016)
- метод /api/getClassList(raceUID) -- 31.05.16
- метод /api/getMarshalList(raceUID) -- 31.05.16


####02.06.2016
- убрать ненужные поля с массивами из гонки при getRaceList -- (01.06.2016)
- убрать массив гонок при getRaceInfo -- (01.06.2016)
- исправление ошибки с пустым городом в getRaceInfo -- (01.06.2016)
- доработка запросов на получение списка гонок с учетом параметров name и city -- 31.05.16
- независимый от консоли запуск(screen) -- (01.06.2016)


####01.06.2016
- список чекпоинтов через запятую в checkpoints, массив перенести в checkpointsTodo -- 31.05.16
- метод /api/getRaceList(dt,city,name) -- 31.05.16
- метод /api/getRaceInfo(raceUID) -- 31.05.16
- создание моделей и перенос кода работы с БД из контроллера в модель -- 31.05.16
- доработка реляционной модели с учетом замечаний по чекпоинтам, схема после доработок -- 31.05.16