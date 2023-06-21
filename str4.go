//Solicite ao usuário duas strings e informe se elas são iguais ou diferentes.

package main

import "fmt"

func main() {
	var (
		str1 string
		str2 string
	)
	fmt.Println("Informe a primeira string:")
	fmt.Scan(&str1)
	fmt.Println("Informe a segunda string")
	fmt.Scan(&str2)
	if str1 == str2 {
		fmt.Println("As strings são iguais")
	} else {
		fmt.Println("As strings são diferentes")
	}
}
