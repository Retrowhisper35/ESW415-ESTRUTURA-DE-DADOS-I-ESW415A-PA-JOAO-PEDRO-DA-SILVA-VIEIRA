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

// adicionarPosicao insere valor no índice posicao (0 = início, tamanho = fim).
// Retorna false se a posição for inválida.
func (l *lista) adicionarPosicao(valor int, posicao int) bool {
	if posicao < 0 || posicao > l.tamanho {
		return false
	}

	if posicao == 0 {
		l.adicionarInicio(valor)
		return true
	}

	if posicao == l.tamanho {
		l.adicionarFim(valor)
		return true
	}

	anterior := l.inicio
	for i := 0; i < posicao-1; i++ {
		anterior = anterior.proximo
	}

	novo := &no{}
	novo.valor = valor
	novo.proximo = anterior.proximo // 1º: o novo aponta para o restante da lista
	anterior.proximo = novo         // 2º: o anterior passa a apontar para o novo
	l.tamanho++

	return true
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

	ok := l.adicionarPosicao(99, 2)
	fmt.Printf("\nadicionarPosicao(99, 2) -> ok=%v (esperado true)\n", ok)
	fmt.Print("Resultado obtido:   ")
	l.imprimir()
	fmt.Println("Resultado esperado: [ 1 5 99 10 20 30 ] (tamanho=6)")

	okInicio := l.adicionarPosicao(0, 0)
	fmt.Printf("\nadicionarPosicao(0, 0) -> ok=%v (esperado true, insere no início)\n", okInicio)
	l.imprimir()

	okFim := l.adicionarPosicao(999, l.tamanho)
	fmt.Printf("\nadicionarPosicao(999, tamanho) -> ok=%v (esperado true, insere no fim)\n", okFim)
	l.imprimir()

	okInvalida := l.adicionarPosicao(0, -1)
	okForaDoIntervalo := l.adicionarPosicao(0, l.tamanho+1)
	fmt.Printf("\nadicionarPosicao(0, -1) -> ok=%v (esperado false)\n", okInvalida)
	fmt.Printf("adicionarPosicao(0, tamanho+1) -> ok=%v (esperado false)\n", okForaDoIntervalo)
}
