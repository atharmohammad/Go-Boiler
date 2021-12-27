package main

import "fmt"

type Voice interface {
	speakHello() string // provided function inside the interface will make each struct the function is a reciever of
	// of a type Voice so now we can use the function print with only one implementation
	// as all language had the same implementation for that function , so now we dont have to define
	// function
}

type EnglishVoice struct{}
type HindiVoice struct{}
type UrduVoice struct{}

func main() {
	eng := EnglishVoice{}
	hi := HindiVoice{}
	ur := UrduVoice{}

	// eng.speakHello() //Since all 3 lang has different implementation for this function we have 3 functions
	// hi.speakHello()
	// ur.speakHello()

	print(eng)
	print(hi)
	print(ur)
}

//Assume speakHello functions have different implementation for different language
func (e EnglishVoice) speakHello() string {
	return "Hello!"
}

func (h HindiVoice) speakHello() string {
	return "Namaste!"
}

func (u UrduVoice) speakHello() string {
	return "Salam!"
}

func print(v Voice) {
	fmt.Println(v.speakHello())
}
