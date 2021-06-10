package main

import (
	"strings"
)

func ExtrairTokens(entrada string) ([]*Token, []Error) {
	inicializarTabelaSimbolos()
	var tokens []*Token
	var errors []Error

	linhas := strings.Split(entrada, "\n")

	for linhaIndex, linhaString := range linhas {
		var coluna = 0
		for {
			token, index := Scanner(linhaString, coluna, linhaIndex == len(linhas)-1)
			if token != nil {
				coluna = index
				token.col = coluna
				token.lin = linhaIndex + 1
			}

			if token != nil {
				tokens = append(tokens, token)
				if strings.HasPrefix(token.classe, "ERRO") {
					errors = append(errors, Error{token.classe, linhaIndex + 1, coluna})
					break
				}
			}

			if (token != nil && token.classe == "EOF") || (linhaIndex != len(linhas)-1 && index == len(linhaString)) {
				break
			}
		}
	}

	return tokens, errors
}
