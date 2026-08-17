package main



func main(){
	//cards:= newdeck()
	// hand,remainingdeck:=deal(cards,5)
	// hand.print()
	// fmt.Println("remaining cards contain:")
	// remainingdeck.print()
    //cards.saveToFile("my_cards")
	cards:=NewFileFromDeck("my_cards")
	cards.shuffle()
	cards.print()

}