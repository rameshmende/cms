package main

import (
	"cms/model"
	"encoding/json"
	"net/http"
)

func (a *App) initStateRoutes() {
	a.Router.HandleFunc("/api/v1/states", a.GetStates).Methods("GET")
	a.Router.HandleFunc("/api/v1/states/{id}", a.GetState).Methods("GET")
	a.Router.HandleFunc("/api/v1/states", a.CreateStates).Methods("POST")
	a.Router.HandleFunc("/api/v1/states/{id}", a.UpdateStates).Methods("PUT")
	a.Router.HandleFunc("/api/v1/states/{id}", a.CreateStates).Methods("DELETE")
}

func (a *App) GetStates(res http.ResponseWriter, req *http.Request) {

}
func (a *App) GetState(res http.ResponseWriter, req *http.Request) {

}
func (a *App) CreateStates(res http.ResponseWriter, req *http.Request) {
	var response model.Response
	var state model.State
	res.Header().Set("Content-Type", "Application/json")
	err := json.NewDecoder(req.Body).Decode(&state)
	if err != nil {
		response.Status = http.StatusUnprocessableEntity
		response.Error = true
		response.Data = err.Error()
		json.NewEncoder(res).Encode(response)
	}

	err = state.Create(a.DB)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Error = true
		response.Data = err.Error()
	} else {
		response.Status = http.StatusOK
		response.Error = false
		response.Data = state
	}
	json.NewEncoder(res).Encode(response)

}
func (a *App) UpdateStates(res http.ResponseWriter, req *http.Request) {

}
func (a *App) DeleteStates(res http.ResponseWriter, req *http.Request) {

}
