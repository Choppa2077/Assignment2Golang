package main

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name         string
	Age          int
	Email        string
	DepartmentID uint // Foreign key
}

type Course struct {
	gorm.Model
	Title        string
	Description  string
	InstructorID uint // Foreign key
}

type Department struct {
	gorm.Model
	Name string
	// Remove Students field
}

type Enrollment struct {
	gorm.Model
	StudentID uint
	CourseID  uint
}

type Instructor struct {
	gorm.Model
	Name  string
	Email string
	// Remove Courses field
}
