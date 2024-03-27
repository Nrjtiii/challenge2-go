package main

import "fmt"

type Orang struct { //Deklarasi struct
	massa  float32
	tinggi float32
}

func (Orang Orang) checkkBmi() float32 { //Fungsi untuk menghitung struct

	bmi := Orang.massa / (Orang.tinggi * Orang.tinggi)

	return bmi
}

func main() {

	mark := Orang{78, 1.69}
	John := Orang{92, 1.95}

	markHigherBMI := mark.checkkBmi() > John.checkkBmi() //membandingkan nilai BMI

	fmt.Println(markHigherBMI)
}
