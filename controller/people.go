package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jlee2920/golang-api.git/models"
	"net/http"
)

// Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request, people []models.Person) {
	json.NewEncoder(w).Encode(people)
}

// Display a single data
func GetPerson(w http.ResponseWriter, r *http.Request, people []models.Person) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Person{})
}

// create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request, people []models.Person) {
	params := mux.Vars(r)
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request, people []models.Person) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}
