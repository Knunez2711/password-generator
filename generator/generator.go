package generator

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	symbols   = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

// Options defines the configuration for password generation.
type Options struct {
	Length     int
	Count      int
	UseUpper   bool
	UseDigits  bool
	UseSymbols bool
}

// Validate checks that the options are sensible.
func (o Options) Validate() error {
	if o.Length < 4 {
		return errors.New("length must be at least 4")
	}
	if o.Count < 1 {
		return errors.New("count must be at least 1")
	}
	return nil
}

// Generator is the interface for anything that can generate passwords.
type Generator interface {
	Generate(opts Options) ([]string, error)
}

// SecureGenerator uses crypto/rand for cryptographically secure passwords.
type SecureGenerator struct{}

// Generate produces a slice of passwords based on the given options.
func (g SecureGenerator) Generate(opts Options) ([]string, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	charset := buildCharset(opts)
	if len(charset) == 0 {
		return nil, errors.New("charset is empty: enable at least one character type")
	}

	passwords := make([]string, 0, opts.Count)
	for range opts.Count {
		pw, err := generateOne(charset, opts.Length)
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, pw)
	}

	return passwords, nil
}

// buildCharset assembles the character pool based on options.
func buildCharset(opts Options) string {
	var sb strings.Builder
	sb.WriteString(lowercase)
	if opts.UseUpper {
		sb.WriteString(uppercase)
	}
	if opts.UseDigits {
		sb.WriteString(digits)
	}
	if opts.UseSymbols {
		sb.WriteString(symbols)
	}
	return sb.String()
}

// generateOne picks random characters from charset using crypto/rand.
func generateOne(charset string, length int) (string, error) {
	result := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := range result {
		idx, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		result[i] = charset[idx.Int64()]
	}

	return string(result), nil
}