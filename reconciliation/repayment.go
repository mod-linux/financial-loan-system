package reconciliation

import (
	"linus/lms/finance"
	"time"
)

type Payment struct {
	Amount                      float64
	PaymentDate                 time.Time
	Foreclosure                 bool
	ReconciliationConfiguration PaymentReconciliationConfiguration
	LoanRepaymentObject         LoanRepaymentObject
}

func (payment Payment) bounceReco(input interface{}) interface{} {
	return input
}

func (payment Payment) penalReco(input interface{}) interface{} {
	return input
}

type PaymentReconciliationConfiguration func(interface{}) interface{}

func (payment Payment) getRecoMethod() []PaymentReconciliationConfiguration {

	return []PaymentReconciliationConfiguration{
		payment.bounceReco,
		payment.penalReco,
	}
}

type LoanRepaymentObject struct {
	Schedule []finance.Schedule
}

func (payment Payment) Reconcile() {
	ExecuteInstructions(payment.getRecoMethod())
}

func ExecuteInstructions(instructions []PaymentReconciliationConfiguration) {
	var input interface{}
	for _, instruction := range instructions {
		input = instruction(input)
	}
}
