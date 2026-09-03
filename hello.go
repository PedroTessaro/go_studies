package main

import "fmt"

const helloPrefix = "Hello, "

func Hello(name, language string) string {
	if name == "" {
		return "Hello World"
	}
	if language == "Portuguese" {
		return "Olá " + name
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("Pedro", ""))
}
