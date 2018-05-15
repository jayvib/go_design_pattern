package vehicle

import (
	"testing"
	"strings"
	"fmt"
)

func TestJunkComplex(t *testing.T) {
	junkVehicle := &JunkVehicle{Wheel: 4, Door: 4, Seat: 5}

	junkComplex := &JunkComplex{ parts:junkVehicle }

	if !strings.Contains(junkComplex.DoorsIGot(), "got 4 doors") {
		t.Error("expecting message to have 4 doors but got", junkComplex.DoorsIGot())
	}

	if !strings.Contains(junkComplex.SeatsIGot(), "got 5") {
		t.Error("expecting message to have 5 seats but got", junkComplex.SeatsIGot())
	}

	if !strings.Contains(junkComplex.WheelsIGot(), "got 4") {
		t.Error("expecting message to have 4 wheels but got", junkComplex.WheelsIGot())
	}

	junkVehicleWithExtraOnePartEach := &PartAddedJunkVehicle{ jv:junkVehicle }

	t.Run("junkVehicleWithExtraOnePartEach", func(t *testing.T){

		junkComplex.parts = junkVehicleWithExtraOnePartEach

		if !strings.Contains(junkComplex.DoorsIGot(), "got 5 doors") {
			t.Error("expecting message to have 5 doors but got", junkComplex.DoorsIGot())
		}

		if !strings.Contains(junkComplex.SeatsIGot(), "got 6") {
			t.Error("expecting message to have 6 seats but got", junkComplex.SeatsIGot())
		}

		if !strings.Contains(junkComplex.WheelsIGot(), "got 5") {
			t.Error("expecting message to have 7 wheels but got", junkComplex.WheelsIGot())
		}

		fmt.Println(junkComplex.IGot())
	})
}
