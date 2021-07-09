package main
import ("fmt" 
"strings")

func main() {
	var kata string
	fmt.Print("masukkan kata : ")
	fmt.Scanf("%s", &kata)
	fmt.Println(palindrome(strings.ToLower(kata)))
}

func palindrome(value string) string {	
	for i:=0; i < len(value)/2; i++ {
		if (value[i] != value[len(value) -1 - i] ) {
			return "bukan polindrome"
		}
	}
	return "polindrome"
}

