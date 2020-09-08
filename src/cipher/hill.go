package cipher

import (
	"math"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func modInverse(a, m int) int {
	for a < 0 {
		a += m
	}
	a = a % m
	result := 1
	for x := 1; x < m; x++ {
		if (a*x)%m == 1 {
			result = x
		}
	}
	return result
}

func Hill(plainText string, key string) string {

	rows := strings.Split(key, ",")
	matrixSize := len(rows)
	matrix := mat.NewDense(matrixSize, matrixSize, nil)

	for i, row := range rows {
		elements := strings.Split(row, "|")
		for j, element := range elements {
			value, _ := strconv.ParseFloat(element, 32)
			matrix.Set(i, j, value)
		}
	}

	plainText = processString(plainText)
	cipherText := ""

	for i := 0; i < len(plainText); i += matrixSize {

		vector := mat.NewVecDense(matrixSize, nil)
		for j := 0; j < matrixSize; j++ {
			if i+j < len(plainText) {
				vector.SetVec(j, float64(int(plainText[i+j])-65))
			} else {
				vector.SetVec(j, 0)
			}
		}

		vector.MulVec(matrix, vector)

		for j := 0; j < matrixSize; j++ {
			if i+j < len(plainText) {
				nextCharacter := int(math.Round(vector.At(j, 0))) % 26
				if nextCharacter < 0 {
					nextCharacter += 26
				}
				cipherText += string(rune(nextCharacter + 65))
			}
		}
	}

	return cipherText
}

func DecipherHill(cipherText string, key string) string {
	rows := strings.Split(key, ",")
	matrixSize := len(rows)
	matrix := mat.NewDense(matrixSize, matrixSize, nil)

	for i, row := range rows {
		elements := strings.Split(row, "|")
		for j, element := range elements {
			value, _ := strconv.ParseFloat(element, 32)
			matrix.Set(i, j, value)
		}
	}

	detMatrix := mat.Det(matrix)
	detMatrixModInv := modInverse(int(math.Round(detMatrix)), 26)
	matrix.Inverse(matrix)
	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			value := float64(int(math.Round((math.Round(matrix.At(i, j) * detMatrix * float64(detMatrixModInv))))) % 26)
			if value < 0 {
				value += 26
			}
			matrix.Set(i, j, value)
		}
	}

	cipherText = processString(cipherText)
	plainText := ""

	for i := 0; i < len(cipherText); i += matrixSize {

		vector := mat.NewVecDense(matrixSize, nil)
		for j := 0; j < matrixSize; j++ {
			if i+j < len(cipherText) {
				vector.SetVec(j, float64(int(cipherText[i+j])-65))
			} else {
				vector.SetVec(j, 0)
			}
		}

		vector.MulVec(matrix, vector)

		for j := 0; j < matrixSize; j++ {
			if i+j < len(cipherText) {
				nextCharacter := int(math.Round(vector.At(j, 0))) % 26
				if nextCharacter < 0 {
					nextCharacter += 26
				}
				plainText += string(rune(nextCharacter + 65))
			}
		}
	}

	return plainText
}
