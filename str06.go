package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str := scanner.Text()

	p := 1

	if len(str) == 0 {
		p = 0
	} else {
		for i := 1; i < len(str); i++ {
			var car = string(str[i])
			var ant = string(str[i-1])
			if car == " " && ant != " " {
				p++
			}
		}
	}

	if p == 1 {
		fmt.Println("A frase tem", p, "palavra.")
	} else {
		fmt.Println("A frase tem", p, "palavras.")
	}
}
