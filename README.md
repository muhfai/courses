# Courses
Repository for Final Assignment of II3160 - Pemrograman Integratif by Muhammad Faisal Aziz (STI ITB 2015 - 18215044)
Welcome to courses project, a webservice that enabled users to find courses (All Courses, or by ID), Create a new course, update a course, or delete a course. Hopefully, this service can be applied in e-learning website of ITB, or implemented in another e-learning platform.

# Functional Requirements
1. System can find a course by its ID using method Find
2. System can list All Course that enabled using method AllCourses
3. System can create a new course using method POST
4. System can update a course using method PUT
5. System can delete a course using method DELETE

# Non-Functional Requirements
1. Using Golang 1.9 for programming language
2. Using MongoDB as database
3. Minimum OS of Ubuntu 16.04
4. Storage of 10 GB
5. Using Gorilla/mux as framework
6. Uses binary-json (bson) to encode JSON format

# Installations
First of all, please install Golang 1.9
Kindly install Gorilla/mux and include library of mongodb, bson, toml (to parse and encode in GoLang)
```
$go get github.com/gorilla/mux
$go get gopkg.in/mgo.v2
$go get github.com/BurntSushi/toml
```
After that, clone this project into your $GOPATH
```
$ cd $GOPATH/src
$ git clone https://github.com/muhfai/courses.git
```
Run Project
```
$ go build
$ ./courses
```

# Access
http://ip:port/courses (to check all)
http://ip:port/courses/{ID_Courses} (to check a course based on its ID)

# POST, PUT, GET, and DELETE
Uses Advance REST Client application. You can download that. Or maybe you can use Postman app (https://www.getpostman.com/)