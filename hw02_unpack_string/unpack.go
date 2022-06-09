package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

type symbol struct {
	value string
	count int
}

var regSymbol = regexp.MustCompile(`\\n\d*|\\{2}\d*|\\\d*|[a-z]\d*|a-z`)

func Unpack(txt string) (string, error) {
	parts, err := parseStrToPart(txt)
	if err != nil {
		return "", err
	}
	symbols, err := parsePartsToSymbols(parts)
	if err != nil {
		return "", err
	}
	err = validateSymbols(symbols)
	if err != nil {
		return "", err
	}
	return symbolsToString(symbols), nil
}

func parseStrToPart(txt string) ([]string, error) {
	res := regSymbol.FindAllString(txt, -1)
	partLength := 0
	for _, str := range res {
		partLength += len(str)
	}
	if len(txt) != partLength {
		return []string{}, ErrInvalidString
	}
	return res, nil
}

func parsePartsToSymbols(collect []string) ([]symbol, error) {
	var err error
	var count int
	var countStr string
	var foundSymbol string
	var firstChar string
	var secondChar string
	var doubleChars string
	symbols := []symbol{}
	for _, item := range collect {
		if len(item) == 1 {
			symbols = append(symbols, symbol{item, 1})
			continue
		}
		firstChar = item[:1]
		secondChar = item[1:2]
		doubleChars = item[:2]
		switch {
		case doubleChars == `\n`:
			countStr = item[2:]
			foundSymbol = doubleChars
		case firstChar == `\`:
			countStr = item[2:]
			foundSymbol = secondChar
		default:
			foundSymbol = firstChar
			countStr = item[1:]
		}
		if countStr == "" {
			count = 1
		} else {
			count, err = strconv.Atoi(countStr)
			if err != nil {
				return []symbol{}, err
			}
		}
		symbols = append(symbols, symbol{foundSymbol, count})
	}
	return symbols, nil
}

func validateSymbols(symbols []symbol) error {
	preSymbol := symbol{value: "", count: 0}
	for _, currentSymbol := range symbols {
		if currentSymbol.count > 1 && preSymbol.value == currentSymbol.value {
			return ErrInvalidString
		}
		preSymbol = currentSymbol
	}
	return nil
}

func symbolsToString(symbols []symbol) string {
	var txtBuilder strings.Builder
	for _, symbol := range symbols {
		str := strings.Repeat(symbol.value, symbol.count)
		txtBuilder.WriteString(str)
	}
	return txtBuilder.String()
}
