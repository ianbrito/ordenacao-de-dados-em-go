package entity

import "strconv"

type Person struct {
	Name  string
	Age   int
	Score int
}

func ParseRow(row []string) *Person {
	name := row[0]

	age, err := strconv.ParseInt(row[1], 10, 0)
	if err != nil {
		panic(err)
	}

	score, err := strconv.ParseInt(row[2], 10, 0)
	if err != nil {
		panic(err)
	}

	return &Person{
		Name:  name,
		Age:   int(age),
		Score: int(score),
	}
}
