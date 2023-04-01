package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite um número: ")
	scanner.Scan()
	str := scanner.Text()

	float, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Print("A string não é um valor em ponto flutuante.")
	} else {
		fmt.Println("A string é o valor em ponto flutuante", float)
	}
}
