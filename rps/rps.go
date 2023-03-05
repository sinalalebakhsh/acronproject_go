package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK         = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER        = 1 // beats rock. (rock + 1) % 3 = 0
	SCISSORS     = 2 // beats paper. (paper + 1) % 3 = 0
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

type Round struct {
	Winner         int
	ComputerChoice string
	RoundResult    string
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	winner := 0

	switch computerValue {
	case ROCK:
		computerChoice = "Computer Chose ROCK"
		break
	case PAPER:
		computerChoice = "Computer Chose PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer Chose SCISSORS"
		break
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a draw"
		winner = DRAW
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "player wins!"
		winner = PLAYERWINS
	} else {
		roundResult = "computer wins!"
		winner = COMPUTERWINS
	}

	var result Round
	result.Winner = winner
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult

	return result
}
