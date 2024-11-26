package main

import (
	"fmt"
)

type messageToSend struct {
	message string
	phoneNumber int
}

type user struct {
	name string
	number int
}

// Private and Public fields -> Embed Private fields within Public user details
type sender struct {
	rateLimit int
	user
}

type messageToSend2 struct {
	message string
	sender user
	recipient user
}

type car struct {
	Make string
	Model string
}

// Embedded struct -> truck.Make -> prevents rewriting
type truck struct {
	car
	bedSize int
}

// Methods
type rect struct {
	width int
	height int
}

// No need to store area
func (r rect) area() int {
	return r.width * r.height
}

type authenticationInfo struct {
	username string
	password string
}

func (authI authenticationInfo) getBasicAuth() string {
	// Sprintf -> Returns a string without printing it
	return fmt.Sprintf(
		"Authorization: BASIC %s:%s", 
		authI.username, authI.password,
	)
}

func test(m messageToSend) {
	fmt.Printf("Sending message '%s' to %v\n", m.message, m.phoneNumber)
	fmt.Println("===================================================================")
}

func canSendMessage(m messageToSend2) bool {
	if m.message == "" || m.sender.name == "" || m.sender.number == 0 || m.recipient.name == "" || m.recipient.number == 0 {
		return false
	}
	return true
}

func test2(m messageToSend2) {
	fmt.Printf("Sending message '%s'\nFrom:\n Name: %s || Number %v\nTo:\n Name: %s || Number %v\n", m.message, m.sender.name, m.sender.number, m.recipient.name, m.recipient.number)
	fmt.Println("===================================================================")
}

func main () {
	toSend := messageToSend{
		phoneNumber: 118902783,
		message: "Thanks for signing up.",
	}
	toSend2 := messageToSend2{}
	toSend2.message = "You have an appointment tomorrow."
	toSend2.sender.name = "Sally"
	toSend2.sender.number = 121290431
	toSend2.recipient.name = "Ryan"
	// toSend2.recipient.number = 789021889

	canSend := canSendMessage(toSend2)
	if canSend {
		test2(toSend2)
	} else {
		fmt.Println("Fill in all required fields.")
	}
	test(toSend)

	// How to instantiate embedded structs
	lanesTruck := truck {
		bedSize: 10,
		car: car {
			Make: "Honda",
			Model: "Civic",
		},
	}
	fmt.Println(lanesTruck.Make, lanesTruck.Model, lanesTruck.bedSize)

	newRectangle := rect {
		width: 5,
		height: 3,
	}
	fmt.Printf("Width: %v | Height: %v || Area = %v", newRectangle.width, newRectangle.height, newRectangle.area())

	auth := authenticationInfo{
		username: "jsil",
		password: "ryan",
	}
	fmt.Printf("Auth Info -> %s", auth.getBasicAuth())

	// Anonymous Structs and Nested Structs - When only used once. Use named structs more often
	// Prevent re-using a struct def I don't intend to re-use
	// myCar := struct {
	// 	Make string
	// 	Model string
	// } {
	// 	Make: "Ford",
	// 	Model: "Mustang",
	// }
	// type car struct {
	// 	Make string
	// 	Model string
	// 	Wheel struct {
	// 		Radius int
	// 		Material string
	// 	}
	// }

}
