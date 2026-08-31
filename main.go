package main

import "fmt"

// no representa um elemento da lista encadeada
type no struct {
	valor   int
	proximo *no
}

type lista struct {
	inicio  *no
	fim     *no
	tamanho int
}

func (l *lista) adicionarInicio(valor int) {
	novo := &no{valor: valor, proximo: l.inicio}
	l.inicio = novo

	if l.fim == nil {
		l.fim = novo
	}
	l.tamanho++
}

func (l *lista) adicionarFim(valor int) {
	novo := &no{valor: valor}

	if l.fim == nil {
		// lista vazia: início e fim apontam para o novo nó
		l.inicio = novo
		l.fim = novo
	} else {
		l.fim.proximo = novo
		l.fim = novo
	}
	l.tamanho++
}

func (l *lista) imprimir() {
	atual := l.inicio
	fmt.Print("[ ")
	for atual != nil {
		fmt.Printf("%d ", atual.valor)
		atual = atual.proximo
	}
	fmt.Printf("] (tamanho=%d)\n", l.tamanho)
}

func main() {
	l := &lista{}

	l.adicionarFim(10)   // [10]
	l.adicionarInicio(5) // [5 10]
	l.adicionarFim(20)   // [5 10 20]
	l.adicionarInicio(1) // [1 5 10 20]
	l.adicionarFim(30)   // [1 5 10 20 30]

	fmt.Println("Resultado esperado: [ 1 5 10 20 30 ] (tamanho=5)")
	fmt.Print("Resultado obtido:   ")
	l.imprimir()
}
