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

func Unpack(txt string) (string, error) {
	symbols, err := parse(txt)
	if err != nil {
		return "", err
	}
	return toString(symbols), nil
}

func toString(symbols []symbol) string {
	var txtBuilder strings.Builder
	for _, symbol := range symbols {
		str := strings.Repeat(symbol.value, symbol.count)
		txtBuilder.WriteString(str)
	}
	return txtBuilder.String()
}

func parse(txt string) ([]symbol, error) {
	reg := regexp.MustCompile(`([A-Za-z][0]+)`)
	txt = reg.ReplaceAllString(txt, "")
	reg = regexp.MustCompile(`([A-Za-z][?\d]*)`)
	foundStrings := reg.FindAllString(txt, -1)
	preSymbol := ""
	txtLength := 0
	regOnlyNumber := regexp.MustCompile(`([\d]+)`)
	symbols := []symbol{}
	for _, currentSymbol := range foundStrings {
		txtLength += len(currentSymbol)
		countRepeatStr := regOnlyNumber.FindString(currentSymbol)
		symbolStr := regOnlyNumber.ReplaceAllString(currentSymbol, "")
		if countRepeatStr == "" {
			preSymbol = currentSymbol
			symbols = append(symbols, symbol{symbolStr, 1})
			continue
		}
		if preSymbol == symbolStr {
			return []symbol{}, ErrInvalidString
		}
		countRepeat, _ := strconv.Atoi(countRepeatStr)
		symbols = append(symbols, symbol{symbolStr, countRepeat})
	}
	if txtLength != len(txt) {
		return []symbol{}, ErrInvalidString
	}
	return symbols, nil
}
