package main

func CreateCourse(title string, description string, instructorID uint) {
	course := Course{
		Title:        title,
		Description:  description,
		InstructorID: instructorID,
	}
	DB.Create(&course)
}

func GetCourseByID(courseID uint) (Course, error) {
	var course Course
	result := DB.First(&course, courseID)
	return course, result.Error
}

func UpdateCourseByID(courseID uint, title string, description string) error {
	var course Course
	result := DB.First(&course, courseID)
	if result.Error != nil {
		return result.Error
	}

	course.Title = title
	course.Description = description

	DB.Save(&course)
	return nil
}

func DeleteCourseByID(courseID uint) error {
	var course Course
	result := DB.Delete(&course, courseID)
	return result.Error
}
