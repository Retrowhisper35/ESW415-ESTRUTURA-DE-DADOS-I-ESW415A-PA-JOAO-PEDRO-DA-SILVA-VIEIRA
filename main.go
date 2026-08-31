package main

import "fmt"

// no representa um elemento da lista encadeada
type no struct {
	valor   int
	proximo *no
}

type lista struct {
	inicio       *no
	fim          *no
	qtdElementos int
}

func (l *lista) adicionarInicio(valor int) {
	novo := &no{valor: valor, proximo: l.inicio}
	l.inicio = novo

	if l.fim == nil {
		l.fim = novo
	}
	l.qtdElementos++
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
	l.qtdElementos++
}

// adicionarPosicao insere valor no índice posicao (0 = início, tamanho = fim).
// Retorna false se a posição for inválida.
func (l *lista) adicionarPosicao(valor int, posicao int) bool {
	if posicao < 0 || posicao > l.qtdElementos {
		return false
	}

	if posicao == 0 {
		l.adicionarInicio(valor)
		return true
	}

	if posicao == l.qtdElementos {
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
	l.qtdElementos++

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

	l.qtdElementos--
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
		l.qtdElementos--
		return valor, true
	}

	anterior := l.inicio
	for anterior.proximo != l.fim {
		anterior = anterior.proximo
	}

	anterior.proximo = nil
	l.fim = anterior
	l.qtdElementos--

	return valor, true
}

// removerPosicao remove e retorna o valor do nó no índice posicao.
// Retorna (0, false) se a posição for inválida ou a lista estiver vazia.
func (l *lista) removerPosicao(posicao int) (int, bool) {
	if posicao < 0 || posicao >= l.qtdElementos {
		return 0, false
	}

	if posicao == 0 {
		return l.removerInicio()
	}

	if posicao == l.qtdElementos-1 {
		return l.removerFim()
	}

	anterior := l.inicio
	for i := 0; i < posicao-1; i++ {
		anterior = anterior.proximo
	}

	removido := anterior.proximo
	anterior.proximo = anterior.proximo.proximo // pula o nó removido
	l.qtdElementos--

	return removido.valor, true
}

// posicao percorre a lista comparando atual.valor a cada passo e retorna o
// índice onde valorProcurado foi encontrado. Retorna (0, false) se não existir.
func (l *lista) posicao(valorProcurado int) (int, bool) {
	atual := l.inicio
	indice := 0

	for atual != nil {
		if atual.valor == valorProcurado {
			return indice, true
		}
		atual = atual.proximo
		indice++
	}

	return 0, false
}

// valorNaPosicao percorre a lista até posicaoProcurada e retorna o valor
// armazenado ali. Retorna (0, false) se a posição não existir.
func (l *lista) valorNaPosicao(posicaoProcurada int) (int, bool) {
	if posicaoProcurada < 0 || posicaoProcurada >= l.qtdElementos {
		return 0, false
	}

	atual := l.inicio
	for i := 0; i < posicaoProcurada; i++ {
		atual = atual.proximo
	}

	return atual.valor, true
}

// tamanho percorre a lista e conta os nós.
func (l *lista) tamanho() int {
	contador := 0
	atual := l.inicio

	for atual != nil {
		contador++
		atual = atual.proximo
	}

	return contador
}

// imprimir exibe os valores da lista no formato "10 -> 20 -> 30 -> nil".
func (l *lista) imprimir() {
	atual := l.inicio
	for atual != nil {
		fmt.Printf("%d -> ", atual.valor)
		atual = atual.proximo
	}
	fmt.Println("nil")
}

func exibirMenu() {
	fmt.Println("\n===== Lista Encadeada =====")
	fmt.Println("1  - Adicionar no início")
	fmt.Println("2  - Adicionar no fim")
	fmt.Println("3  - Adicionar em posição específica")
	fmt.Println("4  - Remover do início")
	fmt.Println("5  - Remover do fim")
	fmt.Println("6  - Remover de posição específica")
	fmt.Println("7  - Buscar posição de um valor")
	fmt.Println("8  - Buscar valor em uma posição")
	fmt.Println("9  - Tamanho da lista")
	fmt.Println("10 - Imprimir lista")
	fmt.Println("0  - Sair")
	fmt.Print("Escolha uma opção: ")
}

func main() {
	l := &lista{}
	opcao := -1

	for opcao != 0 {
		exibirMenu()
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			var valor int
			fmt.Print("Valor: ")
			fmt.Scan(&valor)
			l.adicionarInicio(valor)
			fmt.Println("Adicionado no início.")

		case 2:
			var valor int
			fmt.Print("Valor: ")
			fmt.Scan(&valor)
			l.adicionarFim(valor)
			fmt.Println("Adicionado no fim.")

		case 3:
			var valor, posicao int
			fmt.Print("Valor: ")
			fmt.Scan(&valor)
			fmt.Print("Posição: ")
			fmt.Scan(&posicao)
			if l.adicionarPosicao(valor, posicao) {
				fmt.Println("Adicionado na posição", posicao)
			} else {
				fmt.Println("Posição inválida.")
			}

		case 4:
			valor, ok := l.removerInicio()
			if ok {
				fmt.Println("Removido:", valor)
			} else {
				fmt.Println("Lista vazia.")
			}

		case 5:
			valor, ok := l.removerFim()
			if ok {
				fmt.Println("Removido:", valor)
			} else {
				fmt.Println("Lista vazia.")
			}

		case 6:
			var posicao int
			fmt.Print("Posição: ")
			fmt.Scan(&posicao)
			valor, ok := l.removerPosicao(posicao)
			if ok {
				fmt.Println("Removido:", valor)
			} else {
				fmt.Println("Posição inválida.")
			}

		case 7:
			var valor int
			fmt.Print("Valor: ")
			fmt.Scan(&valor)
			idx, ok := l.posicao(valor)
			if ok {
				fmt.Println("Encontrado na posição", idx)
			} else {
				fmt.Println("Valor não encontrado.")
			}

		case 8:
			var posicao int
			fmt.Print("Posição: ")
			fmt.Scan(&posicao)
			valor, ok := l.valorNaPosicao(posicao)
			if ok {
				fmt.Println("Valor:", valor)
			} else {
				fmt.Println("Posição inválida.")
			}

		case 9:
			fmt.Println("Tamanho:", l.tamanho())

		case 10:
			l.imprimir()

		case 0:
			fmt.Println("Encerrando.")

		default:
			fmt.Println("Opção inválida.")
		}
	}
}
