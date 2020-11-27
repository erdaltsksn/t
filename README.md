# t (Go Translate)

[![PkgGoDev](https://pkg.go.dev/badge/github.com/erdaltsksn/t)](https://pkg.go.dev/github.com/erdaltsksn/t)
![Go (build)](https://github.com/erdaltsksn/t/workflows/Go%20(build)/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/erdaltsksn/t)](https://goreportcard.com/report/github.com/erdaltsksn/t)
![CodeQL](https://github.com/erdaltsksn/t/workflows/CodeQL/badge.svg)

`t` helps you translate Go apps into multiple languages.

## Features

- Configurable
- Strings with variables
- Translation folders that is relative to module
- Translation files in `yaml` format

## Requirements

- [Golang](https://golang.org)

## Getting Started

```sh
go get github.com/erdaltsksn/t
touch main.go
```

**Directory Structure:**

```sh
1 .
2 ├── translations
3 │   ├── en.yaml
4 │   └── tr.yaml
5 └── main.go
```

**main.go:**

```go
package main

import (
	"fmt"

	"github.com/erdaltsksn/t"
	"golang.org/x/text/language"
)

func main() {
	t.Configure(t.Config{
		Language:         language.Turkish,
		FallbackLanguage: language.English,
		TranslationFolder: struct {
			Path     string
			Relative bool
		}{
			Path:     "/translations",
			Relative: true,
		},
	})

	fmt.Println(t.Translate("msgHello"))
	fmt.Println(t.Translate("msgMorning", "Adam"))
}
```

**en.yaml and tr.yaml:**

```yaml
# en.yaml
msgHello: "Hello World"
msgMorning: "Good morning, %v"
# tr.yaml
msgHello: "Merhaba Dünya"
msgMorning: "Günaydın, %v"
```

**Output:**

```txt
Merhaba Dünya
Günaydın, Adam
```

## Installation

```sh
go get github.com/erdaltsksn/t
```

## Updating / Upgrading

```sh
go get -u github.com/erdaltsksn/t
```

## Usage

```go
t.Translate("keyText")
t.Translate("keyText", "Variable")
// Print in different language
message.NewPrinter(language.Chinese).Sprintf("keyText", "Variable-1", "var-2")
```

## Contributing

If you want to contribute to this project and make it better, your help is very
welcome. See [CONTRIBUTING](.github/CONTRIBUTING.md) for more information.

## Security Policy

If you discover a security vulnerability within this project, please follow our
[Security Policy Guide](.github/SECURITY.md).

## Code of Conduct

This project adheres to the Contributor Covenant [Code of Conduct](.github/CODE_OF_CONDUCT.md).
By participating, you are expected to uphold this code.

## Disclaimer

In no event shall we be liable to you or any third parties for any special,
punitive, incidental, indirect or consequential damages of any kind, or any
damages whatsoever, including, without limitation, those resulting from loss of
use, data or profits, and on any theory of liability, arising out of or in
connection with the use of this software.
