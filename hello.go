package main

import "fmt"

const spanish = "Spanish"
const lithuanian = "Lithuanian"

const helloEnglish = "Hello"
const helloSpanish = "Hola"
const helloLithuanian = "Labas"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	greeting := helloEnglish

	switch language {
	case spanish:
		greeting = helloSpanish
	case lithuanian:
		greeting = helloLithuanian
	}

	return fmt.Sprintf("%s, %s!", greeting, name)
}

func main() {
	fmt.Println(Hello("World", ""))
}
