//Escreva um programa que receba uma string e verifique se ela é um número válido em ponto flutuante.
//Informe ao usuário se é ou não.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var valor string

	fmt.Print("Digite um número em ponto flutuante: ")
	fmt.Scanln(&valor)

	_, err := strconv.ParseFloat(valor, 64)
	if err == nil {
		fmt.Println("É um número válido em ponto flutuante.")
	} else {
		fmt.Println("Não é um número válido em ponto flutuante.")
	}
}
