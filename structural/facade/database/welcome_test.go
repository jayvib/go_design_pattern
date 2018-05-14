package database_test

import (
	"testing"
	"github.com/go_design_pattern/structural/facade/database"
	"fmt"
)

func TestWelcomePage(t *testing.T) {
	expected := "Welcome Jayson!"

	message := database.WelcomePage("jayson.vibandor@gmail.com")

	if message != expected {
		t.Error("expected message doesn't match")
	}

	fmt.Println(message)
}
