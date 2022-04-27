package main

import (
	"fmt"
	"gopackages/random"
	"gopackages/wordz"

	"github.com/huandu/xstrings"
)

func main() {
	wordz.Prefix = ""
	wordz.Words = []string{"London", "Moscow", "Rome", "Paris", "Luxembourg", "Berlin", "New-York"}
	fmt.Println("City:", random.City())
	fmt.Println("City (with using Shuffle func):", xstrings.Shuffle(random.City()))
	wordz.Words = []string{"Six", "Seven", "Eight", "Nine", "Ten"}
	fmt.Println("Digit:", random.Digit())
}

