package main

func GetStudentsByDepartment(departmentID uint) ([]Student, error) {
	var students []Student
	result := DB.Where("department_id = ?", departmentID).Find(&students)
	return students, result.Error
}

func GetCoursesByInstructor(instructorID uint) ([]Course, error) {
	var courses []Course
	result := DB.Where("instructor_id = ?", instructorID).Find(&courses)
	return courses, result.Error
}

func GetEnrollmentsByStudent(studentID uint) ([]Enrollment, error) {
	var enrollments []Enrollment
	result := DB.Where("student_id = ?", studentID).Find(&enrollments)
	return enrollments, result.Error
}

// GetTotalStudentsInDepartment finds the total number of students in a department
func GetTotalStudentsInDepartment(departmentID uint) (int64, error) {
	var count int64
	result := DB.Model(&Student{}).Where("department_id = ?", departmentID).Count(&count)
	return count, result.Error
}

// GetCoursesWithEnrollmentCount retrieves the list of courses with the number of enrolled students
func GetCoursesWithEnrollmentCount() ([]map[string]interface{}, error) {
	var courses []Course
	result := DB.Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}

	var coursesWithEnrollmentCount []map[string]interface{}

	for _, course := range courses {
		var count int64
		DB.Model(&Enrollment{}).Where("course_id = ?", course.ID).Count(&count)

		courseWithCount := map[string]interface{}{
			"Course":           course,
			"EnrolledStudents": count,
		}

		coursesWithEnrollmentCount = append(coursesWithEnrollmentCount, courseWithCount)
	}

	return coursesWithEnrollmentCount, nil
}

// EnrollStudentInCourse enrolls a student in a course within a transaction
func EnrollStudentInCourse(studentID, courseID uint) error {
	tx := DB.Begin()

	// Check for errors during the transaction
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Enroll the student in the course
	enrollment := Enrollment{
		StudentID: studentID,
		CourseID:  courseID,
	}

	if err := tx.Create(&enrollment).Error; err != nil {
		tx.Rollback()
		return err
	}

	// You can perform additional updates or checks within the same transaction if needed

	// Commit the transaction if everything is successful
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
