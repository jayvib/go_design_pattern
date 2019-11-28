package animal

type fishMovement interface {
	moveTail()
	moveMouth()
}

type fish struct {
	fishMovement
}

func (f fish) Move() {
	f.moveTail()
	f.moveMouth()
}
