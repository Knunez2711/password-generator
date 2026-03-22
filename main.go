package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Knunez2711/password-generator/generator"
)

func main() {
	length := flag.Int("length", 16, "Length of each password")
	count := flag.Int("count", 1, "Number of passwords to generate")
	noUpper := flag.Bool("no-upper", false, "Exclude uppercase letters")
	noDigits := flag.Bool("no-digits", false, "Exclude digits")
	noSymbols := flag.Bool("no-symbols", false, "Exclude symbols")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: passgen [options]\n\nOptions:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  passgen\n")
		fmt.Fprintf(os.Stderr, "  passgen -length 20 -count 5\n")
		fmt.Fprintf(os.Stderr, "  passgen -length 12 -no-symbols\n")
	}

	flag.Parse()

	opts := generator.Options{
		Length:     *length,
		Count:      *count,
		UseUpper:   !*noUpper,
		UseDigits:  !*noDigits,
		UseSymbols: !*noSymbols,
	}

	g := generator.SecureGenerator{}
	passwords, err := g.Generate(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	for _, pw := range passwords {
		fmt.Println(pw)
	}
}