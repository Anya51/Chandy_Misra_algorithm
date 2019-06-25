// Implementação do algsoritmo de travessia de Tarry
package main

import (
	"fmt"
	"math"
)

//Token ...
type Token struct {
	Sender string
}

type MSG struct {
	dist float64
	id   string
}

type Token struct {
	Sender string
}

//Neighbour ...
type Neighbour struct {
	id    string
	From  chan MSG
	To    chan MSG
	Value int
}

/*
func (id string, token Token, neighs ...Neighbour) {
	var pai Neighbour

	// Redeirecionando todos os canais de entrada para um único canal "in" de entrada
	in := make(chan Token, 1)
	nmap := make(map[string]Neighbour)
	for _, neigh := range neighs {
		nmap[neigh.id] = neigh
		go redirect(in, neigh)
	}

	if token.Sender == "init" {
		// o iniciador
		fmt.Printf("* %s é raiz.\n", id)
		token.Sender = id
		neighs[0].To <- token
		size := len(neighs)
		for i := 1; i < size; i++ {
			tk := <-in
			fmt.Printf("From %s to %s\n", tk.Sender, id)
			tk.Sender = id
			neighs[i].To <- tk
		}
		tk := <-in
		fmt.Printf("From %s to %s\n", tk.Sender, id)
		fmt.Println("Fim!")
	} else {
		// o não iniciador
		tk := <-in
		fmt.Printf("From %s to %s\n", tk.Sender, id)
		for _, neigh := range neighs {
			if pai.id == "" {
				pai = nmap[tk.Sender]
				fmt.Printf("* %s é pai de %s\n", pai.id, id)
			}
			// Entrega o token para o vizinho se ele não for o pai
			if pai.id != neigh.id {
				tk.Sender = id
				neigh.To <- tk
				tk = <-in
				fmt.Printf("From %s to %s\n", tk.Sender, id)
			}
		}
		// Token volta para o pai depois de ter passado enviado para todos os vizinhos
		tk.Sender = id
		pai.To <- tk
	}

}
*/

func redirect(in chan Token, neigh Neighbour) {
	token := <-neigh.From
	in <- token
}

func process1(id string, token Token, neighs ...Neighbour) {
	var parent Neighbour

	in := make(chan Token, 1)
	nmap := make(map[string]Neighbour)
	for _, neigh := range neighs {
		nmap[neigh.id] = neigh
		go redirect(in, neigh)
	}

	if token.Sender == "init" {
		// o iniciador
		dist := 0.0
		//parent = nil
		token.Sender = id
		msg := MSG{dist, id}

		for _, neigh := range neighs {
			neigh.To <- msg
		}

	} else {

		dist := math.Inf(1)
		//token.Parentp = nil
	}
}

func main() {

	pQ := make(chan Token, 1)
	qP := make(chan Token, 1)

	ValueQP := 2

	pR := make(chan Token, 1)
	rP := make(chan Token, 1)

	ValuePR := 2

	qR := make(chan Token, 1)
	rQ := make(chan Token, 1)

	ValueQR := 2

	/*
		rT := make(chan Token, 1)
		rS := make(chan Token, 1)
		sR := make(chan Token, 1)
		sT := make(chan Token, 1)
		tR := make(chan Token, 1)
		tS := make(chan Token, 1)
	*/
	//	go process1("T", Token{}, Neighbour{"R", rT, tR}, Neighbour{"S", sT, tS})
	//	go process1("S", Token{}, Neighbour{"R", rS, sR}, Neighbour{"T", tS, sT})
	go process1("R", Token{}, Neighbour{"Q", qR, rQ, ValueQR}, Neighbour{"P", pR, rP, ValuePR}) //, Neighbour{"T", tR, rT}, Neighbour{"S", sR, rS})
	go process1("Q", Token{}, Neighbour{"P", pQ, qP, ValueQP}, Neighbour{"R", rQ, qR, ValueQR})
	process1("P", Token{Sender: "init"}, Neighbour{"Q", qP, pQ, ValueQP}, Neighbour{"R", rP, pR, ValuePR})

}
