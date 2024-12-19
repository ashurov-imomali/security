package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"main/cesar"
	"main/palifor"
	"main/veginer"
	"main/vertical"
	"os"
)

func main() {
	encryptedText, err := readFile("input.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	fmt.Print("Введите ключ для шифрования/расшифровки: ")
	reader := bufio.NewReader(os.Stdin)
	key, _ := reader.ReadString('\n')

	key = key[:len(key)-1]

	keyRunes := []rune(key)
	textRunes := []rune(encryptedText)

	cesarDecrypted := cesar.Decrypt(3, textRunes)
	fmt.Println("Цезарь (расшифровка):", string(cesarDecrypted))
	cesarEncrypted := cesar.Encrypt(3, textRunes)
	fmt.Println("Цезарь (шифровка):", string(cesarEncrypted))

	verticalDecrypted := vertical.Decrypt([]rune("Ашурмалаевич"), textRunes)
	fmt.Println("Вертикальный шифр (расшифровка):", string(verticalDecrypted))
	verticalEncrypted := vertical.Encrypt(keyRunes, textRunes)
	fmt.Println("Вертикальный шифр (шифровка):", string(verticalEncrypted))
	return

	paliforDecrypted, _ := palifor.Decrypt(keyRunes, textRunes)
	fmt.Println("Плейфейр (расшифровка):", string(paliforDecrypted))
	paliforEncrypted, _ := palifor.Encrypt(keyRunes, textRunes)
	fmt.Println("Плейфейр (шифровка):", string(paliforEncrypted))

	veginerDecrypted := veginer.Decrypt(keyRunes, textRunes)
	fmt.Println("Виженер (расшифровка):", string(veginerDecrypted))
	veginerEncrypted := veginer.Encrypt(keyRunes, textRunes)
	fmt.Println("Виженер (шифровка):", string(veginerEncrypted))

	err = writeFile("cesar_decrypted.txt", cesarDecrypted)
	if err != nil {
		fmt.Println("Ошибка записи в файл cesar_decrypted.txt:", err)
	}

	err = writeFile("vertical_decrypted.txt", verticalDecrypted)
	if err != nil {
		fmt.Println("Ошибка записи в файл vertical_decrypted.txt:", err)
	}

	err = writeFile("palifor_decrypted.txt", paliforDecrypted)
	if err != nil {
		fmt.Println("Ошибка записи в файл palifor_decrypted.txt:", err)
	}

	err = writeFile("veginer_decrypted.txt", []rune(veginerDecrypted))
	if err != nil {
		fmt.Println("Ошибка записи в файл veginer_decrypted.txt:", err)
	}
}

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func writeFile(filename string, content []rune) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(string(content))
	return err
}
