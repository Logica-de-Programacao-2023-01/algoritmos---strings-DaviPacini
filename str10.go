//Escreva um programa que receba duas strings e verifique se elas são anagramas. Informe ao usuário se são ou não.

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var str1, str2 string

	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)

	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	// Remove espaços em branco e converte para letras minúsculas
	str1 = normalizeString(str1)
	str2 = normalizeString(str2)

	// Verifica se as strings são anagramas
	if areAnagrams(str1, str2) {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}

// Função para remover espaços em branco e converter para letras minúsculas
func normalizeString(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ToLower(str)
	return str
}

// Função para verificar se duas strings são anagramas
func areAnagrams(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	// Converte as strings em slices de caracteres
	chars1 := strings.Split(str1, "")
	chars2 := strings.Split(str2, "")

	// Ordena as slices de caracteres
	sort.Strings(chars1)
	sort.Strings(chars2)

	// Verifica se as slices ordenadas são iguais
	for i := 0; i < len(chars1); i++ {
		if chars1[i] != chars2[i] {
			return false
		}
	}

	return true
}
