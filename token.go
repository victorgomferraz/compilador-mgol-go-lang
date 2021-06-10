package main

type Token struct {
	lexema, classe, tipo string
	lin, col int
}

func (t *Token) set(new *Token)  {
	t.lexema = new.lexema
	t.classe = new.classe
	t.tipo = new.tipo
	t.lin = new.lin
	t.col = new.col
}
