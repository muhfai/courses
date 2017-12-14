package config

import (
	"fmt"
	"log"
	"courses/models"
	"courses/mgo.v2"
	"courses/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// SERVER the DB server
const SERVER = "localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "Course"

// DOCNAME the name of the document
const DOCNAME = "courses"

// GetCourses returns the list of Courses
func (r Repository) GetCourses() models.Courses {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to connect into MongoDB Server", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := models.Courses{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// AddCourse inserts a Course in the DB
func (r Repository) AddCourse(course models.Course) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	course.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(course)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UpdateCourse updates an Course in the DB
func (r Repository) UpdateCourse(course models.Course) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	course.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).UpdateId(course.ID_Course, course)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// DeleteCourse deletes an Course
func (r Repository) DeleteCourse(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return "404"
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "500"
	}

	// Write status
	return "200"
}
