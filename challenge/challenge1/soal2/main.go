package main

import (
	"fmt"
	"slices"
	"strings"
)

type CardRules struct{
	jenisKartu string
	prefixes []string
	panjangNomor []int
}

func main(){
	cards:= []CardRules{
		{
			jenisKartu: "China UnionPay",
			prefixes: []string{"62"},
			panjangNomor: []int{16,17,18,19},
		},
		{
			jenisKartu: "Switch",
			prefixes: []string{"4903", "4905", "4911", "4936", "564182","633110", "6333", "6759"},
			panjangNomor: []int{16,18,19},
		},
	}

	input := []string{
		"6212345678901234",
		"6221265678901234",
		"6221261234567890",
		"4903887659126748923",
		"491188765912674892",
		"6331107659126746",
		"5641828876591267489",
	}
	
	//iterasi tiap baris input
	for _, cardNum := range input{
		cardType := "Jenis kartu tidak dikenali"
		 
		//cek aturan tiap kartu
		for  _, rule := range cards {

			//cek panjang kartu
			if !slices.Contains(rule.panjangNomor, len(cardNum)){
				continue
			}

			//cek prefix dari seluruh list prefixes kartu saat ini
			for _, prefix := range rule.prefixes{
				if strings.HasPrefix(cardNum, prefix){
					cardType = rule.jenisKartu
					break
				}
			}

			if cardType != "Jenis kartu tidak dikenali"{
				break
			}
		}
		fmt.Printf("Nomor %s adalah %s \n", cardNum, cardType)
	}
}
