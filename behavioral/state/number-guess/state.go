package number_guess

import "fmt"

type FinishState struct{}

func (f *FinishState) ExecuteState(c *GameContext) bool {
	if c.Won {
		fmt.Println("Congrats, you won")
	} else {
		fmt.Println("You, lose")
	}
	return false
}
