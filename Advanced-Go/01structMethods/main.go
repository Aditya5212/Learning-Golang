package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func (s *Student) updateAge(age int) {
	s.age = age
}

func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, marks := range s.grades {
		sum += marks
	}
	return float32(sum) / float32(len(s.grades))
}

func (s Student) getMaxGrades() int {
	max := s.grades[0]
	for _, marks := range s.grades {
		if marks > max {
			max = marks
		}
	}
	return max
}

func main() {
	student1 := Student{
		"Aditya",
		[]int{50, 60, 70, 80, 90},
		15,
	}
	student2 := Student{
		"Yash",
		[]int{90, 95, 70, 80, 90},
		20,
	}
	student1.updateAge(20)
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println("Average Grade of Aditya:", student1.getAverageGrade())
	fmt.Println("Average Grade of Yash:", student2.getAverageGrade())
	fmt.Println("Max Grade of Aditya:", student1.getMaxGrades())
	fmt.Println("Max Grade of Yash:", student2.getMaxGrades())
}
