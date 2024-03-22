package persistence

import (
	"bigbytes.io/studipay/db"
	"bigbytes.io/studipay/types"
)

func CreatePayment(payment *types.Payment) error {
	client := db.GetClient()

	_, err := client.Exec("INSERT INTO payment (date, payment_datetime, amount, student_id) VALUES ($1, $2, $3, $4)", payment.Date, payment.PaymentDatetime, payment.Amount, payment.StudentID)
	if err != nil {
		return err
	}
	return nil
}

func GetPayment(date string) (*types.Payment, error) {
	client := db.GetClient()
	row := client.QueryRow("SELECT date, payment_datetime, amount, student_id FROM payment WHERE date = $1", date)

	payment := &types.Payment{}
	err := row.Scan(&payment.Date, &payment.PaymentDatetime, &payment.Amount, &payment.StudentID)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func GetPaymentByStudentId(studentId string) ([]*types.Payment, error) {
	client := db.GetClient()
	rows, err := client.Query("SELECT date, payment_datetime, amount, student_id FROM payment WHERE student_id = $1", studentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payments := []*types.Payment{}
	for rows.Next() {
		payment := &types.Payment{}
		err := rows.Scan(&payment.Date, &payment.PaymentDatetime, &payment.Amount, &payment.StudentID)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return payments, nil
}
