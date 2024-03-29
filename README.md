# goric

[![Documentation](https://godoc.org/github.com/digenaldo/goric?status.svg)](http://godoc.org/github.com/digenaldo/goric)
[![Coverage Status](https://coveralls.io/repos/github/digenaldo/goric/badge.svg?1&branch=master)](https://coveralls.io/github/digenaldo/goric?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/digenaldo/goric)](https://goreportcard.com/report/github.com/digenaldo/goric)
[![CircleCI](https://circleci.com/gh/digenaldo/goric/tree/master.svg?style=svg)](https://circleci.com/gh/digenaldo/goric/tree/master)

A Library in Go for validation of Brazilian documents.

### Installation

```bash
go get github.com/digenaldo/goric
```

### Import

```go
import (
	"github.com/digenaldo/goric"
)
```

### Usage

```go
package main

import (
	"github.com/digenaldo/goric"
)

func main() {
	//returns if cpf is valid
	b, err := goric.CpfIsValid("89576827086")

	//returns if cpf size is valid
	b1 := goric.CpfSizeIsValid("89576827086")

	//verifies that the CPF digits are valid
	b2 := goric.CpfDigitsValid("89576827086")
}
```

### To do list

- [x] CPF
- [ ] CNH
- [ ] Título de Eleitor
- [ ] PIS

### References

- [RIC](http://www.justica.gov.br/Acesso/governanca/ric)

### API Methods

To see the full list of API methods check out the [documentation on godoc.org](https://godoc.org/github.com/digenaldo/goric)
