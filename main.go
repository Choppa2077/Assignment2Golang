package main

func main() {
	// Initialize the database
	InitDB()

	// Perform operations on the Department table
	//CreateDepartment("Economics")
	// Retrieve a department by ID
	//department, err := GetDepartmentByID(1)
	//if err == nil {
	//	fmt.Println("Retrieved Department:", department)
	//} else {
	//	fmt.Println("Error:", err)
	//}

	// Update a department by ID
	//err := UpdateDepartmentByID(1, "Information Technology")
	//if err == nil {
	//	fmt.Println("Department Updated Successfully")
	//} else {
	//	fmt.Println("Error:", err)
	//}

	// Delete a department by ID
	//err := DeleteDepartmentByID(2)
	//if err == nil {
	//	fmt.Println("Department Deleted Successfully")
	//} else {
	//	fmt.Println("Error:", err)
	//}

	CreateStudent("Arup", 20, "arup@gmail.com", 2)
	// Retrieve a student by ID
	//student, err := GetStudentByID(1)
	//if err == nil {
	//	fmt.Println("Retrieved Student:", student)
	//} else {
	//	fmt.Println("Error:", err)
	//}

	// Update a student by ID
	//err := UpdateStudentByID(1, "Mukan", 19, "sdudent777@gmail.com")
	//if err == nil {
	//	fmt.Println("Student Updated Successfully")
	//} else {
	//	fmt.Println("Error:", err)
	//}

	// Delete a student by ID
	//err := DeleteStudentByID(2)
	//if err == nil {
	//	fmt.Println("Student Deleted Successfully")
	//} else {
	//	fmt.Println("Error:", err)
	//}
	//CreateInstructor("Professor Bob", "professor999@gmail.com")

	// Retrieve an instructor by ID
	//instructor, err := GetInstructorByID(1)
	//if err == nil {
	//	fmt.Println("Retrieved Instructor:", instructor)
	//} else {
	//	fmt.Println("Error:", err)
	//}
	//

	// Update an instructor by ID
	//err := UpdateInstructorByID(1, "Instructor Smith", "instructor77@gmail.com")
	//if err == nil {
	//	fmt.Println("Instructor Updated Successfully")
	//} else {
	//	fmt.Println("Error:", err)
	//}

	// Delete an instructor by ID
	//err := DeleteInstructorByID(3)
	//if err == nil {
	//	fmt.Println("Instructor Deleted Successfully")
	//} else {
	//	fmt.Println("Error:", err)
	//}

	//CreateCourse("Data analysis", "Intro to data analysis", 2)

	// Retrieve a course by ID
	//course, err := GetCourseByID(1)
	//if err == nil {
	//	fmt.Println("Retrieved Course:", course)
	//} else {
	//	fmt.Println("Error:", err)
	//}

	// Update a course by ID
	//err := UpdateCourseByID(1, "Programming 102", "Intermediate Programming")
	//if err == nil {
	//	fmt.Println("Course Updated Successfully")
	//} else {
	//	fmt.Println("Error:", err)
	//}

	// Delete a course by ID
	//err := DeleteCourseByID(1)
	//if err == nil {
	//	fmt.Println("Course Deleted Successfully")
	//} else {
	//	fmt.Println("Error:", err)
	//}

	//CreateEnrollment(3, 4) // Assuming studentID and courseID

	// Retrieve an enrollment by ID
	//enrollment, err := GetEnrollmentByID(1)
	//if err == nil {
	//	fmt.Println("Retrieved Enrollment:", enrollment)
	//} else {
	//	fmt.Println("Error:", err)
	//}
	//
	// Update an enrollment by ID
	//err := UpdateEnrollmentByID(2, 1, 2) // Assuming new studentID and courseID
	//if err == nil {
	//	fmt.Println("Enrollment Updated Successfully")
	//} else {
	//	fmt.Println("Error:", err)
	//}

	// Delete an enrollment by ID
	//err := DeleteEnrollmentByID(1)
	//if err == nil {
	//	fmt.Println("Enrollment Deleted Successfully")
	//} else {
	//	fmt.Println("Error:", err)
	//}

	// Retrieve all students from a specific department
	//students, err := GetStudentsByDepartment(2) // Assuming departmentID
	//if err == nil {
	//	fmt.Println("Students in Department:", students)
	//} else {
	//	fmt.Println("Error:", err)
	//}

	// Find courses taught by a particular instructor
	//courses, err := GetCoursesByInstructor(2) // Assuming instructorID
	//if err == nil {
	//	fmt.Println("Courses taught by Instructor:", courses)
	//} else {
	//	fmt.Println("Error:", err)
	//}

	// Get all enrollments for a specific student
	//enrollments, err := GetEnrollmentsByStudent(1) // Assuming studentID
	//if err == nil {
	//	fmt.Println("Enrollments for Student:", enrollments)
	//} else {
	//	fmt.Println("Error:", err)
	//}

	// Get the total number of students in a department
	//totalStudents, err := GetTotalStudentsInDepartment(2) // Assuming departmentID
	//if err == nil {
	//	fmt.Println("Total Students in Department:", totalStudents)
	//} else {
	//	fmt.Println("Error:", err)
	//}

	//coursesWithEnrollmentCount, err := GetCoursesWithEnrollmentCount()
	//if err == nil {
	//	fmt.Println("Courses with Enrollment Count:", coursesWithEnrollmentCount)
	//} else {
	//	fmt.Println("Err", err)
	//}

	//Enroll a student in a course within a transaction
	//err := EnrollStudentInCourse(4, 3) // Assuming studentID and courseID
	//if err == nil {
	//	fmt.Println("Student enrolled in the course successfully")
	//} else {
	//	fmt.Println("Error:", err)
	//}
}
