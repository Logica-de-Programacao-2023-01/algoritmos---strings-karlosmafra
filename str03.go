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

	var c1, c2 string

	fmt.Print("Caractere para ser substituído: ")
	fmt.Scan(&c1)
	fmt.Print("Substituir por: ")
	fmt.Scan(&c2)

	if !strings.Contains(str, c1) {
		fmt.Println("A frase não contem o caractere.")
	} else {
		fmt.Println(strings.ReplaceAll(str, c1, c2))
	}
}
