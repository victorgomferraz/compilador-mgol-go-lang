package main


import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestIdentificaPalavraReservada(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	inicializarTabelaSimbolos()

	var palavraReservada = palavrasReservadas[rand.Intn(len(palavrasReservadas)-1)]
	lexemaNaoEncontrado, _ := identificaPalavraReservada(palavraReservada+"NAOENCONTRAR", 0)
	lexema1, _ := identificaPalavraReservada(palavraReservada, 0)
	lexema2, _ := identificaPalavraReservada("PREFIXO "+palavraReservada, 7)
	lexema3, _ := identificaPalavraReservada("PREFIXO "+palavraReservada+" SUFIXO", 7)

	if lexema1 == lexema2 && lexema2 == lexema3 && lexema1 != palavraReservada {
		t.Errorf("Identificação diferente do esperado, retornado: %s, esperado: %s.", lexema1, palavraReservada)
	}
	if lexemaNaoEncontrado != "" {
		t.Errorf("Não deveria encontrar Lexema, lexema :  %s", lexemaNaoEncontrado)
	}

	if _, ok := tabelaSimbolos[lexema1]; !ok {
		t.Errorf("Palavra reservada não esta na tabela de simbolos")
	}
}

func TestIdentificaTokenSemPalavraReservada(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	inicializarTabelaSimbolos()

	verificarTokenSemPalavraReservada(t, "Num", "1", 0)
	verificarTokenSemPalavraReservada(t, "Num", "123", 0)
	verificarTokenSemPalavraReservada(t, "Num", "123.", 0)
	verificarTokenSemPalavraReservada(t, "Num", "12.2", 0)
	verificarTokenSemPalavraReservada(t, "Num", "12E+11", 0)
	verificarTokenSemPalavraReservada(t, "Num", "12.23E+11", 0)
	verificarTokenSemPalavraReservada(t, "Num", "12.23e+11", 0)
	verificarTokenSemPalavraReservada(t, "ERRO1", "12.23e", 0)
	verificarTokenSemPalavraReservada(t, "ERRO1", "12.23e+", 0)
	verificarTokenSemPalavraReservada(t, "Literal", "\"asdasd\"", 0)
	verificarTokenSemPalavraReservada(t, "Literal", "\"asdasd123\"", 0)
	verificarTokenSemPalavraReservada(t, "Literal", "\";\"", 0)
	verificarTokenSemPalavraReservada(t, "Literal", "\"123\"", 0)
	verificarTokenSemPalavraReservada(t, "ERRO2", "\"123", 0)
	verificarTokenSemPalavraReservada(t, "id", "asd", 0)
	verificarTokenSemPalavraReservada(t, "id", "D1", 0)
	verificarTokenSemPalavraReservada(t, "Comentário", "{asdas}", 0)
	verificarTokenSemPalavraReservada(t, "Comentário", "{123}", 0)
	verificarTokenSemPalavraReservada(t, "Comentário", "{{{}", 0)
	verificarTokenSemPalavraReservada(t, "Comentário", "{;}", 0)
	verificarTokenSemPalavraReservada(t, "ERRO3", "{;", 0)
	verificarTokenSemPalavraReservada(t, "OPR", "<", 0)
	verificarTokenSemPalavraReservada(t, "OPR", ">", 0)
	verificarTokenSemPalavraReservada(t, "OPR", "=", 0)
	verificarTokenSemPalavraReservada(t, "OPR", "<=", 0)
	verificarTokenSemPalavraReservada(t, "OPR", ">=", 0)
	verificarTokenSemPalavraReservada(t, "RCB", "<-", 0)
	verificarTokenSemPalavraReservada(t, "OPM", "+", 0)
	verificarTokenSemPalavraReservada(t, "OPM", "-", 0)
	verificarTokenSemPalavraReservada(t, "OPM", "*", 0)
	verificarTokenSemPalavraReservada(t, "OPM", "/", 0)
	verificarTokenSemPalavraReservada(t, "AB_P", "(", 0)
	verificarTokenSemPalavraReservada(t, "FC_P", ")", 0)
	verificarTokenSemPalavraReservada(t, "PT_V", ";", 0)
	verificarTokenSemPalavraReservada(t, "Vir", ",", 0)
	verificarTokenSemPalavraReservada(t, "ERRO4", "!", 0)
}

