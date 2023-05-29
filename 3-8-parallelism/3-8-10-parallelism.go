package main

import "fmt"

func removeDuplicates(inputStream, outputStream chan string) {
	prevValue := ""
	for value := range inputStream {
		if value != prevValue {
			outputStream <- value
		}
		prevValue = value
	}
	close(outputStream)
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)

	go removeDuplicates(inputStream, outputStream)

	input := []string{"apple", "banana", "banana", "apple", "apple", "orange", "orange"}

	// Отправляем значения во входной канал
	go func() {
		for _, v := range input {
			inputStream <- v
		}
		close(inputStream)
	}()

	// Получаем значения из выходного канала
	for v := range outputStream {
		fmt.Println(v)
	}
}
