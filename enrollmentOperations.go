package main

func CreateEnrollment(studentID, courseID uint) {
	enrollment := Enrollment{
		StudentID: studentID,
		CourseID:  courseID,
	}
	DB.Create(&enrollment)
}

func GetEnrollmentByID(enrollmentID uint) (Enrollment, error) {
	var enrollment Enrollment
	result := DB.First(&enrollment, enrollmentID)
	return enrollment, result.Error
}

func UpdateEnrollmentByID(enrollmentID, studentID, courseID uint) error {
	var enrollment Enrollment
	result := DB.First(&enrollment, enrollmentID)
	if result.Error != nil {
		return result.Error
	}

	enrollment.StudentID = studentID
	enrollment.CourseID = courseID

	DB.Save(&enrollment)
	return nil
}

func DeleteEnrollmentByID(enrollmentID uint) error {
	var enrollment Enrollment
	result := DB.Delete(&enrollment, enrollmentID)
	return result.Error
}
