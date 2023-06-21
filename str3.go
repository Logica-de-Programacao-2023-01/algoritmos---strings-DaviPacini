//Escreva um programa que receba uma string e substitua todas as ocorrências desse caractere na string por outro
//caractere especificado pelo usuário. Informe ao usuário o resultado.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		palavra    string
		letra      string
		substituir string
	)
	fmt.Println("Informe uma palavra")
	fmt.Scan(&palavra)
	fmt.Println("Agora informe qual letra você quer que seja substituída:")
	fmt.Scan(&letra)
	fmt.Println("Por qual letra vc quer q seja substituído?")
	fmt.Scan(&substituir)
	resultado := strings.ReplaceAll(palavra, letra, substituir)
	fmt.Print("Resultado: ", resultado)
}
