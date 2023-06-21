//Escreva um programa que receba uma string e conte quantas palavras ela contém. Informe ao usuário o resultado.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Digite uma string: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	words := strings.Fields(input)
	wordCount := len(words)

	fmt.Printf("A string contém %d palavra(s).\n", wordCount)
}
