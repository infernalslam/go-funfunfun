package main

import (
	"fmt"
	"sort"
)

type profile struct {
	Name    string
	Surname string
	Age     int
}

type profiles []profile

func (p profiles) Find(f func(p profile) bool) profile {
	prof := profile{}
	for _, val := range p {
		if f(val) {
			prof = val
			break
		}
	}
	return prof
}

func (p profiles) Filter(f func(p profile) bool) (prof profiles) {
	for _, val := range p {
		if f(val) {
			prof = append(prof, val)
		}
	}
	return prof
}

func (p profiles) Map(f func(p profile) int) []int {
	ages := make([]int, len(p))
	for i, val := range p {
		ages[i] = f(val)
	}
	return ages
}

func (p profiles) ToConsole() {
	for _, val := range p {
		str := fmt.Sprintf("%s %s %d", val.Name, val.Surname, val.Age)
		fmt.Println(str)
	}
}

// sort

func (p profiles) Len() int {
	return len(p)
}

func (p profiles) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p profiles) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func main() {

	payload := profiles{
		{"T1", "T1", 27},
		{"T2", "T2", 22},
		{"T3", "T3", 21},
		{"T4", "T4", 19},
		{"T5", "T5", 20},
		{"T6", "T6", 11},
		{"T7", "T7", 12},
	}

	resp1 := payload.Filter(func(p profile) bool {
		return p.Age >= 21
	}).Map(func(p profile) int {
		return p.Age
	})

	resp2 := payload.Find(func(p profile) bool {
		return p.Name == "T3"
	})

	fmt.Println(resp1)
	fmt.Println(resp2)

	sort.Sort(payload)
	t := payload.Map(func(p profile) int {
		return p.Age
	})

	fmt.Println(t)

}
