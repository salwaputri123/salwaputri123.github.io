package main

import "fmt"

// fungsi untuk pertambahan dua bilangan
func tambah(a, b int) int {
    return a + b
}

func main() {
	fmt.Println("Hello, World!")
    // Contoh penggunaan fungsi tambah
    hasil := tambah(3, 5)
    fmt.Println("Hasil pertambahan:", hasil)
}