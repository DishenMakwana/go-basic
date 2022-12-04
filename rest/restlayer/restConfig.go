package restlayer

import (
	"github.com/gorilla/mux"
	"net/http"
)

func restConfig(router *mux.Router) {
	restRouter := router.PathPrefix("/rest/api").Subrouter()

	restRouter.Methods("GET").Path("/persons").HandlerFunc(SelectAllPersons)
	restRouter.Methods("GET").Path("/person/{name}").HandlerFunc(SelectPersonBasedName)
	restRouter.Methods("POST").Path("/person/add").HandlerFunc(SavePerson)
	restRouter.Methods("POST").Path("/person/edit").HandlerFunc(UpdatePerson)
}

func RestStart(endPoint string) error {
	router := mux.NewRouter()

	restConfig(router)

	return http.ListenAndServe(endPoint, router)
}
