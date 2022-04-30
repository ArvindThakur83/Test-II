package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// func CreateEmployee(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json") //json formate hai
// 	var emp Employees
// 	json.NewDecoder(r.Body).Decode(&emp)
// 	if emp.Position == "Hr" {
// 		fmt.Println("Permission Granted!")
// 		database.Create(&emp)
// 		json.NewEncoder(w).Encode(emp)
// 		fmt.Println(emp.EmpName)
// 	} else {
// 		panic("permission Dennied")
// 	}
// 	fmt.Println("Successfully Created")

// }
// func GetEmployees(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")
// 	var employees []Employees
// 	var emp Employees
// 	json.NewDecoder(r.Body).Decode(&emp)
// 	database.Find(&employees)
// 	if emp.Position == "Hr" || emp.Position == "Admin" {
// 		// database.Find(&employees)
// 		json.NewEncoder(w).Encode(employees)

// 	} else {
// 		fmt.Println("Sorry Acces Dennied!")
// 	}
// 	// database.Find(&employees)            //Find sare record return krega database se
// 	// json.NewEncoder(w).Encode(employees) //or yha user ko return kr dege json format me

// }

// func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("cotent-type", "application/json")
// 	var employee Employees
// 	//start

// 	database.First(&employee, mux.Vars(r)["eid"]) //mux.Vars(r) ek array return krega or apn id ke through value utha rhe hai
// 	json.NewEncoder(w).Encode(employee)
// }

// func UpdateEmployees(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")
// 	var employee Employees
// 	if employee.Position == "Hr" {
// 		fmt.Println("Permission Granted")
// 		database.First(&employee, mux.Vars(r)["eid"]) //database se daa utha ra hai id ke throught
// 		json.NewDecoder(r.Body).Decode(&employee)
// 		database.Save(&employee) //yha pr database me new value save kr ra hai
// 		json.NewEncoder(w).Encode(employee)
// 	} else {
// 		fmt.Println("PErmission Dennied!")
// 	}
// 	// database.First(&employee, mux.Vars(r)["eid"]) //database se daa utha ra hai id ke throught
// 	// json.NewDecoder(r.Body).Decode(&employee)     //yha pr jo v r me data aya hai uss se decode krke new value ko assign r ra hai .Decode(&employee)

// 	// database.Save(&employee)            //yha pr database me new value save kr ra hai
// 	// json.NewEncoder(w).Encode(employee) //encode krke return kr ra hai json ke format me
// }

// func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")
// 	var emp Employees
// 	if emp.Position == "Admin" {
// 		fmt.Println("Permission Granted!")
// 		database.Delete(&emp, mux.Vars(r)["eid"]) //yha fatch krke delete kr diya Delete method ki through
// 		json.NewEncoder(w).Encode("Employee is Deleted !!")
// 	}
// 	// database.Delete(&emp, mux.Vars(r)["eid"])           //yha fatch krke delete kr diya Delete method ki through
// 	// json.NewEncoder(w).Encode("Employee is Deleted !!") //user ko return krdega encode hokr
// }

//Weapons
func CreatWeapon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var wpn Weapons
	json.NewDecoder(r.Body).Decode(&wpn)
	database.Create(&wpn)
	json.NewEncoder(w).Encode(wpn)
}

func GetWeapons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var wpn []Weapons
	database.Find(&wpn)
	json.NewEncoder(w).Encode(wpn)
}

func GetWeaponsById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var wpm Weapons
	database.First(&wpm, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode(wpm)
}
