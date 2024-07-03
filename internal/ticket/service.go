package ticket

import (
	"bufio"
	"fmt"
	"github.com/xvbnm48/golang-terminal-learn/internal/db"
)

func InitDbService(filePath string) error {
	fmt.Println(filePath)
	return db.InitDB(filePath)
}

func CloseDb() {
	db.CloseDb()
}

func CreateTicket(scanner *bufio.Scanner) {
	fmt.Printf("Masukkan nama tiket: ")
	scanner.Scan()
	title := scanner.Text()
	fmt.Printf("masukkan isi tiket:")
	scanner.Scan()
	content := scanner.Text()

	query := `INSERT INTO tickets (title, content) VALUES (?, ?)`
	_, err := db.DB.Exec(query, title, content)
	if err != nil {
		fmt.Println("Error create ticket")
		return
	}

	fmt.Println("Ticket created")
}

func GetAllTicket(scanner *bufio.Scanner) {
	for {
		query := `SELECT * FROM tickets`
		rows, err := db.DB.Query(query)
		if err != nil {
			fmt.Println("Error get all ticket")
			return
		}
		defer rows.Close()

		fmt.Println("\n==============================")
		fmt.Println("    Daftar Tiket")
		fmt.Println("==============================")
		for rows.Next() {
			var ticket Ticket
			err := rows.Scan(&ticket.ID, &ticket.Title, &ticket.Content)
			if err != nil {
				fmt.Println("Error scanning ticket, ", err)
				return
			}
			fmt.Printf("ID: %d\nJudul: %s\nIsi: %s\n", ticket.ID, ticket.Title, ticket.Content)
		}

		err = rows.Err()
		if err != nil {
			fmt.Println("Error rows, ", err)
			return
		}

		fmt.Println("\n==============================")
		fmt.Println("ketik 'kembali' untuk kembali ke menu utama")
		scanner.Scan()
		if scanner.Text() == "kembali" {
			return
		}
	}
}
