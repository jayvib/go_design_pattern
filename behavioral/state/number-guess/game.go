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
	SecretNumber int
	Retries      int
	Won          bool
	Next         GameState
}

type StartState struct{}

func (s *StartState) ExecuteState(c *GameContext) bool {
	c.Next = &AskState{}
	rand.Seed(time.Now().UnixNano())
	c.SecretNumber = rand.Intn(10)
	fmt.Println("Introduce a number a number of retries to set the difficulty:")
	fmt.Fscanf(os.Stdin, "%d\n", &c.Retries)
	return true
}

type AskState struct{}

func (a *AskState) ExecuteState(c *GameContext) bool {
	fmt.Printf("Introduce a number between 0 and 10, you have %d tries left\n", c.Retries)
	var n int
	fmt.Fscanf(os.Stdin, "%d", &n)
	c.Retries = c.Retries - 1
	if n == c.SecretNumber {
		c.Won = true
		c.Next = &FinishState{}
	}

	if c.Retries == 0 {
		c.Next = &FinishState{}
	}
	return true
}
