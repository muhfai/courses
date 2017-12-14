// Muhammad Faisal Aziz - 18215044
// Setting Up API
// Thanks to Udemy for helping in learning Gorilla-mux

package main
 
import (
	"log"
	"net/http"
	// don't forget to include our dearest framework : gorilla mux
	"courses/gorilla/mux"
	"courses/mgo.v2/bson"
	
	// importing all packages
	"courses/config"
	"courses/dao"
	"courses/models"
	// import json
	"encoding/json"
)

var conf = config.Config{}
var cobadao = dao.CoursesDAO{}


// POST a new course
// Decodes the request body into a course object, assign it an ID_Courses, and uses the DAO Insert method to create a course in database
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var course models.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	course.ID = bson.NewObjectId()
	if err := cobadao.Insert(course); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, course)
}

// GET list of all courses
// Uses FindAll method of DAO Library to fetch list of courses from database
func AllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := cobadao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, courses)
}

// CRUD

// GET a course by its ID_Course
// Using mux library to get parameters that the users passed in with the request
func FindCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	course, err := cobadao.FindById(params["ID_Course"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Course tidak ditemukan!")
		return
	}
	respondWithJson(w, http.StatusOK, course)
}

// PUT update an existing course
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var course models.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid")
		return
	}
	if err := cobadao.Update(course); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing course
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var course models.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid")
		return
	}
	if err := cobadao.Delete(course); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// Error handling
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

// JSON responses
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parsing connection into config.toml
// Connecting config into db
func init() {
	conf.Read()

	cobadao.Server = conf.Server
	cobadao.Database = conf.Database
	cobadao.Connect()
}

// Main
// Handling HTTP Routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/courses", AllCourses).Methods("GET")
	r.HandleFunc("/courses", CreateCourse).Methods("POST")
	r.HandleFunc("/courses", UpdateCourse).Methods("PUT")
	r.HandleFunc("/courses", DeleteCourse).Methods("DELETE")
	r.HandleFunc("/courses/{ID_Course}", FindCourse).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}