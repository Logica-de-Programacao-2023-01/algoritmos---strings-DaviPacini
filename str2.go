//Escreva um programa que receba uma string e remova todos os espaços em branco. Informe ao usuário o resultado

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

	result := strings.ReplaceAll(input, " ", "")

	fmt.Println("Resultado:", result)
}
