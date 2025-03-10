# Валютчик

Данный сервис раз в день собирает данные о курсе валют разных стран по отношению к белорусскому рублю: ```https://api.nbrb.by/exrates/rates?periodicity=0```

При разработке сервиса использовались следующие пакеты:
1. [Gin](https://gin-gonic.com/) - минималистичный веб-фреймворк для Go
2. [Gorm](https://gorm.io/) - библиотека ORM для Go
3. [Gocron](https://github.com/go-co-op/gocron) - пакет планирования заданий, который позволяет выполнять функции с заранее определенными интервалами
4. [GoDotEnv](https://github.com/joho/godotenv) - порт проекта Ruby Dotenv для Go (загружает env переменные из файла .env)

## Запуск

Для запуска сервиса необходимо создать базу данных MySQL и сохранить строку подключения в файл ```.env``` (в корневом каталоге) в переменную ```DATABASE_URL```

Также можно воспользоваться строкой подключения по умолчанию

```root:@tcp(127.0.0.1:3306)/currency?charset=utf8mb4&parseTime=True&loc=Local```

* root - имя пользователя (без пароля)
* 127.0.0.1:3306 - локальный адрес MySQL сервера
* currency - имя базы данных, которую использует проект

В этом случае достаточно создать базу данных ```currency```, а после запустить проект

После запуска проекта выполняется автоматическая миграция, которая создаст таблицу валют в базе данных

## API

Данный проект имеет один роут, который возвращает список сохранённых данных о курсе валют:

```/data``` - список всех данных, сохранённых в БД

```/data?date=YYYY-MM-DD``` - список всех данных, сохранённых в БД за определённую дату