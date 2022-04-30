package main

import "fmt"

func main() {
	map_1 := map[string]string{

		"Admin":    "admin",
		"hr":       "hr",
		"employee": "employee",
	}
	fmt.Println(map_1)
	DataMigration()
	HandlerRouting()
}
