package main

import ("fmt"
"strings"
)

func main(){
	// var value string
	// fmt.Print("Masukkan kalimat : ")
	// fmt.Scanln("%s",&value)

	var value = "I am A Great human"
	memecahString(value);
}

func memecahString(s string){
	var katas = strings.Split(s," ")
	for _, kata := range katas{
		if kata == strings.Title(kata) {
			fmt.Print(strings.Title(strings.ToLower(membalikChar(kata))), " ")
		}else{
			fmt.Print(membalikChar(kata), " ")
		}
	}
}

func membalikChar(s string) (result string){
	for _, v := range s {
        result = string(v) + result
    }
    return 
}