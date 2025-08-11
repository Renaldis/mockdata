package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"strings"


	"github.com/Renaldis/mockdata/data"
)

func main() {

	var help bool
	var inputPath, outputPath string

	flag.BoolVar(&help, "h", false, "Tampilkan cara menggunakan")
	flag.BoolVar(&help, "help", false, "Tampilkan cara menggunakan")

	flag.StringVar(&inputPath, "i", "", "Lokasi file JSON sebagai input")
	flag.StringVar(&inputPath, "input", "", "Lokasi file JSON sebagai input")
	flag.StringVar(&outputPath, "o", "", "Lokasi file JSON sebagai output")
	flag.StringVar(&outputPath, "output", "", "Lokasi file JSON sebagai output")

	flag.Parse()

	if help || inputPath == "" || outputPath == "" {
		printUsage()
		os.Exit(0)
	}
	if err := validateInput(inputPath); err != nil {
		fmt.Printf("Invalid input: %s\n", err)
		os.Exit(0)
	}

	if err := validateOutput(outputPath); err != nil {
		fmt.Printf("Invalid input: %s\n", err)
		os.Exit(0)
	}

	var mapping map[string]string

	if err := readInput(inputPath, &mapping); err != nil {
		fmt.Printf("Gagal membaca input: %s\n", err)
		os.Exit(0)
	}

	if err := validatetype(mapping); err != nil {
		fmt.Printf("Gagal validasi tipe data: %s\n", err)
		os.Exit(0)
	}

	result, err := generateOutput(mapping)
	if err != nil {
		fmt.Printf("Gagal membuat data: %s\n", err)
		os.Exit(0)
	}

	if err := writeOutput(outputPath, result); err != nil {
		fmt.Printf("Gagal menulis hasil: %s\n", err)
		os.Exit(0)
	}

}

func printUsage() {
	fmt.Println("Usage: mockdata [-i | --input] <input file> [-o | output] <output file>")
	fmt.Println("-i --input: File input berupa JSON sebagai template")
	fmt.Println("-o --output: File output berupa JSON sebagai hasil")
}

func validateInput(path string) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	return nil
}
func validateOutput(path string) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	fmt.Println("File sudah ada di lokasi")
	confirmOverwrite()
	return nil
}

func confirmOverwrite() {
	fmt.Println("Apakah anda ingin menimpa file yang sudah ada (y/t)")

	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.ToLower(strings.TrimSpace(response))

	if response != "y" && response != "yes" && response != "ya" {
		fmt.Println("membatalkan proses...")
		os.Exit(0)
	}
}

func readInput(path string, mapping *map[string]string) error {
	if path == "" {
		return errors.New("path tidak valid")
	}

	if mapping == nil {
		return errors.New("mapping tidak valid")
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	fileByte, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(fileByte) == 0 {
		return errors.New("input kosong")
	}

	if err := json.Unmarshal(fileByte, &mapping); err != nil {
		return err
	}
	return nil
}

func validatetype(mapping map[string]string) error {

	for _, value := range mapping {
		if !data.Supported[value] {
			return errors.New("tipe data tidak didukung")
		}
	}
	return nil
}

func generateOutput(mapping map[string]string) (map[string]any, error) {
	result := make(map[string]any)

	for key, dataType := range mapping {
		result[key] = data.Generate(dataType)
	}

	return result, nil
}

func writeOutput(path string, result map[string]any) error {
	if path == "" {
		return errors.New("path tidak valid")
	}

	// 0644 -> bawaan linux artinya kita bisa akses read, write, tapi orang lain hanya akses read aja
	// RDWR -> Read write
	// CREATE -> membuat
	// TRUNC -> menghapus/mengosongkan

	flags := os.O_RDWR | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(path, flags, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	resultByte, err := json.MarshalIndent(result, "", "  ")
	// marshaindent json versi rapih ada indentasi
	// marshal aja cuman sebaris nanti json nya
	if err != nil {
		return err
	}

	if _, err := file.Write(resultByte); err != nil {
		return err
	}

	return nil
}
