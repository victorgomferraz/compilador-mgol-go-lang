package analizador_lexico


var palavrasReservadas = []string{"inicio", "varinicio", "varfim", "escreva", "leia", "se", "entao", "fimse", "faca-ate", "fimfaca", "fim", "inteiro", "real", "real"}

var tabelaSimbolos map[string]Token

func inicializarTabelaSimbolos() {
	tabelaSimbolos = make(map[string]Token)
	for _, s := range palavrasReservadas {
		tabelaSimbolos[s] = Token{s, s, ""}
	}
}
