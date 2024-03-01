I have a structs Student, Course, Enrollment, Instructor, Department
So between a many-to-many relationship, also  a one-to-many relationship between Instructors and Courses.
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
