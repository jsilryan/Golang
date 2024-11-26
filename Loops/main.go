package main
import "fmt"

func bulkSend(numMess int) float64 {
	initialCost := 0.0
	for i := 0; i < numMess; i++ {
		// Can't multiply int with float
		initialCost += 1.0 + 0.01 * float64(i)
	}
	return float64(initialCost)
}

func maxMessages(thresh float64) int {
	totals := 0.0
	for i := 0; ; i++ {
		totals += 1.0 + 0.01 * float64(i)
		if (totals > thresh) {
			return i
		}
	}
}

func getMaxMessages2(costMultiplier float64, maxCostInPennies int) int {
	actualCost := 1.0
	maxMessagesToSend := 0
	for actualCost <= float64(maxCostInPennies) {
		maxMessagesToSend ++
		actualCost *= costMultiplier
	}
	return maxMessagesToSend
}

func test(numMess int) {
	fmt.Printf("Sending %v messages.\n", numMess)
	cost := bulkSend(numMess)
	fmt.Printf("Total Cost of Bulk = %.2f\n", cost)
	fmt.Println("===========================================")
}

func test2(thresh float64) {
	fmt.Printf("Threshold: %.2f\n", thresh)
	num := maxMessages(thresh)
	fmt.Printf("Max Messages: %v\n", num)
	fmt.Println("===========================================")
}

func test3(costMultiplier float64, maxCostInPennies int) {
	fmt.Printf("Multiplier: %.2f || Max Cost: %v\n", costMultiplier, maxCostInPennies)
	maxMess := getMaxMessages2(costMultiplier, maxCostInPennies)
	fmt.Printf("Max Messages: %v\n", maxMess)
	fmt.Println("===========================================")
}

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		// Let the 3 and 5 one be at the top
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i % 3 == 0 {
			fmt.Println("fizz")
		} else if i % 5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("===========================================")
}

func printPrimes(max int) {
	// Start at 2 because 1 is not prime
	for n := 2; n < max + 1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		// Skip even numbers
		if n % 2 == 0 {
			continue
		}
		// sqrt(n) => Check if i*i < n -> I don't want numbers > root of n
		// Anything higher than the root has no chance of multiplying evenly into n that numbers less than the root didn't
		// 16 -> 4 || 8 will divide into 16 because 2 already does 
		isPrime := true
		for i := 3; i * i < n + 1; i++ {
			if n % i == 0 {
				isPrime = false
				break
			}
		}
		if !isPrime {
			continue
		}
		fmt.Println(n)
	}
}

func test4(max int) {
	fmt.Printf("Primes up to %v\n", max)
	printPrimes(max)
	fmt.Println("===========================================")
}

func main() {
	totalMess := 10
	test(totalMess)
	thresh := 100.0
	test2(thresh)
	costMult, maxCost := 2.5, 40 // 1.1, 5 | 1.3, 10
	test3(costMult, maxCost)
	// fizzbuzz()
	max := 20
	max2 := 100
	test4(max)
	test4(max2)
}