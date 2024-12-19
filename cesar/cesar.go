package cesar

func Encrypt(n int, text []rune) []rune {
	for i := 0; i < len(text); i++ {
		text[i] = rune((int(text[i]) + n) % '𐀀')
	}
	return text
}

func Decrypt(n int, text []rune) []rune {
	for i := 0; i < len(text); i++ {
		text[i] = rune((int(text[i]) - n + '𐀀') % '𐀀')
	}
	return text
}
