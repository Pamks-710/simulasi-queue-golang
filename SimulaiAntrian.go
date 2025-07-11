package main

import (
	"fmt"
)

var queue []string // Slice untuk menyimpan antrian

// Fungsi untuk menambahkan ke antrian
func enqueue(name string) {
	queue = append(queue, name)
	fmt.Println(name, "berhasil ditambahkan ke antrian.")
}

// Fungsi untuk menghapus dari antrian
func dequeue() {
	if len(queue) == 0 {
		fmt.Println("Antrian kosong.")
		return
	}
	fmt.Println(queue[0], "telah keluar dari antrian.")
	queue = queue[1:] // Menghapus elemen pertama
}

// Fungsi untuk menampilkan antrian saat ini
func displayQueue() {
	if len(queue) == 0 {
		fmt.Println("Antrian kosong.")
	} else {
		fmt.Println("Antrian saat ini:", queue)
	}
}

func main() {
	var choice int
	var name string

	for {
		fmt.Println("\n=== Program Simulasi Antrian ===")
		fmt.Println("1. Tambah ke antrian")
		fmt.Println("2. Keluarkan dari antrian")
		fmt.Println("3. Lihat antrian")
		fmt.Println("4. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Masukkan nama: ")
			fmt.Scan(&name)
			enqueue(name)
		case 2:
			dequeue()
		case 3:
			displayQueue()
		case 4:
			fmt.Println("Terima kasih! Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
