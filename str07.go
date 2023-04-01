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

	if strings.Contains(str, "1") || strings.Contains(str, "2") || strings.Contains(str, "3") || strings.Contains(str, "4") || strings.Contains(str, "5") || strings.Contains(str, "6") || strings.Contains(str, "7") || strings.Contains(str, "8") || strings.Contains(str, "9") || strings.Contains(str, "0") {
		fmt.Println("A string contém pelo menos um número.")
	} else {
		fmt.Println("A string não contém nenhum número.")
	}
}
