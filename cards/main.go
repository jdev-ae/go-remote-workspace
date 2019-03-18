package main

func main() {
	cards := newDeck()
	hand1, remaining1 := cards.deal(5)

	hand1.print()
	remaining1.print()

	cards.writeToFile("mydeck.txt")
	hand1.writeToFile("hand1.txt")
	remaining1.writeToFile("remaining1.txt")

}
