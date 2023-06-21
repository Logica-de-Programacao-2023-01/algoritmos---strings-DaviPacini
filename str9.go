//Solicite ao usuário uma string e substitua todas as ocorrências de uma letra por outra informada pelo usuário.

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
