package main

import (
	"testing"
)

func TestStack_IsEmpty(t *testing.T) {
	var stackEmpty, stack Stack

	if !stackEmpty.IsEmpty() {
		t.Errorf("Pilha devia estar vazia")
	}

	stack.Push("a")
	if stack.IsEmpty() {
		t.Errorf("Pilha não devia estar vazia")
	}

	if len(stack) != 1 {
		t.Errorf("Pilha devia conter %d porem contem %d", 1, len(stack))
	}

	if len(stackEmpty) != 0 {
		t.Errorf("Pilha devia conter %d porem contem %d", 0, len(stack))
	}

}

func TestStack_IsPush(t *testing.T) {
	var stack Stack
	stack.Push("a")
	if stack.IsEmpty() {
		t.Errorf("Pilha não devia estar vazia")
	}

}

func TestStack_IsPop(t *testing.T) {
	var stack Stack
	expect := [5]string{"a", "b", "c", "d", "e"}
	if !stack.IsEmpty() {
		t.Errorf("Pilha devia estar vazia")
	}

	for i, v := range expect {
		stack.Push(v)
		if i+1 != len(stack) {
			t.Errorf("Pilha devia conter %d porem contem %d", i+1, len(stack))
		}
		if stack.IsEmpty() {
			t.Errorf("Pilha não devia estar vazia")
		}
	}

	legthStack := len(stack)
	for i := 0; i < legthStack; i++ {
		value, isOk := stack.Pop()
		if !isOk {
			t.Errorf("Não era esperado erro, posicao loop %d, tamanho stack %d, valor %s", i, len(stack), value)
		}
		if value != expect[len(expect)-1-i] {
			t.Errorf("Valor esperado do POP é %s porem contem %s", expect[len(expect)-1-i], value)
		}
	}

	if !stack.IsEmpty() {
		t.Errorf("Pilha devia estar vazia, ela possui %d elementos", len(stack))
	}

	value, isOk := stack.Pop()

	if value != "" || isOk {
		t.Errorf("Era esperado não obter valores, porem foi obtido: %s", value)
	}

}
