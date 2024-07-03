package main

import (
	"bufio"
	"fmt"
	"github.com/xvbnm48/golang-terminal-learn/internal/ticket"
	"os"
)

func main() {
	err := ticket.InitDbService()
	if err != nil {
		fmt.Println("Error init db")
		return
	}

	defer ticket.CloseDb()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nTiket CLI")
		fmt.Println("1. Buat Tiket")
		fmt.Println("2. Lihat Semua Tiket")
		fmt.Println("3. Hapus Tiket")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih opsi: ")
		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			ticket.CreateTicket(scanner)
		case "2":
			ticket.GetAllTicket(scanner)
		case "3":
			ticket.DeleteTicketById(scanner)
		case "4":
			fmt.Println("Keluar")
			os.Exit(0)
		default:
			fmt.Println("Opsi tidak valid")

		}
	}
}
