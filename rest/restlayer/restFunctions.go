package restlayer

import (
	"encoding/json"
	"fmt"
	"github.com/dishenmakwana/rest/dbtools"
	"github.com/dishenmakwana/rest/model"
	"github.com/gorilla/mux"
	"net/http"
)

func SelectPersonBasedName(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	name, ok := vars["name"]

	if !ok {
		response.WriteHeader(http.StatusBadRequest)

		fmt.Println(response, "Person not Found")
	}

	persons := dbtools.SelectPersonBasedName(name)

	json.NewEncoder(response).Encode(persons)
}

func SelectAllPersons(response http.ResponseWriter, request *http.Request) {
	persons := dbtools.SelectAllPersons()

	json.NewEncoder(response).Encode(persons)
}

func SavePerson(response http.ResponseWriter, request *http.Request) {
	var person model.Person

	err := json.NewDecoder(request.Body).Decode(&person)

	if err != nil {
		fmt.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Println(response, " could not get person :", err)
		return
	}

	dbtools.Save(person)
}

func UpdatePerson(response http.ResponseWriter, request *http.Request) {
	var person model.Person

	err := json.NewDecoder(request.Body).Decode(&person)

	if err != nil {
		fmt.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Println(response, " could not get person :", err)
		return
	}

	dbtools.Update(person)
}
