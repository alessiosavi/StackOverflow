package main

import "fmt"

var studentsCities = []string{"Mumbai", "Delhi", "Ahmedabad", "Mumbai", "Bangalore", "Delhi", "Kolkata", "Pune"}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	fmt.Printf("Cities before remove: %+v\n", studentsCities)
	for i := 0; i < len(studentsCities); i++ {
		if contains(studentsCities[i+1:], studentsCities[i]) {
			studentsCities = remove(studentsCities, i)
			i--
		}
	}
	fmt.Printf("Cities after remove: %+v\n", studentsCities)
}
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
