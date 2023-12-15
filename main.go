package main
//global variable
var url string = "https://www.google.com"
func main(){
	println("url : ", url)
}
//isEven : check if number is even or not
func isEven(number int) bool{
	return number % 2 == 0 //return true if number is even
}
//strLength : return length of string
func strLength(str string) int{
	return len(str)
}
//isValidIsbn : check if isbn is valid or not
func isValidIsbn(isbn string) bool{
	//vector for store isbn identifier
	var isbnIdentifier []string = []string{"978", "979"}
	//check if isbn identifier is valid
	return contains(isbnIdentifier, isbn[0:3])
}
//contains : check if array contains value or not
func contains(arr []string, value string) bool{
	for _, v := range arr{ //loop through array
		if v==value {
			return true
		}
	}
	return false
}
//exceedMax : check if number exceed max or not
func exceedMax(val int)bool{
	const MAKS_VALUE = 999;
	return val>MAKS_VALUE;
}
//sucessSave : check if fail succes saved in the directory
func succesSave(file_name string) bool{
	//vector than contain file name in the curr directory
	var file_dir []string =[]string{"main.go","test_main.go"}
	for _, gofile :=range file_dir{
		if(gofile==file_name){
			return true
		}
	}
	return false
}