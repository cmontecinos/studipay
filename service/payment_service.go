package service

import (
	"bigbytes.io/studipay/persistence"
	"bigbytes.io/studipay/types"
)

func CreatePayment(payment *types.Payment) error {
	err := persistence.CreatePayment(payment)
	if err != nil {
		return err
	}
	return nil
}

func GetPayment(paymentId string) (*types.Payment, error) {
	payment, err := persistence.GetPayment(paymentId)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func GetPaymentsByStudentID(studentID string) ([]*types.Payment, error) {
	payments, err := persistence.GetPaymentByStudentId(studentID)
	if err != nil {
		return nil, err
	}
	return payments, nil
}
