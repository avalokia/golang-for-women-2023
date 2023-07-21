package helpers

import (
	"errors"
	"strconv"
	"strings"
)

// Function to find Peserta based on an arg
func FindPeserta(dataPeserta []Peserta, arg string) (Peserta, error) {
	// Check whether the arg is for name or ID
	argInt, err := strconv.Atoi(arg)
	var peserta Peserta
	if err != nil {
		// Converting string to number is error, means the arg is for the name
		peserta = findPesertaByName(dataPeserta, arg)
	} else {
		// coverting string to number is not error, means the arg is for the ID
		peserta = findPesertaByID(dataPeserta, argInt)
	}
	// Peserta is empty
	if (peserta == Peserta{}) {
		return Peserta{}, errors.New("Peserta tidak ditemukan")
	} else {
		return peserta, nil
	}
}

// Function to find peserta by ID
func findPesertaByID(dataPeserta []Peserta, id int) (peserta Peserta) {
	for _, value := range dataPeserta {
		if value.ID == id {
			return value
		}
	}
	return Peserta{}
}

// Function to find peserta by name. Will return the first peserta found in dummy data
func findPesertaByName(dataPeserta []Peserta, name string) (peserta Peserta) {
	for _, value := range dataPeserta {
		// If the name exactly the same or have part of name
		if value.Nama == name || strings.Contains(value.Nama, name) {
			return value
		}
	}
	return Peserta{}
}
