package main

import (
	"fmt"
	"go-design-pattern/creational/factory/abstractFactory"
)

func main() {
	adidasFactory := abstractFactory.GetSportsFactory("adidas")

	adidasShoe := adidasFactory.MakeShoe()

	printShoeDetails(adidasShoe)

	adidasShort := adidasFactory.MakeShort()

	printShoeDetails(adidasShort)

	nikeFactory := abstractFactory.GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()

	printShortDetails(nikeShoe)

	nikeShort := nikeFactory.MakeShort()

	printShoeDetails(nikeShort)
}

func printShoeDetails(s abstractFactory.IShoe) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}

func printShortDetails(s abstractFactory.IShort) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}
