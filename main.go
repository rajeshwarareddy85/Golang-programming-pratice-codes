package main
import "fmt"
func main(){
	cards:= newdeck()
	hand,remainingdeck:=deal(cards,5)
	hand.print()
	fmt.Println("remaining cards contain:")
	remainingdeck.print()
	
}