package main

import (
	"fmt"
	"errors"
	"io"
	"crypto/md5"
	"encoding/hex"
)

type user struct {
	name string
	phone int
}

type user2 struct {
	name string
	phone int
	toDelete bool
}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("Length of names and phone numbers are not equal.")
	}
	users := make(map[string]user)
	for i, name_1 := range(names) {
		users[name_1] = user{
			name: name_1,
			phone: phoneNumbers[i],
		}
	}
	return users, nil
}

func test(names []string, phoneNumbers []int) {
	fmt.Println("Creating map...")
	new_map, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Map: %v\n", new_map)
	fmt.Println("==================================================")
	fmt.Println("")
}

func deleteIfNecessary(users map[string]user2, name string) (deleted bool, err error) {
	user_1, ok := users[name]
	if !ok {
		return false, fmt.Errorf("User %v does not exist.", name)
	}
	if user_1.toDelete {
		delete(users, name)
		return true, nil
	} else {
		return false, nil
	}
}

func test2(users map[string]user2, name string) {
	fmt.Println("Trying to delete...")
	deleted, err := deleteIfNecessary(users, name)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	if deleted {
		fmt.Println("Deleted", name, ".")
	} else {
		fmt.Println(name, "was not scheduled for deletion.")
	}
	fmt.Println("===================================================")
	fmt.Println("")
}

func getCounts(userIDs []string) map[string]int {
	all_ids := make(map[string]int)

	for _, id := range(userIDs) {
		_, ok := all_ids[id]
		if !ok {
			all_ids[id] = 1
		} else {
			all_ids[id] += 1
		}

		// count := all_ids[id]
		// count++
		// all_ids[id] = count
	}
	return all_ids
}

func test3 (userIDs []string) {
	fmt.Printf("Generating counts for %v user IDs...\n", len(userIDs))
	counts := getCounts(userIDs)
	fmt.Printf("%v\n", counts)
	fmt.Println("===============================================")
	fmt.Println("")
}

// Nested Map
func getNameCounts(names []string) map[rune]map[string]int {
	initialLetters := make(map[rune]map[string]int)
	for _, name := range(names) {
		if name == "" {
			continue
		}
		firstChar := name[0] // It is a byte type - convert to Rune
		_, ok := initialLetters[rune(firstChar)]
		if !ok {
			initialLetters[rune(firstChar)] = make(map[string]int)
		} 
		initialLetters[rune(firstChar)][name] ++		
	}
	return initialLetters
}

func test4 (names []string, initial rune, name string) {
	fmt.Printf("Generating counts for %v names...\n", len(names))

	nameCounts := getNameCounts(names)
	fmt.Printf("%v\n", nameCounts)
	count := nameCounts[initial][name]
	fmt.Printf("Count for [%c][%s]: %d\n", initial, name, count)
	fmt.Println("=============================================")
	fmt.Println("")
}

func main() {
	names := []string{"Ryan", "Bob", "Jacob"}
	numbers := []int {743898392, 884932198, 778102948}
	test(names, numbers)

	names_1 := []string{"Ryan", "Bob", "Jacob"}
	numbers_1 := []int {743898392, 884932198}
	test(names_1, numbers_1)

	users := map[string]user2 {
		"John": {
			name: "John",
			phone: 790340129,
			toDelete: true,
		},
		"Janet" : {
			name: "Janet",
			phone: 710284398,
			toDelete: false,
		},
		"Alloise" : {
			name: "Alloise",
			phone: 723098201,
			toDelete: false,
		},
	}
	fmt.Printf("Current Users:\n%v\n\n", users)
	test2(users, "Alloise")
	test2(users, "John")
	test2(users, "Musk")
	fmt.Printf("New Users:\n%v\n\n", users)

	// List of user data (e.g., names or email addresses)
	userData := []string{"alice@example.com", "bob@example.com", "carol@example.com", "bluesky@gmail.com"}

	// Slice to store user IDs
	var userIDs []string

	// Generate MD5 hashes for each user data
	for _, data := range userData {
		// Create an MD5 hash
		hash := md5.New()
		_, err := io.WriteString(hash, data)
		if err != nil {
			fmt.Println("Error hashing data:", err)
			continue
		}
		// Convert the hash to a hexadecimal string
		userID := hex.EncodeToString(hash.Sum(nil))
		// Append the user ID to the slice
		userIDs = append(userIDs, userID)
	}

	// Print the user IDs
	fmt.Println("Generated User IDs:")
	for i, id := range userIDs {
		fmt.Printf("%d: %s\n", i+1, id)
	}
	fmt.Println("")

	test3(userIDs)

	userIDs2 := []string {"fd", "33i", "33i", "c5", "j9", "hbar", "hbar", "33i", "a23"}
	test3(userIDs2)

	names10 := []string{
		"Alice", "Bob", "Carol", "Dave", "Eve", 
		"Frank", "Grace", "Hank", "Alice", "Ivy",
	}

	// Slice of 17 names (some repeated)
	names17 := []string{
		"Jack", "Kathy", "Liam", "Mona", "Nancy", "Naomi",
		"Oliver", "Pam", "Quincy", "Liam", "Kathy",
		"Ryan", "Sam", "Tina", "Mona", "Uma", "Lisa", 
	}

	// Slice of 21 names (some repeated)
	names21 := []string{
		"Victor", "Wendy", "Xander", "Yara", "Zane", 
		"Victor", "Alice", "Bob", "Xander", "Frank",
		"Grace", "Yara", "Liam", "Hank", "Zane", 
		"Ivy", "Uma", "Pam", "Quincy", "Jack", "Vivian",
	}

	test4(names10, 'A', "Alice")
	test4(names17, 'L', "Lisa")
	test4(names21, 'Y', "Yara")
}