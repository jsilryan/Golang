package main 
import (
	"fmt"
	"errors"
	"os"
	"io"
)

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

// Functions as data
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

// Currying
func selfMath(mathFunc func(int, int) int) func (int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

func getLogger(formatter func(string, string) string) func(string, string) {
	return func(a, b string) {
		fmt.Println(formatter(a, b))
	}
}

func test(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("")
	defer fmt.Println("=====================================")
	logger := getLogger(formatter)
	fmt.Println("\nLogs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func formatterMain(a, b string) string {
	return a + ": " + b
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	// Open src file
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	// Close src file when CopyFile function returns
	defer src.Close()

	// Create destination file
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)

}

type user struct{
	name string
	admin bool
}	

const (
	logDeleted = "User Deleted"
	logNotFound = "User Not Found"
	logAdmin = "Admin Deleted"
)

func logAndDelete(users map[string]user, name string) (log string ){
	user, ok := users[name]
	defer delete(users, name)
	if !ok {
		return logNotFound
	}
	if user.admin {
		return logAdmin
	}
	return logDeleted
}

func test2 (users map[string]user, name string) {
	stmt := logAndDelete(users, name)
	fmt.Println(name, ":", stmt)
}

// Closure
func concatter() func (string) string {
	doc := ""
	return func (word string) string {
		doc += word + " "
		return doc
	}
}

func adder() func(int) int {
	total := 0
	return func (num int) int {
		total += num
		return total
	}
}

type emailBill struct {
	costInPennies int
}

func test3 (bills []emailBill) {
	defer fmt.Println("")
	defer fmt.Println("===========================================")
	countAdder, costAdder := adder(), adder()

	for i, bill := range bills {
		countAdder(i)
		costAdder(bill.costInPennies)
	}
	fmt.Printf("You have %v total email bills that cost %v cents.\n", countAdder(0), costAdder(0))
}

// Anonymous function
func doMath(f func(int) int, nums []int) []int {
	var results []int
	for _, n := range nums {
		results = append(results, f(n))
	}
	return results
}

func main() {
	fmt.Println("Add:", aggregate(2, 3, 4, add))
	fmt.Println("Mul:", aggregate(2, 3, 4, mul))

	squareFunc := selfMath(mul)
	doubleFunc := selfMath(add)
	val := 5
	fmt.Println("Square of", val, ":", squareFunc(val))
	fmt.Println("Double of", val, ":", doubleFunc(val))

	first := "Error Logged"
	checkedErrors := []error {
		errors.New("Storage space is empty."),
		errors.New("Threat detected."),
		errors.New("CPU is pegged."),
		errors.New("Invalid syntax."),
	}
	test(first, checkedErrors, formatterMain)

	users := map[string]user {
		"John" : {
			name: "John",
			admin: true,
		},
		"Gretel" : {
			name: "Gretel",
			admin: false,
		},
		"Jamie" : {
			name: "Jamie",
			admin: true,
		}, 
		"Darnell" : {
			name: "Darnell",
			admin: false,
		},
	}
	test2(users, "John")
	test2(users, "Ryan")
	test2(users, "Darnell")
	fmt.Println("")

	aggregator := concatter()
	aggregator("Mr.")
	aggregator("and")
	aggregator("Mrs.")
	fmt.Println(aggregator("Mutua"))
	fmt.Println("")

	bills := []emailBill {
		{costInPennies: 94},
		{costInPennies: 33},
		{costInPennies: 71},
		{costInPennies: 88},
	}
	test3(bills)

	nums := []int{2,3,4,5}
	allNumsDoubled := doMath(func(x int) int {
		return x * x
	}, nums)
	fmt.Println(nums, "doubled =", allNumsDoubled)
	fmt.Println("")

	messages := []string{
		"Hello, hope you are ok.",
		"I received your message.",
		"There's no place like home.",
	}
	printReports(messages)
}

func printReports(messages []string) {
	defer fmt.Println("")
	defer fmt.Println("==================================")
	for _, message := range messages {
		printCostReport(
			func (msg string) int {
				return len(msg) * 2
			}, message,
		)
	}	
}

func printCostReport(costCalculator func(x string) int, message string) {
	totalCost := costCalculator(message)
	fmt.Printf("Message: '%v' | Cost: Kshs. %v\n", message, totalCost)
}
