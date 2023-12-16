// testmain.go : unit testing for main.go
package main

//generate unit test for main.go
import (
	"frontendmaster.com/go/io/data"
	"testing"
)

// test for isEven function
func TestIsEven(t *testing.T) {
	//test case 1
	if !isEven(2) {
		t.Error("2 is even number")
	}
	//test case 2
	if isEven(3) {
		t.Error("3 is not even number")
	}
}

// test for strLength function
func TestStrLength(t *testing.T) {
	//test case 1
	if strLength("hello") != 5 {
		t.Error("Length of hello is 5")
	}
	//test case 2
	if strLength("world") != 5 {
		t.Error("Length of world is 5")
	}
	//test case for empty string
	if strLength("") != 0 {
		t.Error("Length of empty string is 0")
	}
}

// test for isValidIsbn function
func TestIsValidIsbn(t *testing.T) {
	//test case 1
	t.Run("978-0xmath-calculus is valid isbn", func(t *testing.T) {
		if !isValidIsbn("978-0xmath-calculus") {
			t.Error("978-0xmath-calculus is valid isbn")
		} else {
			t.Log("Passed : 978-0xmath-calculus is valid isbn")
		}
	})
	//test case 2
	t.Run("0x9-0-393-04002-9 is invalid isbn", func(t *testing.T) {
		if isValidIsbn("0x9-0-393-04002-9") {
			t.Error("0x9-0-393-04002-9 is invalid isbn")
		} else {
			t.Log("Passed : 0x9-0-393-04002-9 is valid isbn")
		}
	})
}

// test for contains function
func TestContains(t *testing.T) {
	//test case 1
	t.Run("contains([\"978\", \"979\"], \"978\") is true", func(t *testing.T) {
		if !contains([]string{"978", "979"}, "978") {
			t.Error("contains([\"978\", \"979\"], \"978\") is true")
		} else {
			t.Log("Passed : contains([\"978\", \"979\"], \"978\") is true")
		}
	})
	//test case 2
	t.Run("contains([\"978\", \"979\"], \"977\") is false", func(t *testing.T) {
		if contains([]string{"978", "979"}, "977") {
			t.Error("contains([\"978\", \"979\"], \"977\") is false")
		} else {
			t.Log("Passed : contains([\"978\", \"979\"], \"977\") is false")
		}
	})
}

// test for exceedMax function
func TestExceedMax(t *testing.T) {
	//test case 1
	t.Run("100 is not exceed the maximum", func(t *testing.T) {
		if exceedMax(100) {
			t.Error("100 is not exceed the maximum")
		} else {
			t.Log("Passed : 100 is not exceed the maximum")
		}
	})
	//test case 2
	t.Run("1000 is exceed the maximum", func(t *testing.T) {
		if !exceedMax(1000) {
			t.Error("1000 is exceed the maximum")
		} else {
			t.Log("Passed : 1000 is exceed the maximum")
		}
	})
}

// test for succesSave function
func TestSuccesSave(t *testing.T) {
	//test case 1
	var file_name = "main.go"
	t.Run("main.go is saved to directory", func(t *testing.T) {
		if !succesSave(file_name) {
			t.Error("main.go is saved to database")
		} else {
			t.Log("passed : main.go is saved to database")
		}
	})
	//test case 2
	file_name = "mailer.go"
	t.Run("mailer.go is not saved", func(t *testing.T) {
		if succesSave(file_name) {
			t.Error("mailer.go is not saved")
		} else {
			t.Log("mailer.go is not saved")
		}
	})
}

// test binary search function
func TestBinarySearch(t *testing.T) {
	item := "abillrstuv"
	low := 0
	high := len(item) - 1
	//test case 1
	t.Run("Binary search for \"(a)\"", func(t *testing.T) {
		if binarySearch(item, low, high, 'a') == -1 {
			t.Error("Binary search for \"(a)\"")
		} else {
			t.Log("Passed : Binary search for \"(a)\"")
		}
	})
	//test case 2
	t.Run("Binary searh for \"(d)\"", func(t *testing.T) {
		if binarySearch(item, low, high, 'd') != -1 {
			t.Error("Binary search for \"(d)\"")
		} else {
			t.Log("Passed : Binary search for \"(d)\"")
		}
	})
}

