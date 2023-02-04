package dto

import "github.com/s-kmmr/golang-v2-sample-clean-architecture/domain/model"

// EmployeeShowList 従業員一覧参照用DTO
type EmployeeShowList struct {
	employeeNum   string
	departmentNum uint
	firstName     string
	lastName      string
}

// NewEmployeeShowList コンストラクタ
// このDTOを使用して、従業員一覧参照（検索）処理へパラメータを運ぶ
func NewEmployeeShowList(_employeeNum string, _departmentNum uint, _firstName string, _lastName string) *EmployeeShowList {
	return &EmployeeShowList{
		employeeNum:   _employeeNum,
		departmentNum: _departmentNum,
		firstName:     _firstName,
		lastName:      _lastName,
	}
}

func (l *EmployeeShowList) MakeEmployeeSearchCriteria() *model.EmployeeSearchCriteria {
	return model.NewEmployeeSearchCriteria(
		l.employeeNum,
		l.departmentNum,
		l.firstName,
		l.lastName,
	)
}
