package repository

import (
	"context"

	"github.com/s-kmmr/golang-v2-sample-clean-architecture/domain/model"
	"github.com/s-kmmr/golang-v2-sample-clean-architecture/domain/repository"
)

type employee struct {
	// todo client
}

// NewEmployee コンストラクタ
func NewEmployee() repository.Employee {
	return employee{}
}
func (e employee) Retrieve(ctx context.Context, criteria *model.EmployeeSearchCriteria) ([]*model.Employee, error) {
	// todo impl
	return []*model.Employee{}, nil
}
