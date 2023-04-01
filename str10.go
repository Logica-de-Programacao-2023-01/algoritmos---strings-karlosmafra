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
	s1 := scanner.Text()

	fmt.Print("Digite outra frase: ")
	scanner.Scan()
	s2 := scanner.Text()

	anag := true

	for i := range s1 {
		car := string(s1[i])
		if !strings.Contains(s2, car) && car != " " {
			anag = false
		}
	}

	for i := range s2 {
		car := string(s2[i])
		if !strings.Contains(s1, car) && car != " " {
			anag = false
		}
	}

	if s1 == s2 {
		fmt.Println("As frases são iguais.")
	} else if anag == true {
		fmt.Println("As frases são anagramas.")
	} else {
		fmt.Println("As frases não são anagramas.")
	}
}
