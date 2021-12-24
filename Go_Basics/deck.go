package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	card_types := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	card_num := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	for _, card_t := range card_types {
		for _, card_n := range card_num {
			cards = append(cards, card_n+" of "+card_t)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) save(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func read(filename string) deck {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	str := string(data)
	arr := strings.Split(str, ",")
	cards := deck(arr)
	return cards
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newPos := r.Intn(len(d) - 1)
		d[index], d[newPos] = d[newPos], d[index]
	}
}
