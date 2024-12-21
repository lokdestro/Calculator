package service

import "calc/pkg/calc"

type Service interface {
	Calc(expr string) (float64, error)
}

type service struct {
	calc calc.Calculator
}

func New(c calc.Calculator) Service {
	return &service{calc: c}
}

func (s *service) Calc(expr string) (float64, error) {
	return s.calc.Calc(expr)
}
