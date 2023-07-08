package repo

import (
	"math/rand"

	"github.com/google/uuid"
)

func genToken() string {
	randomUUID, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	return randomUUID.String()
}

func genString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
