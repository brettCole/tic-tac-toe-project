package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// board itself
func board(spaces []string) {
	fmt.Println(spaces[6] + "|" + spaces[7] + "|" + spaces[8])
	fmt.Println("------")
	fmt.Println(spaces[3] + "|" + spaces[4] + "|" + spaces[5])
	fmt.Println("------")
	fmt.Println(spaces[0] + "|" + spaces[1] + "|" + spaces[2])
}

// take player input to assign "X" or "O" and in which order to return
func playerInput() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Player 1: Do you want to play as X or O?")
	player_marker, _ := reader.ReadString('\n')
	player_marker = strings.ToUpper(player_marker)
	if player_marker == "X" {
		return []string{"X", "O"}
	} else {
		return []string{"O", "X"}
	}
}

// assigns desired position with player's marker to board
func placeMarker(spaces []string, marker string, position int) {
	// fmt.Println(marker)
	spaces[position] = marker
}

// checks board if there is a winner
func winCheck(spaces []string, mark string) bool {
	return ((spaces[6] == mark && spaces[7] == mark && spaces[8] == mark) || (spaces[3] == mark && spaces[4] == mark && spaces[5] == mark) || (spaces[0] == mark && spaces[1] == mark && spaces[2] == mark) || (spaces[6] == mark && spaces[3] == mark && spaces[0] == mark) || (spaces[7] == mark && spaces[4] == mark && spaces[1] == mark) || (spaces[8] == mark && spaces[5] == mark && spaces[2] == mark) || (spaces[6] == mark && spaces[4] == mark && spaces[2] == mark) || (spaces[8] == mark && spaces[4] == mark && spaces[0] == mark))
}

// returns boolean indicating whether a space is open on board
func spaceCheck(spaces []string, position int) bool {
	status := false
	if spaces[position] != "X" && spaces[position] != "O" {
		status = true
	}
	return status
}

// players input (number 1-9) and uses space_check() to see if position is open
func playerChoice(spaces []string) int {
	position := 0
	for {
		fmt.Print("Choose your next position (0-8): ")
		fmt.Scanf("%d", &position)
		if spaceCheck(spaces, position) == true {
			return position
		} else {
			continue
		}
	}
}

// randomly decide which player goes first
func chooseRandomFirstPlayer() string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 1
	if val := rand.Intn(max-min+1) + min; val == 0 {
		return "Player 2"
	} else {
		return "Player 1"
	}
}

// check if board is full and returns boolean
func fullBoardCheck(spaces []string) bool {
	for i := range spaces {
		if spaceCheck(spaces, i) {
			return false
		}
	}
	return true
}

// ask if players want to play again
func replay() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to play again? Enter Yes or No: ")
	shouldReplay, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Problem retrieve player input: %v\n", err)
	}
	shouldReplay = strings.TrimRight(shouldReplay, "\n")
	if strings.ToLower(shouldReplay) == "yes" {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("Welcome to Tic Tac Toe!")

	reader := bufio.NewReader(os.Stdin)
	whichPlayerTurn := chooseRandomFirstPlayer()
	markers := playerInput()
	gameOn := false

	for gameOn == false {
		fmt.Printf("%v will go first.\n", whichPlayerTurn)

		fmt.Println("Are you ready to play? Please enter Yes or No.")
		playerReady, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Problem reading input: %v \n", err)
		}
		playerReady = strings.TrimRight(playerReady, "\n")

		if strings.ToLower(playerReady) == "yes" {
			gameOn = true
			break
		} else {
			gameOn = false
			continue
		}
	}

Restart:
	// Reset the board
	var spaces []string
	for i := 0; i < 9; i++ {
		spaces = append(spaces, strconv.Itoa(i))
	}

	for gameOn == true {
		player1Marker, player2Marker := markers[0], markers[1]
		if whichPlayerTurn == "Player 1" {
			// Player 1's turn
			board(spaces)
			pickedPosition := playerChoice(spaces)
			placeMarker(spaces, player1Marker, pickedPosition)

			if winCheck(spaces, player1Marker) {
				board(spaces)
				fmt.Println("Congratulations! You have won the game!")
				gameOn = false
			} else {
				if fullBoardCheck(spaces) {
					board(spaces)
					fmt.Println("No winner! The game is a draw!")
					gameOn = false
					continue
				} else {
					whichPlayerTurn = "Player 2"
				}
			}
		} else {
			// Player 2's turn
			board(spaces)
			pickedPosition := playerChoice(spaces)
			placeMarker(spaces, player2Marker, pickedPosition)

			if winCheck(spaces, player2Marker) {
				board(spaces)
				fmt.Println("Player 2 has won!")
				gameOn = false
			} else {
				if fullBoardCheck(spaces) {
					board(spaces)
					fmt.Println("The game is a draw!")
					gameOn = false
					continue
				} else {
					whichPlayerTurn = "Player 1"
				}
			}
		}
	}
	if replay() {
		gameOn = true
		goto Restart
	} else {
		os.Exit(0)
	}
}
