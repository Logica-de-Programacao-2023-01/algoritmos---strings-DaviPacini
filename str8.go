//Escreva um programa que receba uma string e inverta a ordem dos caracteres. Informe ao usuário o resultado.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	chars := strings.Split(input, "")

	reverseChars := make([]string, len(chars))
	for i := 0; i < len(chars); i++ {
		reverseChars[i] = chars[len(chars)-1-i]
	}

	result := strings.Join(reverseChars, "")

	fmt.Println("Resultado:", result)
}

