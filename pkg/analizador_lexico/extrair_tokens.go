package analizador_lexico

import (
	"strings"
)

func ExtrairTokens(entrada string) ([]Token, []Error) {
	var tokens []Token
	var errors []Error

	linhas := strings.Split(entrada, "\n")

	for linhaIndex, linhaString := range linhas {
		var coluna = 0
		for {
			token, index := Scanner(linhaString, coluna)
			coluna = index
			tokens = append(tokens, token)
			if strings.HasPrefix(token.classe, "ERRO") {
				errors = append(errors, Error{token.classe, linhaIndex + 1, coluna})
				break
			}
			if token.classe == "EOF" {
				break
			}
		}
	}

	return tokens, errors
}
