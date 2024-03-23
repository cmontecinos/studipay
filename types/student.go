package types

type Student struct {
	StudentID      int64  `json:"student_id"`
	Rut            string `json:"rut" validate:"required"`
	Name           string `json:"name" validate:"required"`
	Representative string `json:"representative" validate:"required"`
	Phone          string `json:"phone" validate:"required"`
	Email          string `json:"email" validate:"required"`
}
