package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
)

const dictionaryapi = "https://api.dictionaryapi.dev/api/v2/entries/en/"

func main() {

	showPtr := flag.Bool("show", false, "Show the chosen word")
	flag.Parse()

	word, err := selectRandomWordFromFile()
	if err != nil {
		os.Exit(1)
	}
	var guess string
	if *showPtr {
		fmt.Println("The word is " + word)
	}

	printPos([]int{2, 2, 2, 2, 2}, "*****")
	for {
		fmt.Print("\nEnter guess: ")
		fmt.Scanln(&guess)
		real := checkIfActualWord(guess)
		if len(guess) != 5 || len(guess) < 5 {
			fmt.Println("Guess must be a 5 letter word!")
		} else if !real {
			fmt.Println("Not an actual word!")
		} else {
			out := checkRightLetterWorngPos(guess, word)
			printPos(out, guess)
			won := checkForWin(out)
			if won {
				fmt.Println("CONGRATULATIONS! You have won.\nThe word was " + word + ".")
				break
			}
		}
	}
}

func checkIfActualWord(word string) bool {
	resp, err := http.Get(dictionaryapi + word)
	var statusCode int
	if err != nil {
		fmt.Println(err.Error())
		statusCode = 500
	} else {
		statusCode = resp.StatusCode
	}
	if statusCode != 200 {
		return false
	}
	return true
}

func selectRandomWordFromFile() (word string, err error) {

	file, err := os.Open("dictionary.txt")
	var words []string
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}

	pos, err := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
	word = words[pos.Int64()]
	return
}

func checkForWin(out []int) (won bool) {
	won = true
	for _, val := range out {
		if val != 0 {
			won = false
		}
	}
	return
}

func printPos(pos []int, guess string) {
	green := string([]byte{27, 91, 51, 50, 109})
	reset := string([]byte{27, 91, 48, 109})
	yellow := string([]byte{27, 91, 51, 51, 109})

	for i, val := range pos {
		switch val {
		case 0:
			fmt.Print(green, " "+string(guess[i])+" ", reset)
		case 1:
			fmt.Print(yellow, " "+string(guess[i])+" ", reset)
		case 2:
			fmt.Print(" " + string(guess[i]) + " ")
		}
	}
	fmt.Print("\n")
}

func checkRightLetterWorngPos(guess string, word string) (res []int) {
	mapWord := make(map[rune][]int)
	letterCount := make(map[rune]int)

	for i, char := range word {
		mapWord[char] = append(mapWord[char], i)
	}

	var contains bool
	var rightPos bool
	for i, char := range guess {
		contains = false
		rightPos = false
		if mapWord[char] != nil && letterCount[char] < len(mapWord[char]) {
			letterCount[char] = letterCount[char] + 1
			contains = true
			for _, po := range mapWord[char] {
				if i == po {
					rightPos = true
				}
			}
		}
		if contains && rightPos {
			res = append(res, 0)
		} else if contains && !rightPos {
			res = append(res, 1)
		} else {
			res = append(res, 2)
		}
		contains = false
		rightPos = false
	}
	return
}
