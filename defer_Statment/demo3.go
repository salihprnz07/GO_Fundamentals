package defer_statment

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("Ürün kaydedildi:", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("Loglandı")
}

func Demo3() {
	p := product{productName: "Laptop", unitPrice: 3500}
	defer p.save()
	p = product{productName: "Mouse", unitPrice: 35}
	p.save()
	fmt.Println("Işlem başarılı")
}
