package usecase

import (
	"context"
	"fmt"

	"github.com/s-kmmr/golang-v2-sample-clean-architecture/domain/repository"
	"github.com/s-kmmr/golang-v2-sample-clean-architecture/usecase/dto"
)

type EmployeeShowList interface {
	ShowList(ctx context.Context, param *dto.EmployeeShowList) (dto.EmployeeShowListResults, error)
}

type employeeShowList struct {
	er repository.Employee
}

func NewEmployeeShowList(_er repository.Employee) EmployeeShowList {
	return &employeeShowList{
		er: _er,
	}
}

// ShowList 検索条件を元に、従業員一覧を取得する
func (l *employeeShowList) ShowList(ctx context.Context, param *dto.EmployeeShowList) (dto.EmployeeShowListResults, error) {
	res, err := l.er.Retrieve(ctx, param.MakeEmployeeSearchCriteria())
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve show list %w", err)
	}

	return dto.NewEmployeeShowListResults(res), nil
}
