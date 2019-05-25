package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// Each index is the str for a sequentially greater guess
var hangmanBodyArr = [7]string{ "___ \n|  | \n|   \n| \n| \n|____",
								"___ \n|  | \n|  O\n| \n| \n|____",
								"___ \n|  | \n|  O\n|  |\n| \n|____",
								"___ \n|  | \n|  O\n| /|\n| \n|____",
								"___ \n|  | \n|  O\n| /|\\\n| \n|____",
								"___ \n|  | \n|  O\n| /|\\\n| /\n|____",
								"___ \n|  | \n|  O\n| /|\\\n| / \\\n|____"}

// getLetterMap returns map of alphabet letters to slice of indices where the letter occurs in randomWord
func getLetterMap(randomWord string) map[string][]int{
	letterMap := make(map[string][]int)
	for i := 0; i < len(randomWord); i++{
		curLetter := string(randomWord[i])
		letterMap[curLetter] = append(letterMap[curLetter], i)
	}
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
	numIncorrectGuesses := 0
	var pastGuesses [26]int // 0 for not guessed, else 1 for guessed
	badGuesses := "Letters not in the secret word: " // Invalid guesses
	randomWord := getRandomWord()
	// Initialize user guess to underscores
	guessWord := ""
	for i := 0; i < len(randomWord); i++{
		guessWord += "_ "
	}
	guessWord = guessWord[:len(guessWord)-1]
	// add spaces to randomWord to correspond to underscores in guessword
	tmp := ""
	for i := 0; i < len(randomWord); i++{
		tmp += string(randomWord[i]) + " "
	}
	randomWord = tmp[:len(tmp)-1]
	letterMap := getLetterMap(randomWord)
	fmt.Println("")
	fmt.Println(hangmanBodyArr[numIncorrectGuesses])
	fmt.Println("The secret word: " + guessWord)
	fmt.Println("Guess a lowercase letter in the alphabet: ")
	// While user input is incorrect and user still has guesses
	for {
		// Get user input
		reader := bufio.NewReader(os.Stdin)
		guess, _ := reader.ReadString('\n')
		fmt.Println("")
		guess = string(guess[:len(guess)-2]) // Get rid of the delimiter
		guessIndex := guess[0]-'a'
		if len(guess) == 1 && guessIndex >= 0 && guessIndex <= 25{ // Valid guess
			if pastGuesses[guessIndex] == 0 { // If new guess
				pastGuesses[guessIndex] = 1
				curLetterMap := letterMap[guess]
				if len(curLetterMap) > 0 { // Correct guess
					fmt.Println("Nice guess! The letter '" + guess + "' is in the secret word.")
					for i := 0; i < len(curLetterMap); i++{

						guessWord = guessWord[:curLetterMap[i]] + guess + guessWord[curLetterMap[i]+1:]
					}
					fmt.Println("Updated secret word: " + guessWord)
					if randomWord == guessWord{ // Did user guess the entire word?
						fmt.Println("You've successfully guessed the secret word!")
						break
					}
				} else{ // Incorrect guess
					numIncorrectGuesses++;
					badGuesses += guess + " "
					fmt.Println("Uh-oh! The letter '" + guess + "' is not in the secret word.")
					fmt.Println(hangmanBodyArr[numIncorrectGuesses])
					fmt.Println(badGuesses)
					fmt.Println("Secret word: " + guessWord)
					if(numIncorrectGuesses == len(hangmanBodyArr)-1){ // Is user out of guesses?
						fmt.Println("Game over. The secret word was: " + randomWord)
						break
					}
				}
			} else{
				fmt.Println("You have already guessed the letter '" + guess + "'!")
			}
		} else{
			fmt.Println("Guess a lowercase letter in the alphabet: ")
		}
	}
}
