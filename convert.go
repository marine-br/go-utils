package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//CheckBin verifica se uma cadeia de strings 01010101 é um Binário válido
func CheckBin(n string) bool {
	validation := regexp.MustCompile(`^[01]{1,64}$`)
	return validation.MatchString(n)
}

//CheckDec verifica se uma cadeia de strings 99999 é um Decimal válido
func CheckDec(n string) bool {
	validation := regexp.MustCompile(`^[0-9]{1,64}$`)
	return validation.MatchString(n)
}

//CheckHex verifica se uma cadeia de strings ABCF0 é um Decimal válido
func CheckHex(n string) bool {
	validation := regexp.MustCompile(`^[0-9A-Fa-f]{1,64}$`)
	return validation.MatchString(n)
}

//PadLeft adiciona char à esquerda da string até chegar no qty
func PadLeft(s, char string, qty int) string {
	for len(s) < qty {
		s = char + s
	}
	return s
}

//PadRight adiciona char à esquerda da string até chegar no qty
func PadRight(s, char string, qty int) string {
	for len(s) < qty {
		s = s + char
	}
	return s
}

//UnpadLeft remove os 0 (zeros) à esquerda de uma string
func UnpadLeft(s string) string {
	validation := regexp.MustCompile(`^0+`)
	return validation.ReplaceAllString(s, "")
}

//Hex2Bin Transforma um HEXA string em um Binary String
func Hex2Bin(s string, digitos int) string {
	ui, _ := strconv.ParseUint(s, 16, 64)
	binString := fmt.Sprintf("%0"+strconv.Itoa(digitos)+"b", ui)
	return binString
}

//GetBitRL Retorna uma cadeia de BITs de um Binary String\
//Right to Left
func GetBitRL(bin string, ini int, end int) string {
	splited := strings.Split(bin, "")
	return strings.Join(splited[len(splited)-end:len(splited)-ini+1], "")
}

//GetBitLR Retorna uma cadeia de BITs de um Binary String\
//Left to Right
func GetBitLR(bin string, ini int, end int) string {
	splited := strings.Split(bin, "")
	return strings.Join(splited[ini-1:end], "")
}

//Hex2BinRangeLR Transforma um HEXA string em um Binary String e retorna os digitos dentro do ini e end
func Hex2BinRangeLR(s string, ini int, end int, digitos int) string {
	ui, _ := strconv.ParseUint(s, 16, 64)
	binString := fmt.Sprintf("%0"+strconv.Itoa(digitos)+"b", ui)
	splited := strings.Split(binString, "")
	return strings.Join(splited[ini-1:end], "")
}

//Hex2BinRangeRL Transforma um HEXA string em um Binary String e retorna os digitos dentro do ini e end
func Hex2BinRangeRL(s string, ini int, end int, digitos int) string {
	ui, _ := strconv.ParseUint(s, 16, 64)
	binString := fmt.Sprintf("%0"+strconv.Itoa(digitos)+"b", ui)
	splited := strings.Split(binString, "")
	return strings.Join(splited[len(splited)-end:len(splited)-ini+1], "")
}

//MergeMaps junta dois maps[string]string
func MergeMaps(a map[string]string, b map[string]string) map[string]string {
	for k, v := range b {
		a[k] = v
	}
	return a
}
