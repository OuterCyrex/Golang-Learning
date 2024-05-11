package main

import (
	"fmt"
	"os"
)

type StudentId int
type ClassId int
type Name string
type Student struct {
	StudentId
	ClassId
	Name
}
type MainSystem struct {
	Array []*Student
}

const MAX = 0x3f3f3f

var flag = false

func Welcome() {
	fmt.Println("Operate Menu:")
	fmt.Println("1.Add Student Data")
	fmt.Println("2.Show Student Data")
	fmt.Println("3.Delete Student Data")
	fmt.Println("4.Exit Student Data")
}
func NewSystem() *MainSystem {
	return &MainSystem{
		Array: make([]*Student, 0, MAX),
	}
}
func (s *MainSystem) ShowData() {
	if len(s.Array) == 0 {
		fmt.Println("No Data")
	} else {
		for i := 0; i < len(s.Array); i++ {
			fmt.Println("Student Id:", s.Array[i].StudentId, "ClassId:", s.Array[i].ClassId, "Name:", s.Array[i].Name)
		}
	}
	fmt.Println("Input Any Key to continue...")
	var AnyKey int
	_, scan := fmt.Scan(&AnyKey)
	if scan != nil {
		return
	}
}
func (s *MainSystem) AddData() {
	fmt.Println("Input 0 to Exit")
	for {
		fmt.Println("Please Input Student Id:")
		var id StudentId
		_, scan := fmt.Scan(&id)
		if scan != nil {
			return
		}
		if id == 0 {
			return
		}
		for i := 0; i < len(s.Array); i++ {
			if s.Array[i].StudentId == id {
				fmt.Println("Student Id Already Exists")
				flag = true
			}
		}
		if flag {
			flag = false
			continue
		}
		fmt.Println("Please Input Class Id:")
		var class ClassId
		_, err := fmt.Scan(&class)
		if err != nil {
			return
		}
		fmt.Println("Please Input Name:")
		var name Name
		_, err = fmt.Scan(&name)
		if err != nil {
			return
		}
		newStu := new(Student)
		newStu.StudentId = id
		newStu.ClassId = class
		newStu.Name = name
		s.Array = append(s.Array, newStu)
		fmt.Println("Complete!")
	}
}
func (s *MainSystem) DeleteData() {
	fmt.Println("Please Input Student Id:")
	var id StudentId
	_, err := fmt.Scan(&id)
	if err != nil {
		return
	}
	for i := 0; i < len(s.Array); i++ {
		if s.Array[i].StudentId == id {
			s.Array = append(s.Array[:i], s.Array[i+1:]...)
			fmt.Println("Complete!")
			fmt.Println("Input Any Key to continue...")
			var AnyKey int
			_, err2 := fmt.Scan(&AnyKey)
			if err2 != nil {
				return
			}
			return
		}
	}
	fmt.Println("Invalid Student Id")
	fmt.Println("Input Any Key to continue...")
	var AnyKey int
	_, err = fmt.Scan(&AnyKey)
	if err != nil {
		return
	}
}
func main() {
	fmt.Println("Welcome to Student DataBase System")
	StuArray := NewSystem()
	for {
		Welcome()
		fmt.Println("Please Input your operate:")
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}
		fmt.Println("Your Input:", input)
		switch input {
		case 1:
			StuArray.AddData()
		case 2:
			StuArray.ShowData()
		case 3:
			StuArray.DeleteData()
		case 4:
			fmt.Println("Exiting Student Data...")
			os.Exit(0)
		default:
			fmt.Println("Invalid Input!")
		}
	}
}
