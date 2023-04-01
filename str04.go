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
	s1 := scanner.Text()

	fmt.Print("Digite outra frase: ")
	scanner.Scan()
	s2 := scanner.Text()

	if s1 == s2 {
		fmt.Println("As strings são iguais.")
	} else {
		fmt.Println("As strings são diferentes.")
	}
}
