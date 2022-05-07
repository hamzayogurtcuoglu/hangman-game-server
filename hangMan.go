package main

import (
	"bytes"
	"examplee/word"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func getRandomWord() string {
	rand.Seed(time.Now().UTC().UnixNano())
	wordList := word.GetWordList()
	randomIndex := rand.Intn(len(wordList))
	randomSelectedWord := wordList[randomIndex]
	return randomSelectedWord
}

func main() {
	word := getRandomWord()
	play(word)
	fmt.Println("\nPlay Again? (Y/N)")
	var guess string
	fmt.Scan(&guess)
	for guess == "Y" {
		word := getRandomWord()
		play(word)
		fmt.Println("\nPlay Again? (Y/N)")
		fmt.Scan(&guess)
	}
}

func play(word string) {
	wordCompletion := strings.Repeat("_", (len(word)))
	fmt.Println(wordCompletion)
	guessed := false
	guessedLetters := []byte{}
	guessedWords := []string{}
	tries := 6
	fmt.Println("Let's play Hangman!")
	fmt.Println(displayHangman(tries))
	fmt.Println(wordCompletion)
	fmt.Print("\n")
	for !guessed && tries > 0 {
		var guess string
		fmt.Scan(&guess)
		if len(guess) == 1 && IsLetter(guess) {
			data := []byte(guess)
			if bytes.Contains(guessedLetters, data) {
				fmt.Println("You already guessed the letter", guess)
			} else if !bytes.Contains([]byte(word), data) {
				fmt.Println(guess, "is not in the word.")
				fmt.Println("You have just ", tries, " right")
				tries -= 1
				guessedLetters = append(guessedLetters, data...)
			} else {
				fmt.Println("Good job,", guess, " is in the word!")
				guessedLetters = append(guessedLetters, data...)
				foundedWordAsList := []byte(wordCompletion)
				for index, letter := range word {
					if letter == rune(data[0]) {
						foundedWordAsList[index] = byte(letter)
					}
				}
				if !bytes.Contains(foundedWordAsList, []byte("_")) {
					guessed = true
				}
				wordCompletion = string(foundedWordAsList)
			}
		}
		if len(guess) == len(word) && checkStringAlphabet(guess) {
			if contains(guessedWords, guess) {

				fmt.Println("You already guessed the word", guess)
			} else if guess != word {
				fmt.Println(guess, "is not the word.")
				tries -= 1
				guessedWords = append(guessedWords, guess)
			} else if guess == word {
				guessed = true
				wordCompletion = word
			}
		} else {
			fmt.Println("Not a valid guess.")
		}
		fmt.Println(displayHangman(tries))
		fmt.Println(wordCompletion)
		fmt.Print("\n")
	}
	if guessed {
		fmt.Print("Congrats, you guessed the word! You win!")
	} else {
		fmt.Print("Sorry, you ran out of tries. The word was " + word + ". Maybe next time!")
	}

}

func displayHangman(tries int) string {
	stages := [7]string{
		`
                   --------
                   |      |
                   |      O
                   |     \|/
                   |      |
                   |     / \
                   -
				   `,
		`
                   --------
                   |      |
                   |      O
                   |     \|/
                   |      |
                   |     / 
                   -
				   `,
		`
                   --------
                   |      |
                   |      O
                   |     \|/
                   |      |
                   |      
                   -
				   `,
		`
                   --------
                   |      |
                   |      O
                   |     \|
                   |      |
                   |     
                   -
				   `,
		`
                   --------
                   |      |
                   |      O
                   |      |
                   |      |
                   |     
                   -
				   `,
		`
                   --------
                   |      |
                   |      O
                   |    
                   |      
                   |     
                   -
				   `,
		`
                   --------
                   |      |
                   |      
                   |    
                   |      
                   |     
                   -
				   `,
	}
	return stages[tries]
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func checkStringAlphabet(str string) bool {
	for _, charVariable := range str {
		if (charVariable < 'a' || charVariable > 'z') && (charVariable < 'A' || charVariable > 'Z') {
			return false
		}
	}
	return true
}
