package main

import (
	"fmt"
	"time"
)

func main() {
	// Чтение строки времени со стандартного ввода
	var input string
	fmt.Scan(&input)

	// Парсинг строки времени в тип time.Time
	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		fmt.Println("Ошибка при парсинге строки времени:", err)
		return
	}

	// Форматирование и вывод времени в формате UnixDate
	fmt.Println(t.Format(time.UnixDate))
}
