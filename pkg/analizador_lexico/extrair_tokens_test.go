package analizador_lexico

import (
	"strings"
	"testing"
)

func TestExtrairTokens(t *testing.T) {
	inicializarTabelaSimbolos()
	tokens, erros := ExtrairTokens("123 se aux @ x 1\nfimse")

	if len(tokens) != 6 {
		t.Errorf("Número de tokens diferente do esperado, encontrados : %d, esperados %d", len(tokens), 6)
	} else {
		checarToken(t, tokens[0], "Num", "123")
		checarToken(t, tokens[1], "se", "se")
		checarToken(t, tokens[2], "id", "aux")
		checarToken(t, tokens[3], "ERRO4", "@")
		checarToken(t, tokens[4], "fimse", "fimse")
		checarToken(t, tokens[5], "EOF", "EOF")
	}

	if len(erros) != 1 {
		t.Errorf("Número de erros diferente do esperado, encontrados : %d, esperados %d", len(erros), 1)
	} else {
		if erros[0].coluna != 12 || erros[0].linha != 1 || !strings.HasPrefix(erros[0].erro, "ERRO4") {
			t.Errorf("Erro não encontrado, esperado ")
		}
	}

}

func checarToken(t *testing.T, token Token, classe, lexema string) {
	if token.lexema != lexema {
		t.Errorf("Lexema diferente do experado, obtido : %s, esperado %s", token.lexema, lexema)
	}

	if token.classe != classe {
		t.Errorf("Classe diferente do experado, obtido : %s, esperado %s", token.lexema, lexema)
	}
}
