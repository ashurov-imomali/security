package palifor

var words = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'А', 'Б', 'В', 'Г', 'Д', 'Е', 'Ё', 'Ж', 'З', 'И', 'Й', 'К', 'Л', 'М', 'Н', 'О', 'П', 'Р', 'С', 'Т', 'У', 'Ф', 'Х', 'Ц', 'Ч', 'Ш', 'Щ', 'Ъ', 'Ы', 'Ь', 'Э', 'Ю', 'Я',
	'а', 'б', 'в', 'г', 'д', 'е', 'ё', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я',
	' ', '.', ',', ':', ';', '!', '?', '(', ')', '[', ']', '{', '}', '"', '-', '?',
}

func createMatrix(key []rune) [12][12]rune {
	var matrix [12][12]rune
	used := make(map[rune]bool)
	x, y := 0, 0

	for _, c := range key {
		if !used[c] {
			matrix[x][y] = c
			used[c] = true
			y++
			if y == 12 {
				y = 0
				x++
			}
		}
	}

	for i := 0; i < len(words); i++ {
		c := words[i]
		if !used[c] {
			matrix[x][y] = c
			used[c] = true
			y++
			if y == 12 {
				y = 0
				x++
			}
		}
	}

	return matrix
}

func encryptPair(matrix [12][12]rune, a, b rune) (rune, rune) {
	var aRow, aCol, bRow, bCol int
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			if matrix[i][j] == a {
				aRow, aCol = i, j
			}
			if matrix[i][j] == b {
				bRow, bCol = i, j
			}
		}
	}

	if aRow == bRow {
		return matrix[aRow][(aCol+1)%12], matrix[bRow][(bCol+1)%12]
	} else if aCol == bCol {
		return matrix[(aRow+1)%12][aCol], matrix[(bRow+1)%12][bCol]
	} else {
		return matrix[aRow][bCol], matrix[bRow][aCol]
	}
}

func Encrypt(key []rune, text []rune) ([]rune, error) {
	filler = '?'
	keyMatrix := createMatrix(key)

	for _, c := range words {
		if !contains(text, c) {
			filler = c
			break
		}
	}

	var encrypted []rune
	for i := 0; i < len(text); i += 2 {
		a := text[i]
		b := filler
		if i+1 < len(text) {
			b = text[i+1]
		} else {
			i--
		}

		if a == b {
			b = filler
			i--
		}

		encA, encB := encryptPair(keyMatrix, a, b)
		encrypted = append(encrypted, encA, encB)
	}

	return encrypted, nil
}

func Decrypt(key []rune, text []rune) ([]rune, error) {
	keyMatrix := createMatrix(key)
	var decrypted []rune

	for i := 0; i < len(text); i += 2 {
		a := text[i]
		b := text[i+1]

		var aRow, aCol, bRow, bCol int
		for i := 0; i < 12; i++ {
			for j := 0; j < 12; j++ {
				if keyMatrix[i][j] == a {
					aRow, aCol = i, j
				}
				if keyMatrix[i][j] == b {
					bRow, bCol = i, j
				}
			}
		}

		if aRow == bRow {
			decrypted = append(decrypted, keyMatrix[aRow][(aCol+11)%12], keyMatrix[bRow][(bCol+11)%12])
		} else if aCol == bCol {
			decrypted = append(decrypted, keyMatrix[(aRow+11)%12][aCol], keyMatrix[(bRow+11)%12][bCol])
		} else {
			decrypted = append(decrypted, keyMatrix[aRow][bCol], keyMatrix[bRow][aCol])
		}
	}

	return decrypted, nil
}

func contains(slice []rune, c rune) bool {
	for _, v := range slice {
		if v == c {
			return true
		}
	}
	return false
}

var filler rune
