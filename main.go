package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/ianbrito/ordenacao-de-dados-em-go/internal/entity"
)

func main() {
	src := os.Args[1]
	dst := os.Args[2]

	rows, err := read(src)
	if err != nil {
		panic(err)
	}

	list := entity.NewPersonList()

	file := createFile(dst)
	defer file.Close()

	fs, err := file.Stat()
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		list.Push(entity.ParseRow(row))
	}

	if fs.Name() == "ordenado_por_nome.csv" {
		list.SortByName()
	}

	if fs.Name() == "ordenado_por_idade.csv" {
		list.SortByAge()
	}

	records := [][]string{
		{"Nome", "Idade", "Pontuação"},
	}

	for _, person := range list.Persons {
		records = append(records, []string{
			person.Name,
			strconv.Itoa(person.Age),
			strconv.Itoa(person.Score),
		})
	}

	writer := csv.NewWriter(file)
	err = writer.WriteAll(records)

	if err != nil {
		log.Fatal(err)
	}
}

func createFile(name string) *os.File {
	_, err := os.Stat(name)
	if os.IsExist(err) {
		panic(err)
	}

	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	return file
}

func read(fileName string) ([][]string, error) {

	file, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	if _, err := reader.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := reader.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
