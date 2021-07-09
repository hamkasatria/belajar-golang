package main

import "fmt"

func main(){
	var arr = [...]int{15,1,3}
	var total = 0;
	for _,nilai := range arr{
		total = total+nilai
	}
	genereateFib(1,1,total)
}

func genereateFib(fibAwal int, fibAkhir int, total int){
	if fibAwal+fibAkhir<total {
		genereateFib(fibAkhir, fibAwal+fibAkhir, total)
	}else{
		fmt.Println("perlu di tambah : ",(fibAwal+fibAkhir)-total)
		fmt.Print("untuk mendapatkan nilai fib : ", fibAwal+fibAkhir)
	}
}