package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Data []struct {
	GlobalID int `json:"global_id"`
}

func main() {
	var decData Data

	// Вариант из локального файла
	// data, err := os.Open("data-20190514T0100.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer data.Close()
	//json.NewDecoder(data).Decode(&decData)

	// Вариант использования файла из сети
	resp, err := http.Get("https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&decData)

	sum := 0
	for _, v := range decData {
		sum += v.GlobalID
	}

	fmt.Println(sum)
}
