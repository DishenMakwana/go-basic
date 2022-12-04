package dbtools

import (
	"database/sql"
	"github.com/dishenmakwana/database/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var driverName string
var dataSourceName string

func DBInitializer(dn, dsn string) {
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

func SelectAllStudents() []model.Student {
	db := connect()

	rows, err := db.Query("selec * from student")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

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

	return students
}

func SelectStudentBasedID(id int) model.Student {
	db := connect()

	row := db.QueryRow("select * from student where id = ?", id)

	defer db.Close()

	student := model.Student{}

	err := row.Scan(&student.ID, &student.Name, &student.Age)

	if err != nil {
		log.Fatal(err.Error())
	}

	return student
}
func Save(student model.Student) int64 {
	db := connect()

	defer db.Close()

	save, err := db.Prepare("insert into student(id,name,age) values(?,?,?)")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := save.Exec(student.ID, student.Name, student.Age)

	if err != nil {
		log.Fatal(err.Error())
	}

	studentID, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err.Error())
	}

	return studentID
}

func Update(student model.Student) int64 {
	db := connect()

	defer db.Close()

	update, err := db.Prepare("update student set name=?, age=? where id=?")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := update.Exec(student.Name, student.Age, student.ID)

	if err != nil {
		log.Fatal(err.Error())
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
	}

	return rowAffected
}

func Delete(id int) int64 {
	db := connect()

	defer db.Close()

	delete, err := db.Prepare("delete from student id=?")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := delete.Exec(id)

	if err != nil {
		log.Fatal(err.Error())
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
	}

	return rowsAffected
}
