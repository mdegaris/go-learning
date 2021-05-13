package greeting

import "fmt"

type Language int

const (
	ENGLISH = iota
	FRENCH  = iota
	SPANISH = iota
)

func Greet(lang Language) {
	if lang == ENGLISH {
		hello()
	} else if lang == FRENCH {
		bonjour()
	} else if lang == SPANISH {
		hola()
	} else {
		unknown()
	}
}
func hello() {
	fmt.Println("Hello")
}

func bonjour() {
	fmt.Println("Bonjour")
}

func hola() {
	fmt.Println("Hola")
}

func unknown() {
	fmt.Println("Unknown language")
}
