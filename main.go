package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var caminho = "test/mgol_files/simple2.alg"
	args := os.Args

	if len(args) > 1 && args[1] != "" {
		caminho = args[1]
	}

	fmt.Printf("Lendo arquivo: %s\n", caminho)
	conteudo, _ := ioutil.ReadFile(caminho)
	tokens, errosLexicos := ExtrairTokens(string(conteudo))
	erroSemanticoOuSintatico := Parse(tokens)
	/*
		for _, token := range tokens {
			fmt.Printf("Classe: %s, Lexema: %s, Tipo: %s\n", token.classe, token.lexema, token.tipo)
		}
	*/
	for _, erro := range errosLexicos {
		HandleError(erro.erro, erro.linha, erro.coluna)
	}

	if len(errosLexicos) == 0 && !erroSemanticoOuSintatico {
		gravarCodigoObjeto(caminho)
	}

}
