package main

import (
	"errors"
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	costCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	costSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, err
	}
	return costCustomer + costSpouse, nil
}

func test(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("============")
	fmt.Println("Customer Message:", msgToCustomer)
	fmt.Println("Spouse Message:", msgToSpouse)
	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Total Cost: Kshs. %.2f\n", totalCost)
}

func sendSMS(message string) (float64, error)  {
	const maxTextLen = 25
	const costPerChar = 0.0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("Can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs Kshs. %.2f to be sent to '%v' cannot be sent.", cost, recipient)
}

func test1 (cost float64, recipient string) {
	s := getSMSErrorString(cost, recipient)
	fmt.Println(s)
	fmt.Println("==============================")
}

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("Can't divide %v by 0.", de.dividend)
}

func divide(x, y float64) (float64, error) {
	if y == 0 {
		dError := divideError{dividend: x}
		errorMess := dError.Error()
		return 0.0, fmt.Errorf("%v", errorMess)
	}
	quotient := x / y
	return quotient, nil
}

func divide1(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("No dividing by 0")
	}
	quotient := x / y
	return quotient, nil
}

func test2 (x, y float64) {
	q, err := divide1(x, y)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%.2f / %.2f = %.2f", x, y, q)
}

func main() {
	messageToCust := "Loved having you in."
	spouseMess := "Come again another time fam, we're waiting."
	test(messageToCust, spouseMess)
	messageToCust1 := "Loved having you in."
	spouseMess1 := "Come again soon."
	test(messageToCust1, spouseMess1)
	cost := 10.5
	recipient := "Ryan"
	test1(cost, recipient)
	x := 10.0
	y := 0.0
	y1 := 2.0
	test2(x, y)
	test2(x, y1)
}