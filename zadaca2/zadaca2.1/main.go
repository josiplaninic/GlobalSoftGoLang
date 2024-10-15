package main

import "fmt"

// 1. Napiši strukturu koja predstavlja pravokutnik i sadrži polja za širinu i visinu. 
// Kreiraj metode za izračunavanje površine i opsega pravokutnika,
// metode moraju biti vezane za novo kreiranu strukturu Pravokutnika.
//  U main funkciji inicijalizirajte jedan pravokutnik, te pozovite iznad
//   kreirane metode za računanje površine i opsega.

type Pravokutnik struct {
	sirina float64
	visina float64
}

func (p Pravokutnik) izracunajPovrsinu() float64{
	return p.sirina*p.visina
}
func (p Pravokutnik) izracunajOpseg() float64{
	return 2*p.sirina+2*p.visina
}

func main() {
	pravokutnik := Pravokutnik{
		sirina:2.5,
		visina:3.5,
	}

	fmt.Printf("Visina pravokutnika -> %f \n Sirina pravokutnika -> %f \n", pravokutnik.visina, pravokutnik.sirina)
	fmt.Printf("Opseg pravokutnika -> %f \n" ,pravokutnik.izracunajOpseg())
	fmt.Printf("Povrsina pravokutnika -> %f", pravokutnik.izracunajPovrsinu())


}