package service

import (
	"bigbytes.io/studipay/persistence"
	"bigbytes.io/studipay/types"
)

func CreateStudent(student *types.Student) (*types.Student, error) {
	created, err := persistence.CreateStudent(student)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func GetStudent(rut string) (*types.Student, error) {
	student, err := persistence.GetStudent(rut)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func UpdateStudent(student *types.Student) error {
	err := persistence.UpdateStudent(student)
	if err != nil {
		return err
	}
	return nil
}
