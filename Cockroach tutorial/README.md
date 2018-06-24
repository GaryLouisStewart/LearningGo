# Example application using CockroachDB & golang & GORM to generate some dummy data
[tutorial link](https://www.cockroachlabs.com/docs/stable/build-a-go-app-with-cockroachdb-gorm.html)

## Required libraries

1. [gorm](github.com/jinzhu/gorm)
2. [gorm postgres dialect](github.com/jinzhu/gorm/dialects/postgres)

- Install cockroachdb and run cockroach db with the following command
```
cockroach start \
--insecure \
--store=example-app \
--host=localhost
```
- Create a user account

` cockroach user set dbadmin --insecure`

- Setup the database and grant priveleges

```
cockroach sql --insecure -e 'CREATE DATABASE bank'
cockroach sql --insecure -e 'GRANT ALL ON DATABASE bank to dbadmin'
```

- run the following file.`go run create-sample-data.go`


