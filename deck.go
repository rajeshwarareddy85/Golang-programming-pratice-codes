package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"math/rand"
	"time"
)

type deck []string

func NewDeck() deck {
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

func (d deck) toString() string {
    return strings.Join([]string(d),",")
}

func (d deck) SaveToFile(filename string) error{
	return ioutil.WriteFile(filename,[]byte(d.toString()),0666)
}

func NewFileFromDeck(filename string) deck{
	ds,err:=os.ReadFile(filename)
	if err != nil{
		fmt.Print("Error :",err)
		os.Exit(1)
	}
	s:=strings.Split(string(ds),",")
	return deck(s)
}

func (d deck) shuffle(){
	source:=rand.NewSource(time.Now().UnixNano())
	r:=rand.New(source)
	for i:=range d {
		index:=r.Intn(len(d)-1)
		d[i],d[index]=d[index],d[i]
	}
}