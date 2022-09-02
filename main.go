package main

import (
	"fmt"
)

type RС4 struct {
	bytes [256]byte
	x     int
	y     int
}

//Генератор псевдослучайной последовательности (Pseudo-Random Generation Algorithm):
func (field RС4) numberGeneration() int {

	field.x = (field.x + 1) % 256
	field.y = (field.y + int(field.bytes[field.x])) % 256

	field.x, field.y = field.y, field.x

	result := int(field.bytes[(int(field.bytes[field.x])+int(field.bytes[field.y]))%256])

	return result

}

//Алгоритм ключевого расписания (Key-Scheduling Algorithm):
func (field RС4) keyGeneration(key string) {

	transformKey := []byte(key)
	lenghtKey := len(transformKey)

	for i := 0; i < 256; i++ {
		field.bytes[i] = byte(i)
	}

	j := 0

	for i := 0; i < 256; i++ {
		j = (j + int(field.bytes[i]) + int(transformKey[i%lenghtKey])) % 256
		field.bytes[i], field.bytes[j] = field.bytes[j], field.bytes[i]

	}

}

func encodingRC4()

func decodingRC4()

func mainMenu() {

	fmt.Printf("Выберите действие: \n")
	fmt.Printf("1) Кодирование: \n")
	fmt.Printf("2) Декодирование: \n")
	fmt.Printf("3) Выход: \n")

}

func main() {

	var (
		pick = '0'
		key  string
	)

	mainMenu()

	switch pick {

	case '1':
		fmt.Printf("Введите ключ")
		fmt.Scan(key)
		err := encodingRC4()
		if err != nil {

			panic("Ошибка кодирования:\n")

		} else {

			fmt.Print("Кодирование завершено успешно:\n")
		}

	case '2':

		fmt.Printf("Введите ключ")
		fmt.Scan(key)
		err := decodingRC4()
		if err != nil {

			panic("Ошибка декодирования:\n")

		} else {

			fmt.Print("Декодирование завершено успешно:\n")
		}

	case '3':

		return

	default:

		fmt.Printf("Не известная команда")
		return

	}

}
