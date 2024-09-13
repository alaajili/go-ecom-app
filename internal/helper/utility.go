package helper

import (
	"crypto/rand"
	"strconv"
)

func RandomNumbers(length int) (int, error) {

	const numbers = "0123456789"

	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return 0, err
	}

	numLen := len(numbers)
	for i := 0; i < length; i++ {
		buffer[i] = numbers[int(buffer[i])%numLen]
	}

	return strconv.Atoi(string(buffer))
}