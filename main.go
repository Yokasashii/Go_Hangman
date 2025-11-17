package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var Hangman = []string{
	`
  +---+
  |   |
      |
      |
      |
      |
========= 
	`,
	`
  +---+
  |   |
  O   |
      |
      |
      |
========= 
	`,
	`
  +---+
  |   |
  O   |
  |   |
      |
      |
========= 
	`,
	`
  +---+
  |   |
  O   |
 /|   |
      |
      |
========= 
	`,
	`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
========= 

	`,
	`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
========= 
	`,
	`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
========= 
	`,
}

var WinStreak = 0

func main() {
	var Word = []rune(CreatWord())
	var WordTry []rune
	Try := 6
	fmt.Println("\033[38;2;231;222;121m", `
		██╗  ██╗ █████╗ ███╗   ██╗ ██████╗ ███╗   ███╗ █████╗ ███╗   ██╗
		██║  ██║██╔══██╗████╗  ██║██╔════╝ ████╗ ████║██╔══██╗████╗  ██║
		███████║███████║██╔██╗ ██║██║  ███╗██╔████╔██║███████║██╔██╗ ██║
		██╔══██║██╔══██║██║╚██╗██║██║   ██║██║╚██╔╝██║██╔══██║██║╚██╗██║
		██║  ██║██║  ██║██║ ╚████║╚██████╔╝██║ ╚═╝ ██║██║  ██║██║ ╚████║
		╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝`, "\033[0m")
	Menu(&Word, WordTry, Try)
}

func Menu(Word *[]rune, WordTry []rune, Try int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1 - Play")
		fmt.Println("2 - Leave")
		fmt.Print("Select an option : ")

		Input, _ := reader.ReadString('\n')
		Choice := strings.TrimSpace(Input)

		switch Choice {
		case "1":
			fmt.Println("Lets go Play !")
			*Word = []rune(CreatWord())
			StartGame(*Word, WordTry, Try)
		case "2":
			fmt.Println("Bye bye !!")
			os.Exit(0)
		default:
			ClearScreen()
			fmt.Println("Your choice is incorrect. Please try again !")
		}
	}
}

func StartGame(Word []rune, WordTry []rune, Try int) {
	Game(Word, &WordTry, &Try)
}

func Game(Word []rune, WordTry *[]rune, Try *int) {
	GameWord := []rune{'_', '_', '_', '_', '_'}
	for !GameOver(Word, GameWord, Try) {
		ClearScreen()
		fmt.Println("\033[38;2;231;222;121m", `
		██╗  ██╗ █████╗ ███╗   ██╗ ██████╗ ███╗   ███╗ █████╗ ███╗   ██╗
		██║  ██║██╔══██╗████╗  ██║██╔════╝ ████╗ ████║██╔══██╗████╗  ██║
		███████║███████║██╔██╗ ██║██║  ███╗██╔████╔██║███████║██╔██╗ ██║
		██╔══██║██╔══██║██║╚██╗██║██║   ██║██║╚██╔╝██║██╔══██║██║╚██╗██║
		██║  ██║██║  ██║██║ ╚████║╚██████╔╝██║ ╚═╝ ██║██║  ██║██║ ╚████║
		╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝`, "\033[0m")
		fmt.Println("WinStreak :", WinStreak)
		PrintHangman(Try)
		fmt.Println(strings.ToTitle(string(GameWord)))
		fmt.Println("You already wrote : ", string(*WordTry))
		TapeLetter := ScanKeyboard(*WordTry)
		*WordTry = append(*WordTry, TapeLetter)
		if VerifLetter(TapeLetter, Word) {
			for i := 0; i <= 4; i++ {
				if TapeLetter == (Word)[i] {
					GameWord[i] = TapeLetter
				}
			}
		} else {
			(*Try)--
		}
	}
}

func CreatWord() string {
	Dictionnaire := []string{
		"apple", "bread", "chair", "table", "light", "water", "house", "grass", "music", "plant",
		"beach", "cloud", "heart", "smile", "sound", "sweet", "dream", "fruit", "drink", "sleep",
		"laugh", "world", "night", "dress", "stone", "floor", "green", "white", "black", "brown",
		"happy", "watch", "story", "river", "money", "dance", "sugar", "honey", "storm", "month",
		"hands", "earth", "place", "train", "drive", "glass", "space", "ocean", "grass", "round",
	}
	return Dictionnaire[rand.Intn(50)]
}

func PrintHangman(Try *int) {
	fmt.Println(Hangman[6-*Try])
}

func ScanKeyboard(WordTry []rune) rune {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type a letter : ")
	scanner.Scan()
	input := scanner.Text()
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	if len(input) == 1 {
		Runeletter := []rune(input)
		letter := Runeletter[0]
		if letter >= 'a' && letter <= 'z' {
			return letter
		} else {
			fmt.Println("You must write one letter.")
			return 0
		}
	} else {
		fmt.Println("You must write one letter.")
		return 0
	}
}

func VerifLetter(TapeLetter rune, Word []rune) bool {
	for i := 0; i <= 4; i++ {
		if TapeLetter == (Word)[i] {
			return true
		}
	}
	return false
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func GameOver(Word []rune, GameWord []rune, Try *int) bool {
	if string(Word) == string(GameWord) {
		ClearScreen()
		WinStreak = WinStreak + 1
		fmt.Println("You won because the word was", "\033[35m"+strings.ToTitle(string(Word))+"\033[0m")
		fmt.Println("\033[32m", `
	██╗   ██╗ ██████╗ ██╗   ██╗    ██╗    ██╗██╗███╗   ██╗
        ╚██╗ ██╔╝██╔═══██╗██║   ██║    ██║    ██║██║████╗  ██║
         ╚████╔╝ ██║   ██║██║   ██║    ██║ █╗ ██║██║██╔██╗ ██║
          ╚██╔╝  ██║   ██║██║   ██║    ██║███╗██║██║██║╚██╗██║
           ██║   ╚██████╔╝╚██████╔╝    ╚███╔███╔╝██║██║ ╚████║
           ╚═╝    ╚═════╝  ╚═════╝      ╚══╝╚══╝ ╚═╝╚═╝  ╚═══╝`, "\033[0m")
		return true
	} else if *Try == 0 {
		ClearScreen()
		WinStreak = 0
		fmt.Println("You lost, it's Game Over because the word was", "\033[35m"+strings.ToTitle(string(Word))+"\033[0m", "it's ciao kombucha")
		fmt.Println("\033[31m", `
		 ██████╗  █████╗ ███╗   ███╗███████╗     ██████╗ ██╗   ██╗███████╗██████╗ 
		██╔════╝ ██╔══██╗████╗ ████║██╔════╝    ██╔═══██╗██║   ██║██╔════╝██╔══██╗
		██║  ███╗███████║██╔████╔██║█████╗      ██║   ██║██║   ██║█████╗  ██████╔╝
		██║   ██║██╔══██║██║╚██╔╝██║██╔══╝      ██║   ██║╚██╗ ██╔╝██╔══╝  ██╔══██╗
		╚██████╔╝██║  ██║██║ ╚═╝ ██║███████╗    ╚██████╔╝ ╚████╔╝ ███████╗██║  ██║
		╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝     ╚═════╝   ╚═══╝  ╚══════╝╚═╝  ╚═╝`, "\033[0m")
		return true
	} else {
		return false
	}
}
