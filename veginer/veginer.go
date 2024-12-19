package veginer

var words = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'А', 'Б', 'В', 'Г', 'Д', 'Е', 'Ё', 'Ж', 'З', 'И', 'Й', 'К', 'Л', 'М', 'Н', 'О', 'П', 'Р', 'С', 'Т', 'У', 'Ф', 'Х', 'Ц', 'Ч', 'Ш', 'Щ', 'Ъ', 'Ы', 'Ь', 'Э', 'Ю', 'Я',
	'а', 'б', 'в', 'г', 'д', 'е', 'ё', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я',
	' ', '.', ',', ':', ';', '!', '?', '(', ')', '[', ']', '{', '}', '"', '-', '?',
}

func Encrypt(key []rune, text []rune) string {
	makeLengthEqual(&key, text)

	for i := 0; i < len(text); i++ {
		plainChar := text[i]
		keyChar := key[i]
		encryptedChar := (indexOf(plainChar) + indexOf(keyChar)) % len(words)
		text[i] = words[encryptedChar]
	}
	return string(text)
}

func Decrypt(key []rune, text []rune) string {
	makeLengthEqual(&key, text)

	for i := 0; i < len(text); i++ {
		encryptedChar := text[i]
		keyChar := key[i]
		decryptedChar := (indexOf(encryptedChar) - indexOf(keyChar) + len(words)) % len(words)
		text[i] = words[decryptedChar]
	}
	return string(text)
}

func makeLengthEqual(key *[]rune, text []rune) {
	for i := len(*key); i < len(text); i++ {
		*key = append(*key, (*key)[i%len(*key)])
	}
}

func indexOf(char rune) int {
	for i, r := range words {
		if r == char {
			return i
		}
	}
	return -1
}
