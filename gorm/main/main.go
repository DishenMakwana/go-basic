package main

import (
	"encoding/json"
	"github.com/dishenmakwana/gorm/dbtools"
	"github.com/dishenmakwana/gorm/model"
	"log"
	"os"
)

type Configuration struct {
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

	dbtools.DBInitializer(conf.DataSourceName)

	// create database schema(table)
	dbtools.CreateTable(&model.Student{})

	//student := model.Student{
	//	ID:   3,
	//	Name: "dishen",
	//	Age:  22,
	//}

	// create new student
	//dbtools.Save(student)

	//getAll data
	//var students []model.Student
	//dbtools.Select(students)

	//update
	//student := model.Student{
	//	ID: 1,
	//}
	//
	//data := map[string]interface{}{
	//	"Name": "Makwana",
	//	"Age":  40,
	//}
	//
	//row := dbtools.SingleUpdate(&student, data)
	//
	//fmt.Println("Row affected: ", row)

	//multiple update

	//name := "darji"
	//nameWhereClause := "Name='" + name + "'"
	//data := map[string]interface{}{
	//	"Name": "kyada",
	//	"Age":  30,
	//}
	//rows := dbtools.MultipleUpdate(model.Student{}, nameWhereClause, data)
	//fmt.Println("Row affected: ", rows)

	//age := 22
	//ageWhereClause := "Age=" + strconv.Itoa(age)
	//data1 := map[string]interface{}{
	//	"Name": "darji",
	//	"Age":  "30",
	//}
	//rowsAffected := dbtools.MultipleUpdate(model.Student{}, ageWhereClause, data1)
	//fmt.Println("Row affected: ", rowsAffected)

	//single delete
	//student := model.Student{
	//	ID: 3,
	//}
	//row := dbtools.SingleDelete(&student)
	//fmt.Println("row affected:", row)

	//multiple delete
	//name := "Makwana"
	//nameWhereClause := "Name='" + name + "'"
	//rows := dbtools.MultipleDelete(model.Student{}, nameWhereClause)
	//fmt.Println("rows affected: ", rows)
	//
	//age := 22
	//ageWhereClause := "Age=" + strconv.Itoa(age)
	//rows1 := dbtools.MultipleDelete(model.Student{}, ageWhereClause)
	//fmt.Println("rows affected: ", rows1)
}
