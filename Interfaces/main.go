package main

import (
	"fmt"
	"io"
	"math"
	"time"
	"os"
)

// When a type implements an interface, it can be used as the interface type
// I may not know what shape it is, but since it is a shape, I know I can get the area and the perimeter
// I can pass any struct that implements interface shape

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2 * r.width + 2 * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, your birthday is on %s.", bm.recipientName, bm.birthdayTime)
}

type sendingReport struct {
	reportName string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf("Your '%s' report is ready. You've sent %v reports.", sr.reportName, sr.numberOfSends)
}

func test(m message) {
	sendMessage(m)
	fmt.Println("================================================================================")
} 

// Employee
type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name string
	hourlyPay int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func test1(e employee) {
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("==========================================")
}

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body string
	toAddress string
}

type sms struct {
	isSubscribed bool
	body string
	toPhoneNumber string
}

// func print(p printer) {
// 	p.print()
// }

func (e email) print() {
	fmt.Println((e.body))
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return 0.05 * float64(len(e.body))
	}
	return 0.01 * float64(len(e.body))
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return 0.05 * float64(len(s.body))
	}
	return 0.01 * float64(len(s.body))
}

func test2 (e expense, p printer) {
	fmt.Printf("Printing with cost: Kshs. %.2f ...\n", e.cost())
	p.print()
	fmt.Println("================================================")
}

func getExpenseReport(e expense) (string, float64) {
	em, ok := e.(email)
	if ok {
		return em.toAddress, em.cost() 
	}

	sms, ok := e.(sms)
	if ok {
		return sms.toPhoneNumber, sms.cost()
	}

	return "", 0.0
}

func getExpenseReport2 (e expense) (string, float64) {
	switch v := e.(type){
	case email:
		return v.toAddress, v.cost() 
	case sms:
		return v.toPhoneNumber, v.cost() 
	default:
		return "", 0.0
	}
}

func test3 (e expense){
	address, cost := getExpenseReport2(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: Email going to %s will cost: Kshs. %.2f\n", address, cost)
		fmt.Println("=============================================================")
	case sms:
		fmt.Printf("Report: SMS going to %s will cost: Kshs. %.2f\n", address, cost)
		fmt.Println("=============================================================")
	default:
		fmt.Println("Report: Invalid Expense")
		fmt.Println("=============================================================")
	}
}

type invalid struct {}

func (i invalid) cost() float64 {
	return 0.0
}

// Type Switch
func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("Not int")
	}
}

type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]os.FileInfo, error)
	Stat() (os.FileInfo, error)
}

func main () {
	msg1 := birthdayMessage{
		recipientName: "Ryan",
		birthdayTime: time.Date(2003, 04, 05, 0, 0, 0, 0, time.Local),
	}
	msg2 := sendingReport {
		reportName: "Blue Lagoon",
		numberOfSends: 5,
	}
	test(msg1)
	test(msg2)

	test1(
		fullTime{
			name: "Tion",
			salary: 50000, 
		},
	)
	test1(
		contractor{
			name: "Wayne",
			hourlyPay: 25,
			hoursPerYear: 1300,
		},
	)

	e := email {
		isSubscribed: true,
		body: "Greeting Hedera.",
		toAddress: "jsil@gmail.com",
	}
	s := sms {
		isSubscribed: true,
		body: "Hi",
		toPhoneNumber: "0744668821",
	}
	test2(e, e)
	test3(e)
	test3(s)
	test3(invalid{})

	//Type Assertion
	// c - new circle cast from "s" - instance of shape -> It can be any shape
	// "ok" - bool that's true if s was a circle or false if not one
	// c, ok := s.(circle)
}