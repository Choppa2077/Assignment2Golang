package main

func CreateStudent(name string, age int, email string, departmentID uint) {
	student := Student{
		Name:         name,
		Age:          age,
		Email:        email,
		DepartmentID: departmentID,
	}
	DB.Create(&student)
}

// GetStudentByID retrieves a student by their ID
func GetStudentByID(studentID uint) (Student, error) {
	var student Student
	result := DB.First(&student, studentID)
	return student, result.Error
}

// UpdateStudentByID updates a student's information by their ID
func UpdateStudentByID(studentID uint, name string, age int, email string) error {
	var student Student
	result := DB.First(&student, studentID)
	if result.Error != nil {
		return result.Error
	}

	student.Name = name
	student.Age = age
	student.Email = email

	DB.Save(&student)
	return nil
}

// DeleteStudentByID deletes a student by their ID
func DeleteStudentByID(studentID uint) error {
	var student Student
	result := DB.Delete(&student, studentID)
	return result.Error
}
