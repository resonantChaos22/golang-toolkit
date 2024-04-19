package main

import (
	"fmt"

	"github.com/resonantChaos22/toolkit"
)

func main() {
	var tools toolkit.Tools

	s := tools.RandomString(100)

	fmt.Println("Random String: ", s)
}
