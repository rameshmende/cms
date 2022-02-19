package main

import (
	"cms/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *App) initDoctorRoutes() {
	a.Router.HandleFunc("/api/v1/doctors", a.GetDoctors).Methods("GET")
	a.Router.HandleFunc("/api/v1/doctors/{id}", a.GetDoctor).Methods("GET")
	a.Router.HandleFunc("/api/v1/doctors", a.CreateDoctors).Methods("POST")
	a.Router.HandleFunc("/api/v1/doctors/{id}", a.UpdateDoctors).Methods("PUT")
	a.Router.HandleFunc("/api/v1/doctors/{id}", a.CreateDoctors).Methods("DELETE")
}

func (a *App) GetDoctors(res http.ResponseWriter, req *http.Request) {

}
func (a *App) GetDoctor(res http.ResponseWriter, req *http.Request) {

}
func (a *App) CreateDoctors(res http.ResponseWriter, req *http.Request) {

}
func (a *App) UpdateDoctors(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	vars := mux.Vars(req)
	id := vars["id"]
	docId, err := strconv.Atoi(id)
	var updatedFields map[string]interface{}
	res.Header().Set("Content-Type", "Application/json")
	err = json.NewDecoder(req.Body).Decode(&updatedFields)
	fmt.Println(err)
	fmt.Println(updatedFields)
	if err != nil {
		response.Status = http.StatusUnprocessableEntity
		response.Error = true
		response.Data = err.Error()
		json.NewEncoder(res).Encode(response)
	}

	doctor := model.Doctor{ID: docId}
	err = doctor.Update(a.DB, updatedFields)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Error = true
		response.Data = err.Error()
	} else {
		response.Status = http.StatusOK
		response.Error = true
		response.Data = err.Error()
	}
	json.NewEncoder(res).Encode(response)
}
func (a *App) DeleteDoctors(res http.ResponseWriter, req *http.Request) {

}
