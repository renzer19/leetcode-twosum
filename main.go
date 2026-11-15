package main

import "fmt"

func main() {
	logTransaksi := []int{20000, 30000, 45000, 55000, 70000}
	targetTransfer := 85000

	hasil := cariPasangan(logTransaksi, targetTransfer)
	fmt.Println("Index transaksi yang cocok:", hasil)
	cariPasangan(logTransaksi, targetTransfer)

}

func cariPasangan(data []int, target int) []int {
	list := make(map[int]int)
	for index, am := range data {
		current := am
		dicari := target - current
		indexPasangan, ketemu := list[dicari]
		if ketemu {
			return []int{indexPasangan, index}
		}
		list[am] = index
	}
	return nil
}
