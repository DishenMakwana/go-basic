package dbtools

import (
	"database/sql"
	"github.com/dishenmakwana/gRPC/communication/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DBInitializer struct {
	db *sql.DB
}

func Connect(driverName string, dataSourceName string) (*DBInitializer, error) {
	dbconncetion, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	return &DBInitializer{
		db: dbconncetion,
	}, nil
}

func (DBInitializer *DBInitializer) SelectStudentBasedID(id int64) (model.Student, error) {
	row := DBInitializer.db.QueryRow("select * from students where id=?", id)

	student := model.Student{}

	err := row.Scan(&student.ID, &student.Name, &student.Age)

	if err != nil {
		return student, err
	}

	return student, nil
}

func (DBInitializer *DBInitializer) SelectStudentsBasedName(name string) ([]model.Student, error) {
	rows, err := DBInitializer.db.Query("select * from students where name = ?", name)

	if err != nil {
		log.Fatal(err.Error())
	}

	students := []model.Student{}

	for rows.Next() {
		student := model.Student{}

		err = rows.Scan(&student.ID, &student.Name, &student.Age)

		if err != nil {
			log.Fatal(err.Error())
			continue
		}

		students = append(students, student)
	}

	return students, nil
}
