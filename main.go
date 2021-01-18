package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Дано 𝑛-байтовое  данное  (2,  4,  6,  8,  16,  24,  32  байта).
// Задана  перестановка (1,8,23,0,16,...).
// Написать    функцию,    выполняющую    эту    байтовую перестановку.
// Ввод и вывод данных организуйте в 16-ой системе счисления.

const base = 16

func main() {
	var str string
	var perm string

	fmt.Print("enter num with spaces: ")

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str = sc.Text()

	data := translate(str)

	fmt.Print("enter permutation with spaces: ")

	sc.Scan()
	perm = sc.Text()

	permF := translate(perm)

	permutation, err := permutation(data, permF)
	if err != nil {
		log.Fatal(err)
	}

	printRes(permutation)
}

func printRes(data []uint) {
	for _, v := range data {
		fmt.Printf("%x ", v)
	}

	fmt.Println("")
}

func translate(str string) []uint {
	res := strings.Fields(str)

	ress := make([]uint, len(res))

	for i, v := range res {
		num, err := strconv.ParseUint(v, base, 64)
		if err != nil {
			log.Fatal("not hex")
		}

		ress[i] = uint(num)
	}

	return ress
}

func permutation(data []uint, permF []uint) ([]uint, error) {
	if !isValidData(data) {
		return nil, errors.New("invalid data len")
	}

	dataLen := len(data)
	dataMaxByte := uint(dataLen - 1)

	res := make([]uint, dataLen)

	for i, p := range permF {
		if p > dataMaxByte {
			return nil, fmt.Errorf("permutation not available: maxByte=%d, permByte=%d", dataMaxByte, p)
		}

		res[p] = data[i]
	}

	return res, nil
}

func isValidData(data []uint) bool {
	switch len(data) {
	case 2, 4, 6, 8, 16, 24, 32:
		return true
	default:
		return false
	}
}
