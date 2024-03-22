package types

type Payment struct {
	Date            string  `json:"date"`
	PaymentDatetime string  `json:"payment_datetime"`
	Amount          float64 `json:"amount"`
	StudentID       int     `json:"student_id"`
}
