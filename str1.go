//Escreva um programa que receba uma string e converta todas as letras minúsculas para maiúsculas. Informe ao usuário
//o resultado.

package main

import "fmt"

//x - 32

func main() {
	var word string
	fmt.Println("Informe uma string:")
	fmt.Scan(&word)
	for i := 0; i < len(word); i++ {
		fmt.Print(string(word[i] - 32))
	}
}
