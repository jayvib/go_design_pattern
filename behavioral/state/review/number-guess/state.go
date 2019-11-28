package number_guess

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type GameState interface {
	ExecuteState(*GameContext) bool
}

type GameContext struct {
	retries      int
	secretNumber int
	won          bool
	next         GameState
}

type StartState struct{}

func (s *StartState) ExecuteState(c *GameContext) bool {
	rand.Seed(time.Now().UnixNano())
	c.secretNumber = rand.Intn(10)

	fmt.Println("Introduce a number of retries:")
	fmt.Fscanf(os.Stdin, "%d\n", &c.retries)

	c.next = &AskState{}
	return true
}

type AskState struct{}

func (a *AskState) ExecuteState(c *GameContext) bool {
	// ask for a number
	var n int
	fmt.Println("Enter a number from 1-10:")
	fmt.Fscanf(os.Stdin, "%d\n", &n)

	if c.secretNumber == n {
		c.won = true
		c.next = &FinishState{}
	}
	// decrement the number of retries
	c.retries--
	// check if the answer match
	if c.retries == 0 {
		c.won = false
		c.next = &FinishState{}
	}
	return true
}

type FinishState struct{}

func (a *FinishState) ExecuteState(c *GameContext) bool {
	if c.won {
		fmt.Println("Congrats, you won!")
	} else {
		fmt.Println("You, lose :(")
	}
	return false
}
