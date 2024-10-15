package main

import "fmt"

// 2. Napiši strukturu koja predstavlja adresu, koja sadrži polja za grad i ulicu.
// Zatim napiši strukturu koja predstavlja osobu, koja sadrži ime, godine i adresu.
// Kreiraj metodu koja ispisuje puni opis osobe, uključujući njezinu adresu. (Sva polja) u formatu: Ime Prezime, 20 godina, živi u Grad, Ulica.

type Adresa struct {
	grad  string
	ulica string
}

type Osoba struct {
	ime    string
	godine int
	adresa Adresa
}

func (p Osoba) ispišiPodatke() {
	fmt.Printf("%s, %d godina, živi u %s %s", p.ime, p.godine, p.adresa.grad, p.adresa.ulica)
}

func main() {
	Osoba := Osoba{
		ime:    "Marko Marković",
		godine: 25,
		adresa: Adresa{
			grad:  "Mostar",
			ulica: "Stjepana Radića 3",
		},
	}

	Osoba.ispišiPodatke()
}
