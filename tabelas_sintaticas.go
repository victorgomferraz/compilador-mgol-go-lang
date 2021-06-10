package main

var tabelaSintaticaTerminais map[string]map[string]string
var tabelaSintaticaNaoTerminais map[string]map[string]string
var tabelaSintaticaRegras map[string]map[string]string

func inicializarTabelaSintaticas() {
	tabelaSintaticaTerminais = CsvToMap("./tabelas/terminais.csv")
	tabelaSintaticaNaoTerminais = CsvToMap("./tabelas/nao_terminais.csv")
	tabelaSintaticaRegras = CsvToMap("./tabelas/regras.csv")
}
