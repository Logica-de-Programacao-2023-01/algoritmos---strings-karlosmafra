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

	var frase string

	for i := len(str) - 1; i >= 0; i-- {
		car := string(str[i])
		frase += car
	}

	fmt.Println(frase)
}
