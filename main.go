package main

import (
	"frontendmaster.com/go/io/data"
) //import data package

// global variable
var url string = "https://www.google.com"

func main() {
	//data.getMaxSpeed() error : private function

}

// isEven : check if number is even or not
func isEven(number int) bool {
	return number%2 == 0 //return true if number is even
}

// strLength : return length of string
func strLength(str string) int {
	return len(str)
}

// isValidIsbn : check if isbn is valid or not
func isValidIsbn(isbn string) bool {
	//vector for store isbn identifier
	var isbnIdentifier []string = []string{"978", "979"}
	//check if isbn identifier is valid
	return contains(isbnIdentifier, isbn[0:3])
}

// contains : check if array contains value or not
func contains(arr []string, value string) bool {
	for _, v := range arr { //loop through array
		if v == value {
			return true
		}
	}
	return false
}

// exceedMax : check if number exceed max or not
func exceedMax(val int) bool {
	const MAKS_VALUE = 999
	return val > MAKS_VALUE
}

// sucessSave : check if fail succes saved in the directory
func succesSave(file_name string) bool {
	//vector than contain file name in the curr directory
	var file_dir []string = []string{"main.go", "test_main.go"}
	for _, gofile := range file_dir {
		if gofile == file_name {
			return true
		}
	}
	return false
}

// binarySearch : performing binary search algorithm
func binarySearch(item string, low int, high int, searchItem byte) int {
	if high < low {
		return -1
	}
	mid := (low + high) / 2
	if item[mid] == searchItem {
		return mid
	} else if item[mid] < searchItem {
		return binarySearch(item, mid+1, high, searchItem)
	} else {
		return binarySearch(item, low, mid-1, searchItem)
	}
}

// printData : print data to console
func printData(data string) string {
	return data
}

// getData : get data from data package
func getData() int {
	return data.MaxSpeed
}

// circleArea : calculate circle area
func circleArea(radius float64) float64 {
	return data.Mpi * radius * radius
}

// priceToInt : convert price to int
func priceToInt(price float64) int {
	return int(price)
}

// initAge : init age of student
func initAge() [10]int {
	//init array
	data.StudentAge = [10]int{10, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	return data.StudentAge
}

// initName : init name of student
func initName() []string {
	//init slice
	data.StudentName = []string{"jack", "miller", "jane", "joe", "james", "jessica", "jennifer", "jessie", "julia", "julie"}
	return data.StudentName
}

// totalItemsName : get total items, for student name
func totalItemsName(items interface{}) int {
	return len(items.([]string))
}

// totalItemsAge : get total items, for student age
func totalItemsAge(items interface{}) int {
	return len(items.([10]int))
}

// initStatus : init status of student
func initStatus() map[string]bool {
	//init map
	data.StudentStatus = map[string]bool{
		"jack":        true,
		"miller":      false,
		"jane":        true,
		"joe":         false,
		"james":       true,
		"jessica":     false,
		"jennifer":    true,
		"jessie":      false,
		"julia":       true,
		"jane miller": false,
	}
	return data.StudentStatus
}

// appendPort : append port to well known ports
func appendPort(port string, value int) map[string]int {
	data.WellKnownPorts[port] = value
	return data.WellKnownPorts
}
