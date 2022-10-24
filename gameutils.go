package hangman

import (
	"math/rand"
)

func GetRandomWord(liste []string) string {
	return liste[rand.Intn(len(liste))]
}

func HasOccurenceLetter(word string, letterToCheck rune) bool {
	for _, letter := range word {
		if letter == letterToCheck {
			return true
		}
	}
	return false
}

func GetOccurenceLetter(word string, letterToCheck rune) []int {
	var occ []int
	wordR := []rune(word)
	for i, letter := range wordR {
		if letter == letterToCheck {
			occ = append(occ, i)
		}
	}
	return occ
}

func UpdateGameWord(toFind string, word string, letterToCheck rune) string {
	wordR := []rune(word)
	indexs := GetOccurenceLetter(toFind, letterToCheck)
	for _, index := range indexs {
		wordR[index] = letterToCheck
	}
	word = string(wordR)
	return word
}

func SetupGameWord(startup_word string) string {
	size := len([]rune(startup_word))
	runeTableWord := make([]rune, size)
	for i := 0; i < len(runeTableWord); i++ {
		runeTableWord[i] = '_'
	}
	for nbLettersToShow := len([]rune(runeTableWord))/2 - 1; nbLettersToShow > 0; nbLettersToShow-- {
		randomTableI := rand.Intn(len([]rune(runeTableWord)))
		if runeTableWord[randomTableI] != '_' {
			nbLettersToShow++
		} else {
			runeTableWord[randomTableI] = []rune(startup_word)[randomTableI]
		}
	}
	listOfLettersGiven := make([]rune, len([]rune(runeTableWord)))
	for _, letter := range runeTableWord {
		if letter != '_' {
			listOfLettersGiven = append(listOfLettersGiven, letter)
		}
	}
	for _, letter := range listOfLettersGiven {
		runeTableWord = []rune(UpdateGameWord(startup_word, string(runeTableWord), letter))
		AddGameUsed(letter)
	}
	return string(runeTableWord)
}

//func GetWordInAscii(word string) []string {
//	asciiWord := make([]string, len([]rune(word)))
//	for i, letter := range word {
//		for _, asciiLetter := range GetAsciiArtFromRune(letter) {
//			asciiWord[i] += asciiLetter
//		}
//	}
//	return asciiWord
//}
