package review

type Eater interface {
	Eat()
}

type Swimmer interface {
	Swim()
}

type Speaker interface {
	Speak()
}

type Trainer interface {
	Train()
}

type Shark struct {
	Swimmer
	Eater
}

type SwimmerAthlete struct {
	Swimmer
	Trainer
}

func DoSometingWithAthlete(t Trainer) {
	t.Train()
}


