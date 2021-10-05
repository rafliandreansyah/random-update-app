package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

type Object struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

const FILENAME = "data/data.json"

func readFile() (string, int, int) {

	var dataStatus Object

	fmt.Println("Read file data...")

	data, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		return fmt.Sprintf("Error Readfile: %s", err.Error()), 0, 0
	}

	err = json.Unmarshal(data, &dataStatus)
	if err != nil {
		return fmt.Sprintf("Error ConvertData: %s", err.Error()), 0, 0
	}

	if dataStatus.Status.Wind < 6 || dataStatus.Status.Water < 5 {
		return fmt.Sprintf("Status Aman"), dataStatus.Status.Wind, dataStatus.Status.Water
	} else if (dataStatus.Status.Water >= 6 || dataStatus.Status.Water <= 8) || (dataStatus.Status.Wind >= 7 || dataStatus.Status.Wind <= 15) {
		return fmt.Sprintf("Status Siaga"), dataStatus.Status.Wind, dataStatus.Status.Water
	} else {
		return fmt.Sprintf("Bahaya"), dataStatus.Status.Wind, dataStatus.Status.Water
	}

}

func writeFile() {

	rangeMin := 1
	rangeMax := 100
	windRandom := rand.Intn(rangeMax-rangeMin+1) + rangeMin
	waterRandom := rand.Intn(rangeMax-rangeMin+1) + rangeMin

	status := Status{
		Water: waterRandom,
		Wind:  windRandom,
	}
	dataStatus := Object{
		Status: status,
	}

	file, err := os.Create(FILENAME)
	if err != nil {
		fmt.Printf("Error Write1 => %v", err.Error())
		return
	}

	jsonByte, err := json.Marshal(dataStatus)
	if err != nil {
		fmt.Printf("Error Write2 => %v", err.Error())
		return
	}

	n, err := file.WriteString(string(jsonByte))
	if err != nil {
		fmt.Printf("Error Write3 => %v", err.Error())
		return
	}
	fmt.Println(n)
	fmt.Println(string(jsonByte))
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		writeFile()
		status, wind, water := readFile()

		data := map[string]string{
			"Status": status,
			"Wind":   strconv.Itoa(wind),
			"Water":  strconv.Itoa(water),
		}

		t, err := template.ParseFiles("index.html")
		if err != nil {
			fmt.Println("Error", err.Error())
			return
		}

		t.Execute(w, data)

	})
	http.ListenAndServe(":8000", nil)
}
