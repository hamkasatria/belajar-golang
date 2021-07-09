package main
import ("fmt" 
)

func main() {
	var awal, akhir int
	fmt.Print("masukkan tahun awal : ")
	fmt.Scanf("%d",&awal)

	fmt.Print("masukkan tahun akhir : " )
	fmt.Scanf("%d",&akhir)

	tampilkan(awal,akhir, 4)
	
}

func tampilkan (awal int, akhir int, selisih int){
	for i := awal; i <= akhir; i= i+selisih {
		fmt.Print(i," ")
	}
}

