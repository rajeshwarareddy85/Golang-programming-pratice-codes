package main

import "fmt"

type deck []string

func newdeck() deck {
    cards:=deck{}
	cardsuits:=[]string{"spades","clubs","diamonds","hearts"}
	cardsnum:=[]string{"ace","2","3","4","5","6"}

	for _,sym:=range cardsuits{
		for _,num:=range cardsnum{
			cards=append(cards,sym+" of "+num)
		}
	}
	return cards
}

func (d deck) print(){
	for i,card:=range d{
		fmt.Println(i,card)
	}
}

func deal(d deck,handsize int)(deck, deck){
	return d[:handsize],d[handsize:]
}