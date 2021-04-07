package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	conteudo, _ := ioutil.ReadFile("test/mgol_files/simple.alg")
	tokens, erros := ExtrairTokens(string(conteudo))

	for _, token := range tokens {
		fmt.Printf("Classe: %s, Lexema: %s, Tipo: %s\n", token.classe, token.lexema, token.tipo)
	}
	for _, erro := range erros {
		HandleError(erro.erro, erro.linha, erro.coluna)
	}

}
