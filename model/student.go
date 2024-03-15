package types

type Student struct {
	StudentID      int    `json:"student_id"`
	Rut            string `json:"rut"`
	Name           string `json:"name"`
	Representative string `json:"representative"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
}
