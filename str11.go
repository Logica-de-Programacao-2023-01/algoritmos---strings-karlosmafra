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

	str = strings.ReplaceAll(str, "a", "")
	str = strings.ReplaceAll(str, "e", "")
	str = strings.ReplaceAll(str, "i", "")
	str = strings.ReplaceAll(str, "o", "")
	str = strings.ReplaceAll(str, "u", "")
	str = strings.ReplaceAll(str, "A", "")
	str = strings.ReplaceAll(str, "E", "")
	str = strings.ReplaceAll(str, "I", "")
	str = strings.ReplaceAll(str, "O", "")
	str = strings.ReplaceAll(str, "U", "")
	str = strings.ReplaceAll(str, "á", "")
	str = strings.ReplaceAll(str, "é", "")
	str = strings.ReplaceAll(str, "í", "")
	str = strings.ReplaceAll(str, "ó", "")
	str = strings.ReplaceAll(str, "ú", "")
	str = strings.ReplaceAll(str, "Á", "")
	str = strings.ReplaceAll(str, "É", "")
	str = strings.ReplaceAll(str, "Í", "")
	str = strings.ReplaceAll(str, "Ó", "")
	str = strings.ReplaceAll(str, "Ú", "")
	str = strings.ReplaceAll(str, "â", "")
	str = strings.ReplaceAll(str, "ê", "")
	str = strings.ReplaceAll(str, "ô", "")
	str = strings.ReplaceAll(str, "Â", "")
	str = strings.ReplaceAll(str, "Ê", "")
	str = strings.ReplaceAll(str, "Ô", "")
	str = strings.ReplaceAll(str, "ã", "")
	str = strings.ReplaceAll(str, "Ã", "")

	fmt.Println(str)
}
