package main

import (
	"biodata/helpers"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check whether the arg has input parameter or not
	if len(os.Args) < 2 {
		fmt.Printf("Tolong masukkan nama atau nomor absen\nContoh: 'go run biodata.go John' atau 'go run biodata.go 1'\n")
		return
	}
	// Get the argument value as string
	input := os.Args[1:]
	// Find peserta
	peserta, err := helpers.FindPeserta(helpers.DataDummy, strings.Join(input, " "))
	if err != nil {
		// Peserta not found
		fmt.Println(err)
		return
	}
	// Print peserta
	peserta.PrintPeserta()
}
