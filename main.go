package main

import (
	"fmt"

	"github.com/daszabo/testmod"
	testmodML "github.com/daszabo/testmod/v2"
)

func main() {
	fmt.Println(testmod.Hi("David"))
	g, err := testmodML.Hi("David", "pt")
	if err != nil {
		panic(err)
	}
	fmt.Println(g)
}
