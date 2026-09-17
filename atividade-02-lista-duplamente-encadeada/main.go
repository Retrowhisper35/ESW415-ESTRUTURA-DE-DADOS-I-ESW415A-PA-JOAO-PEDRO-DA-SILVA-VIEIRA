package main

import "fmt"

type No struct {
	valor int
	ant   *No
	prox  *No
}

type Lista struct {
	head *No
	tail *No
}

func (l *Lista) InserirInicio(valor int) {
	novo := &No{valor: valor}

	if l.head == nil {
		l.head = novo
		l.tail = novo
		return
	}

	novo.prox = l.head
	l.head.ant = novo
	l.head = novo
}

func (l *Lista) InserirFimSemTail(valor int) {
	novo := &No{valor: valor}

	if l.head == nil {
		l.head = novo
		l.tail = novo
		return
	}

	atual := l.head
	for atual.prox != nil {
		atual = atual.prox
	}

	atual.prox = novo
	novo.ant = atual
	l.tail = novo
}

func (l *Lista) RemoverHead() {
	if l.head == nil {
		return
	}

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return
	}

	l.head = l.head.prox
	l.head.ant = nil
}

func (l *Lista) RemoverTail() {
	if l.tail == nil {
		return
	}

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return
	}

	l.tail = l.tail.ant
	l.tail.prox = nil
}

func (l *Lista) Imprimir(rotulo string) {
	fmt.Printf("  %-8s ", rotulo)
	for atual := l.head; atual != nil; atual = atual.prox {
		fmt.Printf("%d ", atual.valor)
	}
	fmt.Println()
}

func novaListaInicial() *Lista {
	lista := &Lista{}
	for _, v := range []int{10, 20, 50, 60, 80} {
		lista.InserirFimSemTail(v)
	}
	return lista
}

func questao1() {
	fmt.Println("Questao 1 - Inserir 5 no comeco")
	lista := novaListaInicial()
	lista.Imprimir("antes:")

	lista.InserirInicio(5)

	lista.Imprimir("depois:")
	fmt.Println()
}

func questao2() {
	fmt.Println("Questao 2 - Inserir 5 no final SEM o tail")
	lista := novaListaInicial()
	lista.Imprimir("antes:")

	lista.InserirFimSemTail(5)

	lista.Imprimir("depois:")
	fmt.Println()
}

func questao3() {
	fmt.Println("Questao 3 - Remover o head")
	lista := novaListaInicial()
	lista.Imprimir("antes:")

	lista.RemoverHead()

	lista.Imprimir("depois:")
	fmt.Println()
}

func questao4() {
	fmt.Println("Questao 4 - Remover o tail")
	lista := novaListaInicial()
	lista.Imprimir("antes:")

	lista.RemoverTail()

	lista.Imprimir("depois:")
	fmt.Println()
}

func main() {
	questao1()
	questao2()
	questao3()
	questao4()
}
