package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(myGlass())
	fmt.Println(myGo())
	fmt.Println(myHello())
	fmt.Println(myOpt())
}

func myGlass() string {
	return quote.Glass()
}

func myGo() string {
	return quote.Go()
}

func myHello() string {
	return quote.Hello()
}

func myOpt() string {
	return quote.Opt()
}
