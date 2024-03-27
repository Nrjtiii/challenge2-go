package main

import "fmt"

func average(nilai []float32) float32 {

	var avg float32
	sum := nilai[0]

	for i := 1; i < len(nilai); i++ {

		sum = sum + nilai[i]

	}

	avg = sum / float32(len(nilai))
	return avg

}

func standar(avgKoala, avgLumba float32) {

	if avgKoala > avgLumba {
		fmt.Println("Pemenangnya tim Koala")
	} else if avgKoala < avgLumba {
		fmt.Println("Pemenangnya tim Lumba-Lumba")
	} else {
		fmt.Println("Hasilnya seri")
	}
}

func bonus1(avgKoala, avgLumba float32) {

	if avgKoala > avgLumba && avgKoala >= 100 {
		fmt.Println("Pemenangnya tim Koala")
	} else if avgKoala < avgLumba && avgLumba >= 100 {
		fmt.Println("Pemenangnya tim Lumba-Lumba")
	} else {
		fmt.Println("Tidak memenuhi syarat")
	}
}

func bonus2(avgKoala, avgLumba float32) {
	if avgKoala < 100 {
		if avgLumba >= 100 {
			fmt.Println("Pemenangnya tim Lumba-Lumba")
		} else {
			fmt.Println("Tidak ada tim yang mendapat trofi")
		}
	} else if avgKoala >= 100 {
		if avgLumba < 100 {
			fmt.Println("Pemenangnya adalah tim Koala")
		} else if avgKoala > avgLumba {
			fmt.Println("Pemenangnya adalah tim Koala")
		} else if avgKoala < avgLumba {
			fmt.Println("Pemenangnya adalah tim Lumba-Lumba")
		} else {
			fmt.Println("Hasilnya seri")
		}
	}
}

func main() {

	lumba := []float32{97, 112, 101}
	koala := []float32{109, 95, 106}

	avgLumba := average(lumba)
	avgKoala := average(koala)

	fmt.Println("Nilai tim Koala: ", avgKoala)
	fmt.Println("Nilai tim Lumba-Lumba: ", avgLumba)

	bonus2(avgKoala, avgLumba)

}
