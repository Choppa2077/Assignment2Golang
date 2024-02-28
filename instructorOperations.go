package main

// CreateInstructor creates a new instructor record in the database
func CreateInstructor(name string, email string) {
	instructor := Instructor{
		Name:  name,
		Email: email,
	}
	DB.Create(&instructor)
}

// GetInstructorByID retrieves an instructor by their ID
func GetInstructorByID(instructorID uint) (Instructor, error) {
	var instructor Instructor
	result := DB.First(&instructor, instructorID)
	return instructor, result.Error
}

// UpdateInstructorByID updates an instructor's information by their ID
func UpdateInstructorByID(instructorID uint, name string, email string) error {
	var instructor Instructor
	result := DB.First(&instructor, instructorID)
	if result.Error != nil {
		return result.Error
	}

	instructor.Name = name
	instructor.Email = email

	DB.Save(&instructor)
	return nil
}

// DeleteInstructorByID deletes an instructor by their ID
func DeleteInstructorByID(instructorID uint) error {
	var instructor Instructor
	result := DB.Delete(&instructor, instructorID)
	return result.Error
}
