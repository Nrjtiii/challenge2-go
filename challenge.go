package main

import "fmt"

func average(nilai []float32) float32 {	         		//fungsi untuk mencari nilai rata-rata masing-masing kelompok

	var avg float32
	sum := nilai[0]

	for i := 1; i < len(nilai); i++ {

		sum = sum + nilai[i]

	}

	avg = sum / float32(len(nilai))
	return avg

}

func standar(avgKoala, avgLumba float32) {  			//fungsi dengan aturan standar, fungsi ini akan menerima nilai
														//rata-rata setiap tim kemudian membandingkan nilai tersebut
	if avgKoala > avgLumba {
		fmt.Println("Pemenangnya tim Koala")
	} else if avgKoala < avgLumba {
		fmt.Println("Pemenangnya tim Lumba-Lumba")
	} else {
		fmt.Println("Hasilnya seri")
	}
}

func bonus1(avgKoala, avgLumba float32) {				//fungsi dengan aturan bonus1, fungsi ini akan menerima nilai
														//rata-rata setiap tim kemudian membandingkan nilai tersebut
	if avgKoala > avgLumba && avgKoala >= 100 {			//pada fungsi ini ada aturan tambahan yaitu nilai tim yang menang 
		fmt.Println("Pemenangnya tim Koala")			//harus lebih dari 100
	} else if avgKoala < avgLumba && avgLumba >= 100 {
		fmt.Println("Pemenangnya tim Lumba-Lumba")
	} else {
		fmt.Println("Tidak memenuhi syarat")
	}
}

func bonus2(avgKoala, avgLumba float32) {
													//fungsi dengan aturan bonus2, fungsi ini akan menerima nilai
													//rata-rata setiap tim kemudian membandingkan nilai tersebut
	if avgKoala > avgLumba && avgKoala >= 100 {		//pada fungsi ini ada aturan tambahan yaitu nilai kedua tim 
		fmt.Println("Pemenangnya tim Koala")		//harus lebih dari 100 agar bisa mendapat trofi
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

	lumba := []float32{97, 112, 101} //slice yang menyimpan nilai kelompok lumba-lumba
	koala := []float32{109, 95, 106} //slice yang menyimpan nilai kelompok koala

	avgLumba := average(lumba)   //kode ini memanggil fungsi average untuk menghitung nilai rata-rata
	avgKoala := average(koala)	 //setiap tim dan menyimpannya pada sebuah variabel

	fmt.Println("Nilai tim Koala: ", avgKoala)			//baris kode untuk menampilkan nilai kedua tim
	fmt.Println("Nilai tim Lumba-Lumba: ", avgLumba)

	bonus2(avgKoala, avgLumba) 	//kode untuk memanggil fungsi yang diinginkan, 								
}								//terdapat tiga fungsi yang tersedia yaitu standar, bonus1, bonus2, gunakan sesuai keinginan

