package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Grades contains JSON Keys(properties)
type Grades struct {
	LastName  string
	FirstName string
	Ssn       string
	Test1     int
	Test2     int
	Test3     int
	Test4     int
	Final     int
	Grade     string
}

func main() {
	csvFile, err := os.Open("./grades.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var grade Grades
	for index, each := range csvData {
		grade.LastName = each[0]
		grade.FirstName = each[1]
		grade.Ssn = each[2]
		grade.Test1, _ = strconv.Atoi(each[3])
		grade.Test2, _ = strconv.Atoi(each[4])
		grade.Test3, _ = strconv.Atoi(each[5])
		grade.Test4, _ = strconv.Atoi(each[6])
		grade.Final, _ = strconv.Atoi(each[7])
		grade.Grade = each[8]
		// Convert to JSON
		jsonData, err := json.Marshal(grade)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(string(jsonData))
		fileName := fmt.Sprintf("csv-row-%d.json", index)
		jsonFile, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
		}
		jsonFile.Write(jsonData)
		jsonFile.Close()
	}

}
