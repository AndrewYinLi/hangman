package main

import (
	"bufio"
	"log"
	"os"
	"math/rand"
	"time"
	"fmt"
)

// Each index is the str for a sequentially greater guess
var hangmanBodyArr = [6]string{ "___ \n|  | \n|  O\n| \n| \n|____",
								"___ \n|  | \n|  O\n|  |\n| \n|____",
								"___ \n|  | \n|  O\n| /|\n| \n|____",
								"___ \n|  | \n|  O\n| /|\\\n| \n|____",
								"___ \n|  | \n|  O\n| /|\\\n| /\n|____",
								"___ \n|  | \n|  O\n| /|\\\n| / \\\n|____"}

// getLetterMap returns map of alphabet letters to slice of indices where the letter occurs in randomWord
func getLetterMap(randomWord string) map[string][]int{
	letterMap := make(map[string][]int)


	return letterMap
}

// getRandomWord returns a random str from "dictionary.txt"
func getRandomWord() string{
	dictionaryLen := 0
	// Get length of dictionary for generating random word line number
	file, err := os.Open("dictionary.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dictionaryLen++;
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// Get line number for random word from dictionary
	rand.Seed(time.Now().UnixNano())
	lineNumber := rand.Intn(dictionaryLen)
	// Get random word from dictionary
	randomWord := ""
	dictionaryIndex := 0
	file, err = os.Open("dictionary.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		if dictionaryIndex == lineNumber{
			randomWord = scanner.Text()
			break
		}
		dictionaryIndex++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return randomWord
}

func main() {
	randomWord := getRandomWord()
	// Initialize user guess to underscores
	guessWord := ""
	for i := 0; i < len(randomWord); i++{
		guessWord += "_"
	}
	// While user input is incorrect and user still has guesses
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		if randomWord == guessWord{
			break
		}
	}

}
