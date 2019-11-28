package main

import "github.com/jayvib/go_design_pattern/behavioral/state/number-guess"

func main() {
	start := number_guess.StartState{}
	game := number_guess.GameContext{
		Next: &start,
	}
	for game.Next.ExecuteState(&game) {
	}

}
