package main

import (
	"encoding/json"
	"fmt"
	"github.com/dishenmakwana/database/dbtools"
	"github.com/dishenmakwana/database/model"
	"log"
	"os"
)

type Configuration struct {
	DriverName     string `json:"driverName"`
	DataSourceName string `json:"dataSourceName"`
}

func main() {
	file, err := os.Open("configuration/config.json")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	conf := new(Configuration)

	json.NewDecoder(file).Decode(conf)

	dbtools.DBInitializer(conf.DriverName, conf.DataSourceName)

	students := dbtools.SelectAllStudents()

	for _, std := range students {
		fmt.Println("ID: ", std.ID, ", Name: ", std.Name, " Age: ", std.Age)
	}

	student1 := dbtools.SelectStudentBasedID(1)

	fmt.Println("ID: ", student1.ID, ", Name: ", student1.Name, " Age: ", student1.Age)

	student := model.Student{
		ID:   7,
		Name: "rajan",
		Age:  30,
	}

	lastInsertID := dbtools.Save(student)

	fmt.Println("last inserted ID: ", lastInsertID)

	rowsAffected := dbtools.Update(student)

	fmt.Println("RowsAffected: ", rowsAffected)

	rowsAffectedByDelete := dbtools.Delete(7)

	fmt.Println("RowsAffected by Delete:", rowsAffectedByDelete)
}
