package main

import "fmt"

const (
	helloEnglishPrefix    = "Hello, "
	helloPortuguesePrefix = "Olá, "
	helloSpanishPrefix    = "Holla, "
	helloFrenchPrefix     = "Bonjour, "

	french     = "French"
	portuguese = "Portuguese"
	spanish    = "Spanish"
)

func Hello(name, language string) string {
	if name == "" {
		return "Hello World"
	}

	prefix := greetingProfile(language)

	return prefix + name
}

func greetingProfile(language string) (prefix string) {
	prefix = helloEnglishPrefix
	switch language {
	case french:
		prefix = helloFrenchPrefix

	case portuguese:
		prefix = helloPortuguesePrefix

	case spanish:
		prefix = helloSpanishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Pedro", ""))
}
