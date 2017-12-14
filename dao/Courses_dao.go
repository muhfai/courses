// Created by Muhammad Faisal Aziz (18215044) after learning gorilla-mux and MongoDB for almost 2 weeks
// II3160 - Pemrograman Integratif 2017
package dao

import (
	"log"
	"courses/models"
	mgo "gopkg.in/mgo.v2"// importing packages
	"gopkg.in/mgo.v2/bson"
)

// Start : Establish connection between Server and Database

type CoursesDAO struct {
	Server   string
	Database string
}

// For MongoDB Database
var db *mgo.Database

// Collect courses
const (
	COLLECTION = "courses"
)

// Establish connections to MongoDB Database
func (c *CoursesDAO) Connect() {
	session, err := mgo.Dial(c.Server)
	if err != nil {
		log.Fatal(err)
	}
	//create session
	db = session.DB(c.Database)
}

// Query-ing database in service(i'm sorry, i didn't know how to say it :( )
// Uses method db.C(COLLECTION) in order to implement

// Listing all of courses available in database
func (c *CoursesDAO) FindAll()([]models.Course, error){
	var courses []models.Course
	err := db.C(COLLECTION).Find(bson.M{}).All(&courses)
	return courses,err
}

// Finding a course based on its ID
func (c *CoursesDAO) FindById(ID_Course string) (models.Course, error) {
	var course models.Course
	err := db.C(COLLECTION).Find(bson.M{"ID_Course": ID_Course}).One(&course)
	return course, err
}

// Below are CRUD of Course DB

// Insert course into database
func (c *CoursesDAO) Insert(course models.Course) error {
	err := db.C(COLLECTION).Insert(&course)
	return err
}

// Delete an existing course
func (c *CoursesDAO) Delete(course models.Course) error {
	err := db.C(COLLECTION).Remove(&course)
	return err
}

// Update an existing course
func (c *CoursesDAO) Update(course models.Course) error {
	err := db.C(COLLECTION).Update(bson.M{"ID_Course": course.ID_Course}, &course)
	return err
}