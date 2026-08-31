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

// removerInicio remove e retorna o valor do primeiro nó.
// Retorna (0, false) se a lista estiver vazia.
func (l *lista) removerInicio() (int, bool) {
	if l.inicio == nil {
		return 0, false
	}

	removido := l.inicio
	l.inicio = removido.proximo

	if l.inicio == nil {
		// removeu o único elemento: fim também fica vazio
		l.fim = nil
	}

	l.tamanho--
	return removido.valor, true
}

// removerFim remove e retorna o valor do último nó.
// Retorna (0, false) se a lista estiver vazia.
func (l *lista) removerFim() (int, bool) {
	if l.fim == nil {
		return 0, false
	}

	valor := l.fim.valor

	if l.inicio == l.fim {
		// único elemento
		l.inicio = nil
		l.fim = nil
		l.tamanho--
		return valor, true
	}

	anterior := l.inicio
	for anterior.proximo != l.fim {
		anterior = anterior.proximo
	}

	anterior.proximo = nil
	l.fim = anterior
	l.tamanho--

	return valor, true
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

	fmt.Println("\n--- removerInicio / removerFim ---")
	l2 := &lista{}
	l2.adicionarFim(10)
	l2.adicionarFim(20)
	l2.adicionarFim(30) // [10 20 30]

	valor, ok2 := l2.removerInicio()
	fmt.Printf("removerInicio() -> valor=%d ok=%v (esperado 10 true)\n", valor, ok2)
	l2.imprimir()

	valor, ok2 = l2.removerFim()
	fmt.Printf("removerFim() -> valor=%d ok=%v (esperado 30 true)\n", valor, ok2)
	l2.imprimir()

	// lista com um único elemento: remoção deve zerar início e fim
	valor, ok2 = l2.removerInicio()
	fmt.Printf("removerInicio() (único elemento) -> valor=%d ok=%v (esperado 20 true)\n", valor, ok2)
	l2.imprimir()

	// lista vazia
	valor, ok2 = l2.removerInicio()
	fmt.Printf("removerInicio() (lista vazia) -> valor=%d ok=%v (esperado 0 false)\n", valor, ok2)

	valor, ok2 = l2.removerFim()
	fmt.Printf("removerFim() (lista vazia) -> valor=%d ok=%v (esperado 0 false)\n", valor, ok2)
}
