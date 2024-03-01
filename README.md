I have a structs Student, Course, Enrollment, Instructor, Department
So between  Students and Courses a many-to-many relationship, also  a one-to-many relationship between Instructors and Courses.Between Students and Departments one to many realtionship.
In Connection.go  i implement InitDB() funtion and inside of this define my dsn for my postgresql database 
dsn := "host=localhost user=postgres password=m511 dbname=gormcrud port=5432 sslmode=disable TimeZone=Asia/Almaty"
before this i already create database named gormcrud in pgAdmin.

and then after defining dsn call agorm.Open() and put inside dsn and config then call a DB.AutoMigrate() and put inside pointers for each structure that i have.
After runninh this tables are created in pgAdmin.
(Notice that for each operation we every time need call InitDB() inside of main.go

All operations like a CREATE,READ,UPDATE,DELETE for each table is implemented separately, for example CRUD functionality for Students table is in studentOperations.go file and so on.
For creation we get name,email,age, and departmentId for relationship between student and department.In this function we create varibale student that metach to student structure. and  call DB.Create() and inside of this function we put a pointer to studetn variable.
So and in other CREATE funtions we do the same take the parameters create a variable that match to structure and call DB.Create() and put pointer.

In GET funtions we take the parametrs and put this to DB.First() function call ,also put a id.

In UPDATE funtions do like in the get call the DB.First() for finding what we should update and  reassign properties of structure like this for example:
student.Name = name
student.Age = age
student.Email = email
after this we call DB.Save() and put pointer to this

In DELETE we just get id of thing taht we should delete and call DB.Delete() and put inside a pinter to structure and id .

All queries that write in doc of assignment i implement in qyering.go file

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
