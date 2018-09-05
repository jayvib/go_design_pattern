package recap

import "testing"

func TestCareTaker_Add(t *testing.T) {
	originator := originator{}
	originator.state = State{Description: "Idle"}

	careTaker := careTaker{}
	mem := originator.NewMomento()
	if mem.state.Description != "Idle" {
		t.Error("expected state was not found")
	}

	currentLen := len(careTaker.momentoList)
	careTaker.Add(mem)

	if len(careTaker.momentoList) != currentLen+1 {
		t.Error("no new elements were added on the list")
	}
}

func TestCareTaker_Memento(t *testing.T) {
	originator := originator{}
	careTaker := careTaker{}

	originator.state = State{"Idle"}
	careTaker.Add(originator.NewMomento())

	mem, err := careTaker.Momento(0)
	if err != nil {
		t.Fatal(err)
	}

	if mem.state.Description != "Idle" {
		t.Error("Unexpected state")
	}

	mem, err = careTaker.Momento(-1)
	if err == nil {
		t.Fatal("An error is expected when asking for a negative number but no error was found")
	}
}

func TestOriginator_ExtractAndStoreState(t *testing.T) {
	originator := originator{state: State{"Idle"}}
	idleMomento := originator.NewMomento()

	originator.ExtractAndStoreState(idleMomento)
	if originator.state.Description != "Idle" {
		t.Error("Unexpected state found")
	}
}
