package main

import (
	"math/rand"
)

//Generates a random string of size n
func randomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	runeArr := make([]rune, n)
	for i := range runeArr {
		runeArr[i] = letter[rand.Intn(len(letter))]
	}
	return string(runeArr)
}
