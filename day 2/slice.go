package main

import "fmt"

func main() {
	var days = [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}
	var slice1 = days[1:4]

	fmt.Println(slice1)
	// cek panjang data
	fmt.Println(len(slice1))
	// Cek kapasitas data
	fmt.Println(cap(slice1))

	// ubah array, index 0 bukan mengacu pada data array melainkan pada data si slice1
	// slice1[0] = "Lebaran"
	// fmt.Println(days)

	var slice2 = days[5:]
	fmt.Println(slice2)

	// Func append
	var slice3 = append(slice2, "Lebaran")
	fmt.Println("slice 2", slice3)

	slice3[1] = "Bukan Minggu"
	fmt.Println("slice 3", slice3)

	fmt.Println(slice2)
	fmt.Println(days)

	// function slice make, membuat slice dari awal
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Ipay"
	newSlice[1] = "Zulfadli"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println(copySlice)

	// notes: harus hati hati ketika membuat array. Jika tidak teliti maka yang dibuat adalah slice, bukan array
	iniArray := [...]int{
		1, 2, 3, 4, 5,
	}

	iniSlice := []int{
		1, 2, 3, 4, 5,
	}

	fmt.Println("ini array", iniArray)
	fmt.Println("ini slice", iniSlice)
}
