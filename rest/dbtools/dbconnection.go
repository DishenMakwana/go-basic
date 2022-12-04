package dbtools

import (
	"github.com/dishenmakwana/rest/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

import "database/sql"

var driverName string
var dataSourceName string

func dbInitialize(dn, dsn string) {
	driverName = dn
	dataSourceName = dsn
}

func connect() (db *sql.DB) {

	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
func SelectAllPersons() []model.Person {
	db := connect()

	rows, err := db.Query("select * from person")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	persons := []model.Person{}

	for rows.Next() {
		person := model.Person{}

		err = rows.Scan(&person.ID, &person.Name, &person.Age)

		if err != nil {
			log.Fatal(err.Error())
			continue
		}

		persons = append(persons, person)
	}

	return persons
}

func SelectPersonBasedName(name string) model.Person {
	db := connect()

	rows := db.QueryRow("select * from person where name = ?", name)

	defer db.Close()

	person := model.Person{}

	err := rows.Scan(&person.ID, &person.Name, &person.Age)

	if err != nil {
		log.Fatal(err.Error())
	}

	return person
}

func Save(person model.Person) int64 {
	db := connect()

	defer db.Close()

	save, err := db.Prepare("insert into person (id,name,age) values (?,?,?)")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := save.Exec(person.ID, person.Name, person.Age)

	if err != nil {
		log.Fatal(err.Error())
	}

	personID, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err.Error())
	}

	return personID
}

func Update(person model.Person) int64 {

	db := connect()

	defer db.Close()

	update, err := db.Prepare("update person set name = ?, age=? where id=?")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := update.Exec(person.Name, person.Age, person.ID)

	if err != nil {
		log.Fatal(err.Error())
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
	}

	return rowsAffected
}

func Delete(id int) int64 {
	db := connect()
	defer db.Close()

	delete, err := db.Prepare("delete from person where id=?")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := delete.Exec()

	if err != nil {
		log.Fatal(err.Error())
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
	}

	return rowsAffected
}
