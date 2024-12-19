package vertical

import (
	"fmt"
	"log"
	"strings"
)

var words = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'А', 'Б', 'В', 'Г', 'Д', 'Е', 'Ё', 'Ж', 'З', 'И', 'Й', 'К', 'Л', 'М', 'Н', 'О', 'П', 'Р', 'С', 'Т', 'У', 'Ф', 'Х', 'Ц', 'Ч', 'Ш', 'Щ', 'Ъ', 'Ы', 'Ь', 'Э', 'Ю', 'Я',
	'а', 'б', 'в', 'г', 'д', 'е', 'ё', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я',
	' ', '.', ',', ':', ';', '!', '?', '(', ')', '[', ']', '{', '}', '"', '-', '?',
}

func Encrypt(key1, text []rune) []rune {
	matrix := [][]int{}

	key := []rune{}
	keyMap := map[rune]bool{}
	for _, char := range key1 {
		if !keyMap[char] {
			key = append(key, char)
			keyMap[char] = true
		}
	}

	if strings.Contains(string(text), string(filler)) {
		fmt.Println("Filler changed!")
		for _, char := range words {
			if !strings.Contains(string(text), string(char)) {
				filler = char
				break
			}
		}
	}

	for len(text)%len(key) != 0 {
		text = append(text, filler)
	}

	n := len(text) / len(key)
	for i := 0; i < n; i++ {
		row := []int{}
		for j := 0; j < len(key); j++ {
			row = append(row, indexOf(text[i*len(key)+j], words))
		}
		matrix = append(matrix, row)
	}

	keyPrevIndex := make([]int, len(key))
	for i := range key {
		keyPrevIndex[i] = i
	}

	for i := 0; i < len(key); i++ {
		for j := 0; j < len(key); j++ {
			if indexOf(key[i], words) < indexOf(key[j], words) {
				key[i], key[j] = key[j], key[i]
				keyPrevIndex[i], keyPrevIndex[j] = keyPrevIndex[j], keyPrevIndex[i]
			}
		}
	}

	result := []rune{}
	for _, index := range keyPrevIndex {
		for _, row := range matrix {
			result = append(result, words[row[index]])
		}
	}

	return result
}

func appendUnique(slice []rune) []rune {
	newSlice := []rune{}
	mp := make(map[rune]bool)
	for _, r := range slice {
		if _, ok := mp[r]; !ok {
			newSlice = append(newSlice, r)
			mp[r] = true
		}
	}
	return newSlice
}

func Decrypt(key1, text []rune) []rune {
	key1 = appendUnique(key1)
	log.Println(string(key1))
	key := append([]rune{}, key1...)
	for len(text)%len(key) != 0 {
		text = append(text, filler)
	}

	keyPrevIndex := [][]int{}
	for i := range key {
		keyPrevIndex = append(keyPrevIndex, []int{i, i})
	}

	for i := 0; i < len(key); i++ {
		for j := 0; j < len(key); j++ {
			if indexOf(key[i], words) < indexOf(key[j], words) {
				key[i], key[j] = key[j], key[i]
				keyPrevIndex[i][0], keyPrevIndex[j][0] = keyPrevIndex[j][0], keyPrevIndex[i][0]
			}
		}
	}

	for i := 0; i < len(key); i++ {
		for j := 0; j < len(key); j++ {
			if keyPrevIndex[i][0] < keyPrevIndex[j][0] {
				keyPrevIndex[i], keyPrevIndex[j] = keyPrevIndex[j], keyPrevIndex[i]
			}
		}
	}

	n := len(text) / len(key)
	result := []rune{}
	for i := 0; i < n; i++ {
		for _, indices := range keyPrevIndex {
			result = append(result, text[indices[1]*n+i])
		}
	}

	finalResult := []rune{}
	for _, char := range result {
		if char != filler {
			finalResult = append(finalResult, char)
		}
	}

	return finalResult
}

var filler rune = '?'

func indexOf(char rune, alphabet []rune) int {
	for i, r := range alphabet {
		if r == char {
			return i
		}
	}
	return -1
}
