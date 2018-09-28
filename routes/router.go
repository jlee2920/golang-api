package router

import (
	"github.com/gorilla/mux"
	"github.com/jlee2920/golang-api.git/controller"
	"github.com/jlee2920/golang-api.git/models"
	"net/http"
)

type Server struct {
	Router *mux.Router
}

var People []models.Person

func (s *Server) InitializeRoutes() {
	r := mux.NewRouter()
	People = append(People, models.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &models.Address{City: "City X", State: "State X"}})
	People = append(People, models.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &models.Address{City: "City Z", State: "State Y"}})

	// People routes
	r.HandleFunc("/people", s.GetPeople).Methods("GET")
	r.HandleFunc("/people/{id}", s.GetPerson).Methods("GET")
	r.HandleFunc("/people/{id}", s.CreatePerson).Methods("POST")
	r.HandleFunc("/people/{id}", s.DeletePerson).Methods("DELETE")

	// Authorization
	r.HandleFunc("/authorize", s.Authorize).Methods("POST")

	//Token Creation
	r.HandleFunc("/tokens", s.CreateToken).Methods("POST")
	s.Router = r
}

// Display all from the people var
func (s *Server) GetPeople(w http.ResponseWriter, r *http.Request) {
	controller.GetPeople(w, r, People)
}

// Display a single data
func (s *Server) GetPerson(w http.ResponseWriter, r *http.Request) {
	controller.GetPerson(w, r, People)
}

// create a new item
func (s *Server) CreatePerson(w http.ResponseWriter, r *http.Request) {
	controller.CreatePerson(w, r, People)
}

// Delete an item
func (s *Server) DeletePerson(w http.ResponseWriter, r *http.Request) {
	controller.DeletePerson(w, r, People)
}

func (s *Server) Authorize(w http.ResponseWriter, r *http.Request) {
	controller.Authorize(w, r)
}

func (s *Server) CreateToken(w http.ResponseWriter, r *http.Request) {
	controller.CreateToken(w, r)
}

