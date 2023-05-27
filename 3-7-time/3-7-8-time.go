package main

import (
	"fmt"
	"time"
)

const baseTime int64 = 1589570165

func main() {
	// По условию нужно прочитать из стандартного ввода, а потом преобразовать в time.Duration()
	// в таком случае нужно было бы делать так: duraion := time.Duration(min*time.Minutes)+.....
	var min, sec time.Duration
	fmt.Scanf("%d мин. %d сек.", &min, &sec)
	duration := min*time.Minute + sec*time.Second

	// Создание базовой даты из подсказки
	// Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.
	baseDate := time.Unix(baseTime, 0).UTC()
	resultDate := baseDate.Add(duration)

	// Вывод в формате UnixDate
	fmt.Println(resultDate.Format(time.UnixDate))
}
