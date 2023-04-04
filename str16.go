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
	S := scanner.Text()

	fmt.Print("Digite outra frase: ")
	scanner.Scan()
	s := scanner.Text()

	if strings.Contains(S, s) {
		fmt.Println("A segunda frase está contida na primeira.")
	} else {
		fmt.Println("A segunda frase não está contida na primeira.")
	}
}
