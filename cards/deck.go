package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() deck {
	suites := []string{"♣", "♠", "♦", "♥"}
	values := []string{"A", "2", "3", "J", "Q", "K"}

	cards := deck{}

	for _, suit := range suites {
		for _, value := range values {
			cards = append(cards, suit+" "+value)
		}
	}

	return cards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, "\t"+card)
	}
	fmt.Println("-----------")
}

// regular function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// receiver function
func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) writeToFile(file string) error {
	return ioutil.WriteFile(file, []byte(d.toString()), 0666)
}
