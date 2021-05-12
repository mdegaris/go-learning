package greeting

import "fmt"

func English() {
	hello()
}

func French() {
	bonjour()
}

func hello() {
	fmt.Println("Hello")
}

func bonjour() {
	fmt.Println("Bonjour")
}
