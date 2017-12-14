package models

import "courses/mgo.v2/bson"

// Represents a Course, we uses bson to tell the mgo driver
// bson = binary JSON, a bin­ary-en­coded seri­al­iz­a­tion of JSON-like doc­u­ments
type Course struct {
	ID		bson.ObjectId
	ID_Course         string `bson:"ID_Course" json:"ID_Course"`
	Course_Name        string        `bson:"Course_Name" json:"Course_Name"`
	Lecturer  string        `bson:"Lecturer" json:"Lecturer"`
    Topic  string        `bson:"Topic" json:"Topic"`
	Attendance int        `bson:"Attendance" json:"Attendance"`
}

//Course is an array of Course
type Courses []Course