func TestScanner(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	inicializarTabelaSimbolos()

	for _, palavraReservada := range palavrasReservadas {
		verificarTokenScanner(t, palavraReservada, palavraReservada, 0)
	}
	verificarTokenScanner(t, "Num", "1", 0)
	verificarTokenScanner(t, "Num", "1", 0)
	verificarTokenScanner(t, "Num", "123", 0)
	verificarTokenScanner(t, "Num", "123.", 0)
	verificarTokenScanner(t, "Num", "12.2", 0)
	verificarTokenScanner(t, "Num", "12E+11", 0)
	verificarTokenScanner(t, "Num", "12.23E+11", 0)
	verificarTokenScanner(t, "Num", "12.23e+11", 0)
	verificarTokenScanner(t, "ERRO1", "12.23e", 0)
	verificarTokenScanner(t, "ERRO1", "12.23e+", 0)
	verificarTokenScanner(t, "Literal", "\"asdasd\"", 0)
	verificarTokenScanner(t, "Literal", "\"asdasd123\"", 0)
	verificarTokenScanner(t, "Literal", "\";\"", 0)
	verificarTokenScanner(t, "Literal", "\"123\"", 0)
	verificarTokenScanner(t, "ERRO2", "\"123", 0)
	verificarTokenScanner(t, "id", "asd", 0)
	verificarTokenScanner(t, "id", "D1", 0)
	verificarTokenScanner(t, "Comentário", "{asdas}", 0)
	verificarTokenScanner(t, "Comentário", "{123}", 0)
	verificarTokenScanner(t, "Comentário", "{{{}", 0)
	verificarTokenScanner(t, "Comentário", "{;}", 0)
	verificarTokenScanner(t, "ERRO3", "{;", 0)
	verificarTokenScanner(t, "OPR", "<", 0)
	verificarTokenScanner(t, "OPR", ">", 0)
	verificarTokenScanner(t, "OPR", "=", 0)
	verificarTokenScanner(t, "OPR", "<=", 0)
	verificarTokenScanner(t, "OPR", ">=", 0)
	verificarTokenScanner(t, "RCB", "<-", 0)
	verificarTokenScanner(t, "OPM", "+", 0)
	verificarTokenScanner(t, "OPM", "-", 0)
	verificarTokenScanner(t, "OPM", "*", 0)
	verificarTokenScanner(t, "OPM", "/", 0)
	verificarTokenScanner(t, "AB_P", "(", 0)
	verificarTokenScanner(t, "FC_P", ")", 0)
	verificarTokenScanner(t, "PT_V", ";", 0)
	verificarTokenScanner(t, "Vir", ",", 0)
	verificarTokenScanner(t, "ERRO4", "!", 0)
	verificarTokenScanner(t, "EOF", "", 0)
}

func verificarTokenSemPalavraReservada(t *testing.T, tokenString, lexema string, indexInicio int) {
	token, indexFim := identificaTokenSemPalavraReservada(lexema, indexInicio, true)
	verificarToken(t, tokenString, lexema, indexInicio, token, indexFim)

}

func verificarTokenScanner(t *testing.T, tokenString, lexema string, indexInicio int) {
	token, indexFim := Scanner(lexema, indexInicio, true)
	verificarToken(t, tokenString, lexema, indexInicio, token, indexFim)
}

func verificarToken(t *testing.T, tokenString, lexema string, indexInicio int, token Token, indexFim int) {
	if token.classe != "EOF" {
		if indexInicio+len(token.lexema) != indexFim {
			fmt.Println(token)
			t.Errorf("Posição final de leitura esta incorreta, esperava : %d , foi obtido : %d", indexInicio+len(token.lexema), indexFim)
		}

		if token.classe != tokenString || token.lexema != lexema {
			t.Errorf("Token incorreto, lexema e classe esperados : ( %s ,%s ) , obtidos : ( %s ,%s )", lexema, tokenString, token.lexema, token.classe)
		}
	} else {
		if indexInicio != indexFim {
			t.Errorf("Posição final de leitura esta incorreta, esperava : %d , foi obtido : %d", indexInicio, indexFim)
		}

		if token.classe != "EOF" || token.lexema != "EOF" {
			t.Errorf("Token incorreto, lexema e classe esperados : ( %s ,%s ) , obtidos : ( %s ,%s )", "EOF", "EOF", token.lexema, token.classe)
		}
	}
}