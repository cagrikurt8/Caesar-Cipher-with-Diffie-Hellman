package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var words []string

	for i := 0; i < 7; i++ {
		scanner.Scan()
		words = append(words, scanner.Text())
	}

	g, err := strconv.Atoi(words[2])
	p, err1 := strconv.Atoi(words[6])

	if err != nil {
		log.Fatal(err)
	}

	if err1 != nil {
		log.Fatal(err1)
	}

	b, c, A, B := 7, 1, 0, 0

	for i := 0; i < b; i++ {
		B = c * g % p
		c = B
	}

	fmt.Println("OK")

	for i := 0; i < 3; i++ {
		scanner.Scan()

		if i == 2 {
			A, err = strconv.Atoi(scanner.Text())

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Println("B is ", B)

	S := math.Pow(float64(A), float64(b))
	secret := math.Mod(S, float64(p))
	shift := math.Mod(secret, 26)

	fmt.Println(encrypt("Will you marry me?", int(shift)))

	var answer string

	fmt.Scan(&answer)

	if strings.Contains(answer, ",") {
		fmt.Println(encrypt("Great!", int(shift)))
	} else if strings.Contains(answer, "'") {
		fmt.Println(encrypt("What a pity!", int(shift)))
	}
}

func encrypt(message string, shift int) string {
	ALPHABET := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	var encrypted = ""

	for i := 0; i < len(message); i++ {
		letter := message[i]
		letter_idx := letterIndex(ALPHABET, string(letter))

		if letter_idx != -1 {
			letter_idx += shift

			if letter_idx > 25 {
				letter_idx -= 26
			}

			if unicode.IsUpper(rune(letter)) {
				encrypted += ALPHABET[letter_idx]
			} else {
				encrypted += strings.ToLower(ALPHABET[letter_idx])
			}

		} else {
			encrypted += string(letter)
		}
	}
	return encrypted
}

func letterIndex(alp [26]string, letter string) int {
	for i := 0; i < len(alp); i++ {
		if alp[i] == strings.ToUpper(letter) {
			return i
		}
	}
	return -1
}
