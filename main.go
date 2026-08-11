package main

func main(){
	cards:=deck{"spades","clubs","diamonds","hearts"}
	cardsnum:=deck{"ace","2","3","4","5","6"}

	for _,sym:=range cards{
		for _,num:=range cardsnum{
			cards=append(cards,sym+" of "+num)
		}
	}
	cards.print()
	
}