package hello_world

import "fmt"

const spanish = "Spanish"
const lithuanian = "Lithuanian"

const helloEnglish = "Hello"
const helloSpanish = "Hola"
const helloLithuanian = "Labas"

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloSpanish
	case lithuanian:
		prefix = helloLithuanian
	default:
		prefix = helloEnglish
	}

	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s, %s!", getGreetingPrefix(language), name)
}

func main() {
	fmt.Println(Hello("World", ""))
}
