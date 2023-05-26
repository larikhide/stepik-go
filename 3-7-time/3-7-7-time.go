package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// Разделение входной строки на две даты
	dates := strings.Split(scanner.Text(), ",")

	// Преобразование строковых дат в тип Time
	layout := "02.01.2006 15:04:05"
	date1, err := time.Parse(layout, dates[0])
	if err != nil {
		log.Fatal(err)
	}
	date2, err := time.Parse(layout, dates[1])
	if err != nil {
		log.Fatal(err)
	}

	// Вычисление продолжительности периода
	diff := date1.Sub(date2)
	if diff < 0 {
		diff *= -1
	}

	// Вывод результатов
	fmt.Println(diff)
}
