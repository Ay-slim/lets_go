package main

import "fmt"

func main() {
	// The if statement is basically intuitive coming from JS. No ternary operators doth exist here.
	if ("a" == "b") {
		fmt.Println("We're fucked")
	} else {
		fmt.Println("Sike! We're still fucked")
	}

	//  You can declare a variable scoped to an if block and use it within the if block- much like variables declared for a for loop
	if rah := 9; rah < 2 {
		fmt.Println("This makes no sense")
	} else if rah == 9 {
		fmt.Println("Thou sayest")
	} else {
		fmt.Println("For fuck's sake")
	}

	// Switch statements
	// A switch statement with a "subject" variable checks if each case aligns with the variable value
	r := 3
	switch r {
	case 8:
		fmt.Println("Is it eight?")
	case 3:
		fmt.Println("It is 3")
	}
	// Without a subject variabe, it can check conditions just like an if
	switch {
	case r == 8:
		fmt.Println("No subject is it eight?")
	case r == 3:
		fmt.Println("No subject, it is 3")
	}
	}