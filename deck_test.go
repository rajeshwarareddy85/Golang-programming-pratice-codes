package main

import (
	"testing"
	"os"
)


func TestNewDeck(t *testing.T){
	deck:=NewDeck()

	if len(deck)!=24{
		t.Errorf("expected deck is 24 but you got %v",len(deck))
	}

	if deck[0]!="spades of ace"{
		t.Errorf("expected spades of ace but got %v",deck[0])
	}

	if deck[len(deck)-1]!="hearts of 6"{
		t.Errorf("expected hearts of 6 but you got %v",deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewFileFromDeck(t *testing.T){
	os.Remove("_decktesting")
	d:=NewDeck()
	d.SaveToFile("_decktesting")
	load:=NewFileFromDeck("_decktesting")
	if len(load)!=24{
		t.Errorf("expected deck 24 length but you got %v",len(load))
	}
}