//Solicite ao usuário uma string e informe se ela contém pelo menos um número.

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	Possui_Numero := false

	for _, char := range input {
		if unicode.IsDigit(char) {
			Possui_Numero = true
			break
		}
	}

	if Possui_Numero {
		fmt.Println("A string contém pelo menos um número.")
	} else {
		fmt.Println("A string não contém nenhum número.")
	}
}
