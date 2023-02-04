package repository

import (
	"context"

	"github.com/s-kmmr/golang-v2-sample-clean-architecture/domain/model"
)

// Employee インターフェース
type Employee interface {
	Retrieve(ctx context.Context, criteria *model.EmployeeSearchCriteria) ([]*model.Employee, error)
}
