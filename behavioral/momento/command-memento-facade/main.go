package main

import "fmt"

func main() {
	m := MementoFacade{}
	m.SaveSetting(Volume(4))
	m.SaveSetting(Mute(false))
	assertAndPrint(m.RestoreSetting())
	assertAndPrint(m.RestoreSetting())
}

func assertAndPrint(c Command) {
	switch cast := c.(type) {
	case Volume:
		fmt.Printf("Volume:\t%d\n", cast)
	case Mute:
		fmt.Printf("Mute:\t%t\n", cast)
	}
}
