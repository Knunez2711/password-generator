# Password Generator

A CLI tool for generating cryptographically secure passwords, written in Go.

## Features

- Uses `crypto/rand` — cryptographically secure (not `math/rand`)
- Configurable length, count, and character sets
- Clean separation between CLI and core logic
- Idiomatic Go: structs, interfaces, error handling

## Usage

```bash
go build -o passgen .

# Default: 16 chars, all character types
./passgen

# 20 characters, 5 passwords
./passgen -length 20 -count 5

# No symbols
./passgen -length 12 -no-symbols

# Only lowercase + digits
./passgen -no-upper -no-symbols
```

## Options

| Flag         | Default | Description                  |
|--------------|---------|------------------------------|
| `-length`    | 16      | Length of each password      |
| `-count`     | 1       | Number of passwords to generate |
| `-no-upper`  | false   | Exclude uppercase letters    |
| `-no-digits` | false   | Exclude digits               |
| `-no-symbols`| false   | Exclude symbols              |

## Go Concepts Covered

- `structs` and methods
- `interfaces` (`Generator`)
- `crypto/rand` with `math/big` for unbiased random selection
- CLI flags with the `flag` package
- Error handling with descriptive `error` returns
- Package structure and module organization

## Project Structure

```
password-generator/
├── main.go              # CLI entrypoint, flag parsing
├── generator/
│   └── generator.go     # Core logic: Options, Generator interface, SecureGenerator
├── go.mod
└── README.md
```