// testmain.go : unit testing for main.go
package main
//generate unit test for main.go
import (
	"testing"
)
//test for isEven function
func TestIsEven(t *testing.T){
	//test case 1
	if !isEven(2){
		t.Error("2 is even number")
	}
	//test case 2
	if isEven(3){
		t.Error("3 is not even number")
	}
}
//test for strLength function
func TestStrLength( t *testing.T){
	//test case 1
	if(strLength("hello") !=5){
		t.Error("Length of hello is 5")
	}
	//test case 2
	if(strLength("world") !=5){
		t.Error("Length of world is 5")
	}
	//test case for empty string
	if(strLength("") !=0){
		t.Error("Length of empty string is 0")
	}
}
//test for isValidIsbn function
func TestIsValidIsbn(t *testing.T){
	//test case 1
	t.Run("978-0xmath-calculus is valid isbn", func(t *testing.T){
		if(!isValidIsbn("978-0xmath-calculus")){
			t.Error("978-0xmath-calculus is valid isbn")
		}else{
			t.Log("Passed : 978-0xmath-calculus is valid isbn")
		}
	})
	//test case 2
	t.Run("0x9-0-393-04002-9 is invalid isbn", func(t *testing.T){
		if(isValidIsbn("0x9-0-393-04002-9")){
			t.Error("0x9-0-393-04002-9 is invalid isbn")
		}else{
			t.Log("Passed : 0x9-0-393-04002-9 is valid isbn")
		}
	})
}
//test for contains function
func TestContains(t *testing.T){
	//test case 1
	t.Run("contains([\"978\", \"979\"], \"978\") is true", func(t *testing.T){
		if (!contains([]string{"978","979"}, "978")){
			t.Error("contains([\"978\", \"979\"], \"978\") is true")
		}else{
			t.Log("Passed : contains([\"978\", \"979\"], \"978\") is true")
		}
	})
	//test case 2
	t.Run("contains([\"978\", \"979\"], \"977\") is false", func(t *testing.T){
		if(contains([]string{"978","979"}, "977")){
			t.Error("contains([\"978\", \"979\"], \"977\") is false")
		}else{
			t.Log("Passed : contains([\"978\", \"979\"], \"977\") is false")
		}
	})
}
//test for exceedMax function
func TestExceedMax(t *testing.T){
	//test case 1
	t.Run("100 is not exceed the maximum", func(t *testing.T){
		if(exceedMax(100)){
			t.Error("100 is not exceed the maximum")
		}else{
			t.Log("Passed : 100 is not exceed the maximum")
		}
	})
	//test case 2
	t.Run("1000 is exceed the maximum", func (t *testing.T)  {
		if(!exceedMax(1000)){
			t.Error("1000 is exceed the maximum")
		}else{
			t.Log("Passed : 1000 is exceed the maximum")
		}
	})
}
func TestSuccesSave(t *testing.T){
	//test case 1
	var file_name = "main.go"
	t.Run("main.go is saved to directory", func (t *testing.T)  {
		if(!succesSave(file_name)){
			t.Error("main.go is saved to database")
		}else{
			t.Log("passed : main.go is saved to database")
		}
	})
	//test case 2
	file_name="mailer.go"
	t.Run("mailer.go is not saved", func (t *testing.T)  {
		if(succesSave(file_name)){
			t.Error("mailer.go is not saved")
		}else{
			t.Log("mailer.go is not saved")
		}
	})
}