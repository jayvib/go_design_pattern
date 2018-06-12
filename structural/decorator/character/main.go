package main

import "fmt"

type Eater interface {
	Eat(string) string
}

func NewPerson(name string) Person {
	return Person{
		name: name,
	}
}

type Person struct {
	name string
}

func (p Person) Eat(f string) string {
	return fmt.Sprintf("%s is eating %s", p.name, f)
}


type EaterFunc func(string) string

func (e EaterFunc) Eat(food string) string {
	return e(food)
}

type Decorator func(Eater) Eater

func Decorate(e Eater, ds ...Decorator) Eater {
	decorated := e
	for _, d := range ds {
		decorated = d(decorated)
	}
	return decorated
}

// Decorator tools

func BGM(song string) Decorator {
	return func(e Eater) Eater {
		return EaterFunc(func(f string) string {
			eat := e.Eat(f)
			bgm := "Playing: " + song
			eat = fmt.Sprintf("|%s|%s|%s|", bgm, eat , bgm)
			return eat
		})
	}
}

func BigFan() Decorator {
	return func(e Eater) Eater {
		return EaterFunc(func(f string) string {
			eat := e.Eat(f)
			fan := "Fanning with big leaf fan"
			eat = fmt.Sprintf("|%s|%s|%s|", fan, eat, fan)
			return eat
		})
	}
}

func StoryTeller(story string) Decorator {
	return func(e Eater) Eater {
		return EaterFunc(func(f string) string {
			eat := e.Eat(f)
			story := "A beatiful story"
			eat = fmt.Sprintf("|%s|%s|%s|", story, eat, story)
			return eat
		})
	}
}

func main() {
	p := NewPerson("Jayson")
	var eater Eater
	eater = p

	eater = Decorate(eater, BGM("Heaven Knows"), BigFan(), StoryTeller("The Alchemist))"))
	fmt.Println(eater.Eat("Crab"))
}