package main

var palavrasReservadas = []string{"inicio", "varinicio", "varfim", "escreva", "leia", "se", "entao", "fimse", "facaAte", "fimFaca", "fim", "inteiro", "real", "lit", "int", "literal"}

var tabelaSimbolos map[string]*Token

func inicializarTabelaSimbolos() {
	tabelaSimbolos = make(map[string]*Token)
	for _, s := range palavrasReservadas {
		t:= Token{s, s, "",0, 0}
		tabelaSimbolos[s] = &t
	}

	t:= Token{"inteiro", "inteiro", "inteiro",0, 0}
	tabelaSimbolos["inteiro"] = &t

	t2:= Token{"lit", "lit", "literal",0, 0}
	tabelaSimbolos["lit"] = &t2

	t3:= Token{"real", "real", "real",0, 0}
	tabelaSimbolos["real"] = &t3
}