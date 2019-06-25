// Implementação do algsoritmo de travessia de Tarry
package main

import (
	"fmt"
	"math"
)

//Token ...
type Token struct {
	Sender  string
	Distp   float64
	Parentp *int
	value   int
}

//Neighbour ...
type Neighbour struct {
	ID   string
	From chan Token
	To   chan Token
}

func redirect(in chan Token, neigh Neighbour) {
	token := <-neigh.From
	fmt.Println(token)
	in <- token
}

/*
func (id string, token Token, neighs ...Neighbour) {
	var pai Neighbour

	// Redeirecionando todos os canais de entrada para um único canal "in" de entrada
	in := make(chan Token, 1)
	nmap := make(map[string]Neighbour)
	for _, neigh := range neighs {
		nmap[neigh.Id] = neigh
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
			if pai.Id == "" {
				pai = nmap[tk.Sender]
				fmt.Printf("* %s é pai de %s\n", pai.Id, id)
			}
			// Entrega o token para o vizinho se ele não for o pai
			if pai.Id != neigh.Id {
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
func process1(ID string, token Token, neighs ...Neighbour) {

	in := make(chan Token, 1)
	nmap := make(map[string]Neighbour)
	for _, neigh := range neighs {
		nmap[neigh.ID] = neigh
		go redirect(in, neigh)
	}

	if token.Sender == "init" {
		// o iniciador
		fmt.Printf("* %s é raiz.\n", ID)
		token.Distp = 0
		token.Parentp = nil
		token.Sender = ID

		neighs[0].To <- token

		tk := <-in
		fmt.Printf("From %s to %s\n", tk.Sender, ID)
	} else {

		tk := <-in

		fmt.Printf("hi From %s to %s\n", tk.Sender, ID)
		neighs[1].To <- tk
		token.Distp = math.Inf(1)
		token.Parentp = nil
	}
}

func main() {

	pQ := make(chan Token, 1)
	pR := make(chan Token, 1)
	qP := make(chan Token, 1)
	qR := make(chan Token, 1)
	rQ := make(chan Token, 1)
	rP := make(chan Token, 1)
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
	go process1("R", Token{}, Neighbour{"Q", qR, rQ}, Neighbour{"P", pR, rP}) //, Neighbour{"T", tR, rT}, Neighbour{"S", sR, rS})
	go process1("Q", Token{}, Neighbour{"P", pQ, qP}, Neighbour{"R", rQ, qR})
	process1("P", Token{Sender: "init"}, Neighbour{"Q", qP, pQ}, Neighbour{"R", rP, pR})

}
