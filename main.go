package main

import (
	"fmt"
	newcolor "gopackages/color"
	"gopackages/random"
	"gopackages/wordz"

	"github.com/fatih/color"

	"github.com/huandu/xstrings"

	"github.com/FaNaGeR1995/rebrain/utils"
)

func main() {

	//С помощью нашей новой библиотеки проверяем слайс на наличие определенного значения
	//Если его там находим, то сообщим это в stdout и закончим выполнение программы
	isExist := utils.Contains(wordz.Words, "Two")
	if isExist {
		fmt.Println("Slice Words contain finding value")
		return
	}

	newcolor.Greet()
	fmt.Println("Hello world")
	color.Red("Hello world again")

	fmt.Println(wordz.Hello)
	wordz.Words = []string{"Moscow", "New-York", "Amsterdam", "Barcelona", "Paris"}
	fmt.Println(wordz.Random())

	wordz.Prefix = ""
	fmt.Println(random.City())
	fmt.Println(random.Digit())

	fmt.Println(xstrings.Shuffle(random.City()))
	fmt.Println(xstrings.Shuffle(random.Digit()))

}
