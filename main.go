package main

import (
	"errors"
  "github.com/gin-gonic/gin"
	"fmt"
	"strconv"
)

type Student struct {
	Id   int
	Name string
	Age  int
}

var students []Student

func getStudents() []Student {
	return students
}

func getStudentByID(id int) *Student {
	for _, student := range students {
		if student.Id == id {
			return &student
		}
	}

	return nil
}

func addStudent(student Student) error {
	for _, s := range students {
		if s.Id == student.Id {
			return errors.New("ID already exists")
		} else if s.Name == student.Name {
			return errors.New("Name already exists")
		}
	}

	students = append(students, student)
	return nil
}

func updateStudent(student Student) error {
	for i, s := range students {
		if s.Id == student.Id {
			students = append(students[:i], students[i+1:]...)
			students = append(students, student)
			return nil
		}
	}

	return errors.New("No student found")
}

func deleteStudent(id int) error {
	for i, student := range students {
		if student.Id == id {
			students = append(students[:i], students[i+1:]...)
			return nil
		}
	}

	return errors.New("No student found")
}

func printStudents(students []Student) {
	for _, student := range students {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", student.Id, student.Name, student.Age)
	}
}


func main() {
	addStudent(Student{Id: 1, Name: "Student One", Age: 20})
	addStudent(Student{Id: 2, Name: "Student Two", Age: 22})
	addStudent(Student{Id: 3, Name: "Student Three", Age: 21})
	addStudent(Student{Id: 4, Name: "Student Four", Age: 23})

	r := gin.Default()
	r.GET("/students", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
			"data": getStudents(),
		})
	})
	r.GET("/student/:id", func(c *gin.Context) {
		id_str := c.Param("id")
		id, err := strconv.Atoi(id_str)
		
		if err != nil {
			c.JSON(200, gin.H{
				"message": "error",
				"data": "Request error",
			})
		} else {
			student := getStudentByID(id)

			if student == nil {
				c.JSON(200, gin.H{
					"message": "error",
					"data": "No student found",
				})
			} else {
				c.JSON(200, gin.H{
					"data": student,
					"message": "success",
				})
			}
		}
	})
	r.POST("/student", func(c *gin.Context) {
		var student Student
		c.BindJSON(&student)		
		
		err := addStudent(student)
		if err != nil {
			c.JSON(200, gin.H{
				"message": "error",
				"data": "Something went wrong",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "success",
				"data": getStudents(),
			})
		}
	})
	r.PUT("/student", func(c *gin.Context) {
		var student Student
		c.BindJSON(&student)
		
		err := updateStudent(student)
		if err != nil {
			c.JSON(200, gin.H{
				"message": "error",
				"data": "No student found",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "success",
				"data": getStudents(),
			})
		}
	})
	r.DELETE("/student/:id", func(c *gin.Context) {
		id_str := c.Param("id")
		id, err := strconv.Atoi(id_str)

		if err != nil {
			c.JSON(200, gin.H{
				"message": "error",
				"data": "Request error",
			})
		} else {
			err_two := deleteStudent(id)
			
			if err_two != nil {
				c.JSON(200, gin.H{
					"message": "error",
					"data": "Something went wrong",
				})
			} else {
				c.JSON(200, gin.H{
					"message": "success",
					"data": getStudents(),
				})
			}
		}
	})
	r.Run()
}
