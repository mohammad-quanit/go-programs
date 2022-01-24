package main

import (
	"fmt"

	"github.com/mohammad-quanit/go-programs/packer/numbers"
	"github.com/mohammad-quanit/go-programs/packer/strings"
	"github.com/mohammad-quanit/go-programs/packer/strings/greeting"
)

func main() {
	fmt.Println(numbers.IsPrime(4))
	fmt.Println(strings.Reverse("quanit"))
	fmt.Println(greeting.Welcome)
	fmt.Println(greeting.MorningText)
	fmt.Println(greeting.EveningText)
}
