package main

import "fmt"

func main() {
	//fmt.Println("Hello World")
	//var
	var age int = 90
	var name string = "100.20"
	apakahMuda := true
	phi := 3.14

	var sisi int = 10
	luas := sisi * sisi
	fmt.Println(age, name, apakahMuda, phi)
	fmt.Println(luas)

	var umur int = 17
	if umur == 17 {
		fmt.Println("Sweet Seventeen")
	} else if umur < 17 {
		fmt.Println("Umur muda")
	} else {
		fmt.Println("Umur Invalid")
	}

	if umur > 17 && umur < 20 {
		fmt.Println("Sweet Seventeen")
	}

	switch umur {
	case 10:
		fmt.Println("10")
	case 20:
		fmt.Println("20")
	case 30:
		fmt.Println("30")
	default:
		fmt.Println("Invalid")
	}

	for i:= 0; i < 100; i=i+1{
		fmt.Println("Hello World : ", i)
	}
	var luass int
	var messages string
	luass, messages = tambah(10, 30)
	fmt.Println(luass, messages)
	// tambah(20, 50)
}
func tambah(a int, b int) (int, string) {
	// var a = 10
	// var b = 20
	var jumlah = a + b
	// fmt.Println(jumlah)
	return jumlah, "Success"
}
