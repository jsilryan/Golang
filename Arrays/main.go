package main

import (
	"fmt"
 	"errors"
)

type cost struct {
	day int
	value float64
}

const (
	planFree = "free"
	planPro = "pro"
)

func getMessageWithRetries() [3]string {
	return [3]string{
		"Click here to sign up.",
		"Remember to sign up.",
		"Get Bonus for signing up.", // Put last comma if I use newline
	}
}

func getMessagesWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()
	if plan == planPro {
		return allMessages[:], nil
	}
	if plan == planFree {
		return allMessages[0:2], nil
	}
	return nil, errors.New("Unsupported Plan")
}

func send(name string, messages []string) {
	fmt.Printf("Sending to %v...", name)
	fmt.Println()

	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf("Sending: '%v'\n", msg)
	}
} 

func test(name string, plan string) {
	defer fmt.Println("==========================================")

	messages, err := getMessagesWithRetriesForPlan(plan)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	send(name, messages)
	
}

func getMessagesCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		cost := float64(len(message)) * 0.01
		costs[i] = cost
	}
	return costs
}

func test2(plan string) {
	messages, err := getMessagesWithRetriesForPlan(plan)
	if err != nil {
		fmt.Printf("%v", err)
		return 
	}
	costs := getMessagesCosts(messages)
	fmt.Println("Messages - Cost")
	for i := 0; i < len(messages); i++ {
		fmt.Printf("%d. %v - Kshs. %.2f\n", i+1, messages[i], costs[i])
	}
	fmt.Println("====END REPORT====")
	fmt.Println("")
}

func sum(nums ...float64) float64 {
	total := 0.0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

func test3 (nums ...float64) {
	fmt.Printf("Summing %v costs...\n", len(nums))
	total := sum(nums...)
	fmt.Printf("Total Bill: %v\n", total)
	fmt.Println("====END REPORT====")
	fmt.Println("")
}

func getCostsByDay(costs []cost) []float64 {
	costsByDay := []float64{} // Used this instead of Make - very similar
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value
	}
	return costsByDay
}

func test4(costs []cost) {
	perDay := getCostsByDay(costs)
	fmt.Println("Day : Total Cost")
	for i := 0; i < len(perDay); i++ {
		fmt.Printf("%v : %.2f\n", i+1, perDay[i])
	}
	fmt.Println("====END REPORT====")
	fmt.Println("")
}

func createMatrix(rows, cols int) {
	fmt.Printf("%v * %v Matrix:\n", rows, cols)
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, i * j)
		}
		matrix = append(matrix, row)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%v ", matrix[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("====END REPORT====")
	fmt.Println("")
}

func indexOfFirstBadWord (msg []string, badWords []string) int {
	for i, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}

func test5(msg []string, badWords []string) {
	i := indexOfFirstBadWord(msg, badWords)
	fmt.Printf("Scanning message: %v for bad words:\n", msg)
	for _, badWord := range badWords {
		fmt.Println("--", badWord)
	}
	fmt.Printf("1st Bad Word Index: %v\n", i)
	fmt.Println("===============================================")
	fmt.Println("")
}

func main() {
	test("Valhala", planFree)
	test("Ryan", planPro)
	test("Oscar", "No Plan")
	
	test2(planFree)

	nums1 := []float64{4, 20, 5}
	nums2 := []float64{23, 49, 17, 33, 1}
	test3(nums1...)
	test3(nums2...)

	costs := []cost {
		{0, 4.0},
		{1, 2.1},
		{1, 3.1},
		{5, 2.5},
	}

	test4(costs)

	createMatrix(4,5)

	badWords := []string{"crap", "shit", "dang", "frick"}
	message := []string{"Hi", "Davi", ";", "crap", "a", "bird", "shit", "on", "me"}
	test5(message, badWords)

}