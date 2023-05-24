package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

func main() {
	// Чтение данных из стандартного ввода
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	// Вариант для чтения из файла
	// file, err := os.Open("text.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// data, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Декодирование JSON данных в структуру Group
	var group Group
	err = json.Unmarshal(data, &group)
	if err != nil {
		log.Fatal(err)
	}

	// Рассчитываем общее количество оценок и количество студентов в группе
	totalCount := 0
	for _, student := range group.Students {
		totalCount += len(student.Rating)
	}

	// Рассчитываем среднее количество оценок для группы
	average := float64(totalCount) / float64(len(group.Students))

	// Создаем структуру для ответа
	response := struct {
		Average float64 `json:"Average"`
	}{
		Average: average,
	}

	// Кодируем результат в формате JSON
	output, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	// Выводим результат на стандартный вывод
	io.WriteString(os.Stdout, string(output))
}
