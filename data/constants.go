package data

// MaxSpeed understand how package works, public variable
const MaxSpeed = 60 //using const keyword, title case
const Mpi = 3.1452  //using const keyword, title case
// private variable
var minSpeed = 40 //using var keyword, camel case
var maxSpeed = 100

// private function
// getMinSpeed : get min speed
func getMinSpeed() int {
	return minSpeed
}

// getMaxSpeed : get max speed
func getMaxSpeed() int {
	return maxSpeed
}
