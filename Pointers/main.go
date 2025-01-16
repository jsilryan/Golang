package main

import (
	"fmt"
	"strings"
)

type Message struct {
	Recipient string
	Text      string
}

func sendMessage(m Message) {
	recipient1 := &m.Recipient
	text1 := &m.Text
	fmt.Printf("To: %v\n", *recipient1)
	fmt.Printf("Message: %v\n", *text1)

	*recipient1 = "Abdul"
	*text1 = "Have you watched the Liverpool game?"

	fmt.Printf("To: %v\n", m.Recipient)
	fmt.Printf("Message: %v\n", m.Text)
	fmt.Println("===================================")
	fmt.Println("")
}

// Takes in a pointer to a string - mutate without copying explicitly
func removeProfanity(message *string) {
    // Case sensitive
    if message == nil {
        fmt.Println("Invalid message!")
        return
    }
    messageVal := *message
    messageVal = strings.ReplaceAll(messageVal, "Dang", "****")
    messageVal = strings.ReplaceAll(messageVal, "Shoot", "*****")
    messageVal = strings.ReplaceAll(messageVal, "Heck", "****")
    *message = messageVal
}

func test(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
    message := ""
    nilMess := &message
    nilMess = nil
    removeProfanity(nilMess)
    fmt.Println("")
}

type car struct {
    color string
}

// Pointer of c is passed onto the method
func (c *car) setColor(color string) {
    c.color = color
}

type email struct {
    message string
    fromAddress string
    toAddress string
}

func (e email) print() {
    fmt.Println("Message:", e.message)
    fmt.Println("From:", e.fromAddress)
    fmt.Println("To:", e.toAddress)
    fmt.Println("")
}

func (e *email) setMessage(newMessage string) {
    e.message = newMessage
}

func test2(e *email, newMessage string) {
    fmt.Println("--- Before ---")
    e.print()
    e.setMessage(newMessage)
    fmt.Println("--- After ---")
    e.print()
    fmt.Println("====================================")
    fmt.Println("")
}

func main() {
	fullMessage := Message{
		Recipient: "Ryan",
		Text:      "How are you sir?",
	}

	sendMessage(fullMessage)

    messages := [](string){
        "Dang, should have scored that.",
        "Wow, what the heck has he done!",
        "I'm impressed.",
        "Shoot.",
    }
    test(messages)

    var x int = 50
    var y *int = &x // Normal pointer
    fmt.Printf("X = %v\n", x)
    fmt.Printf("&X = %v\n", &x)
    fmt.Printf("Y = %v\n", y)
    fmt.Printf("*Y = %v\n", *y)
    *y = 100
    fmt.Printf("*Y = %v\n", *y)
    fmt.Printf("*X = %v\n", x)

    c := car {
        color: "White",
    }
    fmt.Println(c.color)
    c.setColor("Red")
    fmt.Println(c.color)

    test2(
        &email {
            message: "My first draft.",
            fromAddress: "3waystreet@gmail.com",
            toAddress: "2ways@gmail.com",
        },
        "My second draft.",
    )
}
