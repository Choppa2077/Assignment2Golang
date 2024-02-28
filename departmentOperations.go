package main

func CreateDepartment(name string) {
	department := Department{
		Name: name,
	}
	DB.Create(&department)
}

// GetDepartmentByID retrieves a department by its ID
func GetDepartmentByID(departmentID uint) (Department, error) {
	var department Department
	result := DB.First(&department, departmentID)
	return department, result.Error
}

// UpdateDepartmentByID updates a department's name by its ID
func UpdateDepartmentByID(departmentID uint, newName string) error {
	var department Department
	result := DB.First(&department, departmentID)
	if result.Error != nil {
		return result.Error
	}

	department.Name = newName

	DB.Save(&department)
	return nil
}

// DeleteDepartmentByID deletes a department by its ID
func DeleteDepartmentByID(departmentID uint) error {
	var department Department
	result := DB.Delete(&department, departmentID)
	return result.Error
}
