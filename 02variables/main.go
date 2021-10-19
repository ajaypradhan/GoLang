package main

import "fmt"

// public : Capital character at starting make that public
const LoginToken string = "asndajndakdaksd" 

func main() {
	var username string = "ajay"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n",username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n",isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n",smallVal)

	var smallFloat float32 = 255.45116
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n",smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("anotherVariable is of type: %T \n",anotherVariable)

	var anotherString string
	fmt.Println(anotherString);
	fmt.Printf("anotherString is of type: %T \n",anotherString)
	

	// implicit type or implicit way to declare variable
	var website ="learncodeonline.com"
	fmt.Println(website)

	// no var style
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("LoginToken is of type: %T \n",LoginToken)

}