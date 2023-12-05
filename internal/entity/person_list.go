package entity

import (
	"sort"
)

type PersonList struct {
	Persons []*Person
}

func (l *PersonList) Push(x interface{}) {
	l.Persons = append(l.Persons, x.(*Person))
}

func (l *PersonList) Swap(i, j int) {
	l.Persons[i], l.Persons[j] = l.Persons[j], l.Persons[i]
}

func (l *PersonList) SortByName() {
	sort.Slice(l.Persons, func(i, j int) bool {
		return l.Persons[i].Name < l.Persons[j].Name
	})
}

func (l *PersonList) SortByAge() {
	sort.Slice(l.Persons, func(i, j int) bool {
		return l.Persons[i].Age < l.Persons[j].Age
	})
}

func NewPersonList() *PersonList {
	return &PersonList{}
}
