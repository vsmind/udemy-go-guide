package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreetingE(eb)
	printGreetingS(sb)
	fmt.Println("*****After refactoring*****")
	printGreeting(eb)
	printGreeting(sb)
}

// we can have only receiver type if we are not going to use value inside of function
func (englishBot) getGreeting() string {
	return "Hello there"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

// similar functions can be exchanged
func printGreetingE(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreetingS(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
