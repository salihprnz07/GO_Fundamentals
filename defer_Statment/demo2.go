package defer_statment

import "fmt"

func Demo(sayi int) string {
	defer DeferFunc()

	if sayi%2 == 0 {
		fmt.Println("FÇift")
		return "Çift"
	}

	if sayi%2 != 0 {
		fmt.Println("tek")
	}

	return "Belli değil"
}

func Test() {
	sonuc := Demo(8)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti")
}
