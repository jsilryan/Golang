package main 

// Declares package file belongs to - creates the .exe files
// Main function should be entry point

import "fmt" // Formatted io - printing text on console

func concat (s1 string, s2 string) string {
	return s1 + s2
}

func add (x, y int) int {
	return x + y
}

func incrementSends(sent, toAdd int) int {
	return sent + toAdd
}

// Multiple return values
func getNames() (string, string) {
	return "Ryan", "Silu"
}

// getCoords and getCords2 are the same
// Naming return values acts a documentation 
func getCoords() (x, y int) {
	// x and y are initialized with 0 values

	return // Automatically returns x and y - Harm readability - Explicit returns are better than implicit returns
}

func getCoords2() (int, int) {
	var x int
	var y int
	return x, y
}

func yearsUntilEvents(age int) (adultAge, drinkingAge, carRentalAge int) {
	adultAge = 18 - age
	drinkingAge = 21 - age
	carRentalAge = 25 - age

	if adultAge < 0 {
		adultAge = 0
	}
	if drinkingAge < 0 {
		drinkingAge = 0
	}
	if carRentalAge < 0 {
		carRentalAge = 0
	}
	return adultAge, drinkingAge, carRentalAge
}

func main()  {
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string
	
	congrats := "Congratulations"
	penniesPerText := 2.0

	fmt.Printf(
		"%v\n%f\n%v\n%v\n%v\n%T",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		username,
		congrats,
		penniesPerText,
	)

	messageLen := 10
	maxMessageLen := 20

	fmt.Println("\nTrying to send message of length:", messageLen, "with max length:", maxMessageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message Sent.")
	} else {
		fmt.Println("Message Not Sent.")
	}
	fmt.Println(concat("\nLouis Lane", " treated Superman like shit."))

	sendsSoFar := 400
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("You've sent", sendsSoFar, "messages")

	firstName, _ := getNames()
	fmt.Println("Welcome", firstName)

	age := 20
	adultAge, drinkingAge, carRentalAge := yearsUntilEvents(age)
	fmt.Println("Years until Adult:", adultAge)
	fmt.Println("Years until Drinking:", drinkingAge)
	fmt.Println("Years until Car Rental:", carRentalAge)

}

