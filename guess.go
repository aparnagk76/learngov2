package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	numberOfAttempts := 5
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	random := rand.Intn(100) + 1
	success := false
	for i := 1; i <= numberOfAttempts; i++ {
		fmt.Println("enter the guess number between 1-100")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		target, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if target < random {
			fmt.Println("yout target is low")
		} else if target > random {
			fmt.Println("your target is high")
		} else if target == random {
			fmt.Println("your guess is correct")
			success = true
			break
		}
		fmt.Println("you have left ", numberOfAttempts-i)
	}
	if !success {
		fmt.Println("your guesses are worng and the correct number is ", random)
	}
}
