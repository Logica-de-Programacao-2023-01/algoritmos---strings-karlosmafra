package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str := scanner.Text()

	var cont string

	for i := len(str) - 1; i >= 0; i-- {
		var car = string(str[i])
		cont += car
	}

	if strings.ToUpper(str) == strings.ToUpper(cont) {
		fmt.Println("A frase é um palíndromo")
	} else {
		fmt.Println("A frase não é um palíndromo")
	}
}
