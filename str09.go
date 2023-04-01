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
	fmt.Print("Caractere a ser trocado: ")
	fmt.Scan(&c1)
	fmt.Print("Trocar pelo caractere: ")
	fmt.Scan(&c2)

	if strings.Contains(str, c1) {
		fmt.Println(strings.ReplaceAll(str, c1, c2))
	} else {
		fmt.Printf("A frase não contém o caractere '%s'", c1)
	}
}
