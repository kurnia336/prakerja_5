package main

import "fmt"

func main() {
	var bilangan int = 11
	var hasil int
	hasil=prima(bilangan)
	if hasil==1{
		fmt.Println("bilangan prima")
	}else{
		fmt.Println("bukan bilangan prima")
	}

	kelipatan7(49)
	kelipatan7(2)

	fmt.Println(luasTrapesium(3, 5, 10))

}
func prima(angka int) int {
	if angka == 1{
		return 0;
	}
	for i:=2; i <= angka/2; i++{
		if angka % i == 0{
			return 0;
		}
	}
	return 1;
}

func kelipatan7(angka int) int{
	if angka % 7 == 0{
		fmt.Println("kelipatan 7")
	}else{
		fmt.Println("bukan kelipatan 7")
	}
	return 1;
}

func luasTrapesium(a int, b int, t int) int{
	hasil := ((a + b) / 2) * t
	return hasil
}