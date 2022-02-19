package main

import "net/http"

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

}
func (a *App) UpdateStates(res http.ResponseWriter, req *http.Request) {

}
func (a *App) DeleteStates(res http.ResponseWriter, req *http.Request) {

}
