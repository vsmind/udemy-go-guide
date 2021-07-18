package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of deck which is a slice of string
type deck []string

/*
Receiver *d* of type deck, any variable of type deck has access to print() method
d - instance of deck variable
by convention receiver should be mentioned as 1 or 2 literals abbreviation
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	creatureCards := []string{"Bonecrusher Giant", "Brazen Borrower", "Galazeth Prismari", "Goldspan Dragon"}
	//creatureCount := []int{4,4,4,4}

	// we can use _ instead of variable we are not using
	// we dont need indexes in this slices thi mean we can replace it with _
	for _, creature := range creatureCards {
		cards = append(cards, creature)
	}

	return cards
}

/*
we are going to return to separate values of type deck
 [0:2] - return values from slice from 0 to 2 exclude
 [:2]  -
 [2:]  -
*/

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		println("Error: ", err.Error())
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	// Random generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
