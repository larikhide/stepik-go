package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при чтении ввода:", err)
		return
	}
	// Удаление символов новой строки из входной строки
	input = strings.TrimSpace(input)

	// Парсинг входной строки в объект времени
	t, err := time.Parse("2006-01-02 15:04:05", input)
	if err != nil {
		fmt.Println("Ошибка при парсинге строки времени:", err)
		return
	}

	// Получение времени обеда (13:00) для сравнения
	lunchTime := time.Date(t.Year(), t.Month(), t.Day(), 13, 0, 0, 0, t.Location())

	// Проверка, нужно ли перенести событие на следующий день
	if t.Hour() >= lunchTime.Hour() {
		t = t.Add(24 * time.Hour)
	}

	// Вывод результата в нужном формате
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}
