package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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
