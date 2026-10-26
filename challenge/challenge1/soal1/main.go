package main

import "fmt"

//spec 1
type ConversionResult struct {
	Currency string
	Amount   float64
}

//spec 2
var toIDR = map[string]float64{
	"USD": 15000,
	"EUR": 16000,
	"JPY": 140,
	"SGD": 11000,
}

func main() {
	
	var IDR float64

	fmt.Print("Masukkan jumlah uang dalam Rupiah: ")
	fmt.Scan(&IDR)
	currencies := []string{"USD", "EUR", "JPY", "SGD"}

	//spec3
	converted :=[]ConversionResult{} 

	//spec 4
	for _, currency := range currencies{
		rate, ok := toIDR[currency]
		
		if ok{
			result:= ConversionResult{
				Currency: currency,
				Amount: IDR/rate,
			}
			//spec 5
			converted = append(converted, result)
		}else{
			fmt.Println("Currency not found")
		}
		
	}

	fmt.Printf("\nKonversi dari %.0f IDR:", IDR)
	fmt.Printf("\nUSD: %.3f", converted[0].Amount)
	fmt.Printf("\nEUR: %.3f", converted[1].Amount)
	fmt.Printf("\nJPY: %.3f", converted[2].Amount)
	fmt.Printf("\nSGD: %.3f", converted[3].Amount)
	
}