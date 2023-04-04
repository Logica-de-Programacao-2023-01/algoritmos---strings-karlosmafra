package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma sequência numérica (sem espaço ou vírgula): ")
	scanner.Scan()
	str := scanner.Text()

	dcr := true

	ints, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		fmt.Println("A string não é uma sequência de números inteiros.")
	} else {
		for i := 0; i < len(str)-1; i++ {
			if str[i+1] > str[i] {
				dcr = false
			}
		}
		if dcr == true {
			fmt.Println("A string é uma sequência decrescente.")
			fmt.Println(ints)
		} else {
			fmt.Println("A string não é uma sequência decrescente.")
			fmt.Println(ints)
		}
	}
}
