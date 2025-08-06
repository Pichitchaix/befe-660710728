package main

import (
	"errors"
	"fmt"
)

type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Year  int     `json:"year"`
	GPA   float64 `json:"gpa"`
}

func (s *Student) IsHonor() bool {
	return s.GPA >= 3.50
}

func (s *Student) Validate() error {
	if s.Name == "" {
		return errors.New("Name is required")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("year must be between 1-4")
	}
	if s.GPA < 0 || s.GPA > 4 {
		return errors.New("GPA must be between 0-4")
	}
	return nil
}

func main() {
	// var st Student = Student{ID: "1", Name: "Pichitchai", Email: "pichitchai7679@gmail.com", Year: 4, GPA: 3.76}
	Students := []Student{
		{ID: "1", Name: "Pichitchai", Email: "pichitchai7679@gmail.com", Year: 4, GPA: 3.76},
		{ID: "2", Name: "Pichitchai2", Email: "pichitchai27679@gmail.com", Year: 4, GPA: 1.76},
	}
	newStudent := Student{ID: "3", Name: "Pichitchai3", Email: "pichitchai37679@gmail.com", Year: 4, GPA: 3.56}
	Students = append(Students, newStudent)

	for i, student := range Students {
		fmt.Printf("%dHonor %v\n", i, student.IsHonor())
		fmt.Printf("%dValidation %v\n", i, student.Validate())
	}

}
