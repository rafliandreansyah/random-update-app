package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Object struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

const FILENAME = "data/data.json"

func readFile() string {

	var dataStatus Object

	fmt.Println("Read file data...")

	data, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		return fmt.Sprintf("Error Readfile: ", err.Error())
	}

	err = json.Unmarshal(data, &dataStatus)
	if err != nil {
		return fmt.Sprintf("Error ConvertData: ", err.Error())
	}

	if dataStatus.Status.Wind < 6 || dataStatus.Status.Water < 5 {
		return fmt.Sprintf("Status Aman")
	} else if (dataStatus.Status.Water >= 6 || dataStatus.Status.Water <= 8) || (dataStatus.Status.Wind >= 7 || dataStatus.Status.Wind <= 15) {
		return fmt.Sprintf("Status Siaga")
	} else {
		return fmt.Sprintf("Bahaya")
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
	writeFile()
	for _ = range time.Tick(time.Second * 2) {
		fmt.Println("Reload data")
		writeFile()
		// readFile()
	}
}
