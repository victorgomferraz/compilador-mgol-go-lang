package main

import (
	"testing"
)

func TestCsvToMap(t *testing.T) {
	rulesMap := CsvToMap("./tabelas/regras.csv")
	assertEquals(t, "V", rulesMap["3"]["P1"])
	assertEquals(t, "varinicio LV", rulesMap["3"]["P2"])
	assertEquals(t, "2", rulesMap["3"]["TamanhoB"])
}

func assertEquals(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("CSV diferente do experado, obtido : %s, esperado %s", actual, expected)
	}
}
