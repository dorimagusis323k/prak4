package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Grades []float64
}

var students = make(map[string]Student)

func main() {
	for {
		var choice int
		fmt.Println("1. Добавить")
		fmt.Println("2. Фильтр")
		fmt.Println("3. Выход")
		fmt.Scanln(&choice)

		if choice == 1 {
			var name string
			fmt.Print("Имя: ")
			fmt.Scanln(&name)

			var grade float64
			var grades []float64
			for i := 0; i < 3; i++ {
				fmt.Printf("Оценка %d: ", i+1)
				fmt.Scanln(&grade)
				grades = append(grades, grade)
			}

			students[name] = Student{Name: name, Grades: grades}
		} else if choice == 2 {
			var threshold float64
			fmt.Print("Порог: ")
			fmt.Scanln(&threshold)

			for name, student := range students {
				sum := 0.0
				for _, grade := range student.Grades {
					sum += grade
				}
				avg := sum / float64(len(student.Grades))
				if avg < threshold {
					fmt.Println(name)
				}
			}
		} else if choice == 3 {
			break
		}
	}
}
