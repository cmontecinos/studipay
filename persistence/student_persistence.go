package persistence

import (
	"database/sql"
	"errors"
	"fmt"

	"bigbytes.io/studipay/db"
	"bigbytes.io/studipay/types"
)

func CreateStudent(student *types.Student) (*types.Student, error) {
	client := db.GetClient()

	var studentID int64
	err := client.QueryRow("INSERT INTO student (rut, name, representative, phone, email) VALUES ($1, $2, $3, $4, $5) RETURNING student_id", student.Rut, student.Name, student.Representative, student.Phone, student.Email).Scan(&studentID)
	if err != nil {
		return nil, err
	}
	student.StudentID = studentID

	return student, nil
}

var ErrStudentNotFound = errors.New("student not found")

func GetStudent(rut string) (*types.Student, error) {
	client := db.GetClient()
	row := client.QueryRow("SELECT rut, name, representative, phone, email FROM student WHERE rut = $1", rut)

	student := &types.Student{}
	err := row.Scan(&student.Rut, &student.Name, &student.Representative, &student.Phone, &student.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Student with rut %s not found\n", rut)
			return nil, ErrStudentNotFound
		}
		return nil, err
	}

	return student, nil
}

func UpdateStudent(student *types.Student) error {
	client := db.GetClient()
	_, err := client.Exec("UPDATE student SET name = $1, representative = $2, phone = $3, email = $4 WHERE rut = $5", student.Name, student.Representative, student.Phone, student.Email, student.Rut)
	if err != nil {
		return err
	}

	return nil
}