// test print data to console using fmt
func TestPrintData(t *testing.T) {
	//case 1 : print data to console, hello world
	t.Run("print hello world", func(t *testing.T) {
		if printData("hello world") != "hello world" {
			t.Error("print hello world")
		} else {
			t.Log("Passed : print hello world")
		}
	})
	//case 2 : print data to console, jack miller
	t.Run("print jack miller", func(t *testing.T) {
		if printData("jack miller") == "jack millers" {
			t.Error("print jack miller")
		} else {
			t.Log("Passed : print jack miller")
		}
	})
}

// test circle area function
func TestCircleArea(t *testing.T) {
	//case 1 : calculate circle area with radius 10
	t.Run("calculate circle area with radius 10", func(t *testing.T) {
		if circleArea(10) > 314.52 {
			t.Error("calculate circle area with radius 10")
		} else {
			t.Log("Passed : calculate circle area with radius 10")
		}
	})
	//case 2 : calculate circle area with radius 5
	t.Run("calculate circle area with radius 5", func(t *testing.T) {
		if circleArea(5) > 78.63 {
			t.Error("calculate circle area with radius 5")
		} else {
			t.Log("Passed : calculate circle area with radius 5")
		}
	})
}

// test convert price to int
func TestPriceToInt(t *testing.T) {
	//case 1 : convert price to int
	t.Run("convert price to int", func(t *testing.T) {
		if priceToInt(10000.000) != 10000 {
			t.Error("convert price to int")
		} else {
			t.Log("Passed : convert price to int")
		}
	})
	//case 2 : convert price to int
	t.Run("convert price to int", func(t *testing.T) {
		if priceToInt(20000.000) != 20000 {
			t.Error("convert price to int")
		} else {
			t.Log("Passed : convert price to int")
		}
	})
}

// test initAge function
func TestInitAge(t *testing.T) {
	//case 1 : init age
	t.Run("init age", func(t *testing.T) {
		if initAge()[0] != 10 {
			t.Error("init age")
		} else {
			t.Log("Passed : init age")
		}
	})
	//case 2 : init age
	t.Run("init age", func(t *testing.T) {
		if initAge()[1] != 12 {
			t.Error("init age")
		} else {
			t.Log("Passed : init age")
		}
	})
}

// test initName function
func TestInitName(t *testing.T) {
	//case 1 : init name
	t.Run("init name", func(t *testing.T) {
		if initName()[0] != "jack" {
			t.Error("init name")
		} else {
			t.Log("Passed : init name")
		}
	})
	//case 2 : init name
	t.Run("init name", func(t *testing.T) {
		if initName()[1] != "miller" {
			t.Error("init name")
		} else {
			t.Log("Passed : init name")
		}
	})
}

// test totalItems function
func TestTotalItems(t *testing.T) {
	initName()
	//case 1 : total items
	t.Run("total items", func(t *testing.T) {
		if totalItemsName(data.StudentName) != 10 {
			t.Error("total items for student name")
		} else {
			t.Log("Passed : total items for student name")
		}
	})
	//case 2 : total items
	t.Run("total items", func(t *testing.T) {
		if totalItemsAge(data.StudentAge) != 10 {
			t.Error("total items for student age")
		} else {
			t.Log("Passed : total items for student age")
		}
	})
}

// test init status of student
func TestInitStatus(t *testing.T) {
	//case 1 : init status
	t.Run("init status", func(t *testing.T) {
		if initStatus()["jack"] != true {
			t.Error("init status")
		} else {
			t.Log("Passed : init status")
		}
	})
	//case 2 : init status
	t.Run("init status", func(t *testing.T) {
		if initStatus()["jane"] == false {
			t.Error("init status")
		} else {
			t.Log("Passed : init status")
		}
	})
}

// test appendPort function
func TestAppendPort(t *testing.T) {
	//test case 1
	appendPort("smtp", 25)
	appendPort("udp", 53)
	t.Run("append port", func(t *testing.T) {
		if data.WellKnownPorts["smtp"] != 25 {
			t.Error("append port")
		} else {
			t.Log("Passed : append port")
		}
	})
	//case 2 : udp port
	t.Run("append port", func(t *testing.T) {
		if data.WellKnownPorts["udp"] != 53 {
			t.Error("append port")
		} else {
			t.Log("Passed : append port")
		}
	})
}
