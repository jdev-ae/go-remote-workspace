package main

import "fmt"

func main() {
	fmt.Println("Example for loop over Slice")

	cards := []string{
		"King of Spades",
		"Ace of Diamonds",
		"Eight of Clubs",
	}

	fmt.Println(cards)
	fmt.Println("All Cards")

	for index, card := range cards {
		fmt.Println(index, card)
	}

}
