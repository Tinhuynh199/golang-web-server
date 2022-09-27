package main

//Using net/http to write rest api json, response results in html
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Student struct {
	Id int 
	Name string 
	Age int8
	Class []int
}

type Students []Student

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "This is my Home Page")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is my About Page")
}

func musicAPI(w http.ResponseWriter, r *http.Request) {
	var data = map[string]interface{} {
		"name": "Legend Never Die",
		"singer" : "Mr.T", 
	};
	json.NewEncoder(w).Encode(data)
}

func studentAPI(w http.ResponseWriter, r *http.Request) {
	var data = Student {1, "Huynh Xuan Tin", 20, []int {1,2,3}}
	json.NewEncoder(w).Encode(data)
}

func listStudentAPI(w http.ResponseWriter, r *http.Request) {
	var listStu = Students {
		{1, "Huynh Xuan Tin", 20, []int {1,2,3}},
		{2, "Dang minh triet", 21, []int {1,2,3}},
		{3, "trinh minh tien", 22, []int {1,2,3}},
	}
	json.NewEncoder(w).Encode(listStu)
}

func main() {
	fmt.Println("My Basic Website is running using net/http...")

	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/api/music", musicAPI)
	http.HandleFunc("/api/student", studentAPI)
	http.HandleFunc("/api/students", listStudentAPI)
	log.Fatal(http.ListenAndServe(":8080", nil))
}