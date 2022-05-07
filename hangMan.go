package main

import (
	"bytes"
	"examplee/word"
	"fmt"
	"math/rand"
	"net"
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
	server()
}

func play(conn net.Conn) {

	conn.Write([]byte("WELCOME My Hangman GAME\r\n"))

	for {

		word := getRandomWord()
		wordCompletion := strings.Repeat("_", (len(word)))
		guessed := false
		guessedLetters := []byte{}
		guessedWords := []string{}
		tries := 6

		conn.Write([]byte("Let's play Hangman!"))
		conn.Write([]byte("\r\n"))
		conn.Write([]byte(displayHangman(tries)))
		conn.Write([]byte("\r\n"))
		conn.Write([]byte(wordCompletion))
		conn.Write([]byte("\r\n"))
		fmt.Println("Let's play Hangman!")
		fmt.Println(displayHangman(tries))
		fmt.Println(wordCompletion)
		fmt.Print("\n")

		for !guessed && tries > 0 {
			fmt.Println("The word --> ", word)
			conn.Write([]byte("\r\n"))
			conn.Write([]byte("Type just a letter : "))
			conn.Write([]byte("\r\n"))
			guessT := make([]byte, 1)
			conn.Read(guessT)
			guess := string(guessT)
			if len(guess) == 1 && IsLetter(guess) {
				data := []byte(guess)
				if bytes.Contains(guessedLetters, data) {
					mes := "You already guessed the letter " + guess
					fmt.Println(mes)
					conn.Write([]byte(mes))
				} else if !bytes.Contains([]byte(word), data) {
					mes1 := guess + " is not in the word."
					mes2 := "You have just " + string(tries) + " right"
					fmt.Println(guess, "is not in the word.")
					fmt.Println("You have just ", tries, " right")
					conn.Write([]byte(mes1))
					conn.Write([]byte(mes2))

					tries -= 1
					guessedLetters = append(guessedLetters, data...)
				} else {
					mes1 := "Good job, " + guess + " is in the word!"
					conn.Write([]byte(mes1))
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
					mes := "You already guessed the word " + guess
					conn.Write([]byte(mes))
					fmt.Println("You already guessed the word", guess)

				} else if guess != word {
					mes := guess + " is not the word."
					conn.Write([]byte(mes))
					fmt.Println(guess, "is not the word.")
					tries -= 1
					guessedWords = append(guessedWords, guess)
				} else if guess == word {
					guessed = true
					wordCompletion = word
				}
			}
			conn.Write([]byte("\r\n"))
			conn.Write([]byte(displayHangman(tries)))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte(wordCompletion))
			conn.Write([]byte("\r\n"))
			fmt.Println(displayHangman(tries))
			fmt.Println(wordCompletion)
			fmt.Print("\n")
		}
		if guessed {
			conn.Write([]byte("Congrats, you guessed the word! You win!"))
			fmt.Print("Congrats, you guessed the word! You win!")
		} else {
			mes := "Sorry, you ran out of tries. The word was " + word + ". Maybe next time!"
			conn.Write([]byte(mes))
			fmt.Print("Sorry, you ran out of tries. The word was " + word + ". Maybe next time!")
		}

		var guess string
		conn.Write([]byte("\r\nPlay Again? (y/n)"))
		fmt.Println("\nPlay Again? (y/n)")
		guessT := make([]byte, 1)
		conn.Read(guessT)
		guess = string(guessT)
		if guess != "y" {
			break
		}
	}
	conn.Close()
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
	stages[0] = strings.Replace(stages[0], "\n", "\r\n", 6)
	stages[1] = strings.Replace(stages[1], "\n", "\r\n", 6)
	stages[2] = strings.Replace(stages[2], "\n", "\r\n", 6)
	stages[3] = strings.Replace(stages[3], "\n", "\r\n", 6)
	stages[4] = strings.Replace(stages[4], "\n", "\r\n", 6)
	stages[5] = strings.Replace(stages[5], "\n", "\r\n", 6)
	stages[6] = strings.Replace(stages[6], "\n", "\r\n", 6)
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
