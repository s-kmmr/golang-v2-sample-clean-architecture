package dto

import "github.com/s-kmmr/golang-v2-sample-clean-architecture/domain/model"

// EmployeeShowListResult  従業員一覧参照結果用DTO
// controllerへ渡す
type EmployeeShowListResult struct {
	id            uint
	employeeNum   string
	departmentNum uint
	firstName     string
	lastName      string
}

// EmployeeShowListResults 従業員一覧参照結果スライスDTO
type EmployeeShowListResults []*EmployeeShowListResult

func newEmployeeShowListResult(_id uint, _employeeNum string, _departmentNum uint, _firstName string, _lastName string) *EmployeeShowListResult {
	return &EmployeeShowListResult{
		id:            _id,
		employeeNum:   _employeeNum,
		departmentNum: _departmentNum,
		firstName:     _firstName,
		lastName:      _lastName,
	}
}

// NewEmployeeShowListResults コンストラクタ
func NewEmployeeShowListResults(e []*model.Employee) []*EmployeeShowListResult {
	if len(e) < 1 {
		return nil
	}
	r := make([]*EmployeeShowListResult, len(e))
	for i, v := range e {
		r[i] = newEmployeeShowListResult(
			v.ID(),
			v.EmployeeNum(),
			v.DepartmentNum(),
			v.FirstName(),
			v.LastName(),
		)
	}
	return r
}
