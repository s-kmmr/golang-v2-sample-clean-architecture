package model

import (
	"fmt"
)

// Employee 従業員を表現するモデル
type Employee struct {
	id uint
	// 従業員番号
	employeeNum string
	// 所属部署番号
	departmentNum uint
	// 名前
	firstName string
	// 苗字
	lastName string
}

// NewEmployee コンストラクタ
// 従業員番号・部署番号・苗字・名前は必須
func NewEmployee(_id uint, _employeeNum string, _departmentNum uint, _firstName string, _lastName string) (*Employee, error) {
	if len(_employeeNum) < 1 || _departmentNum < 1 || len(_firstName) < 1 || len(_lastName) < 1 {
		return nil, fmt.Errorf("missing required parameter. employee num = %s department num = %d first name = %s last name = %s", _employeeNum, _departmentNum, _firstName, _lastName)
	}
	return &Employee{
		id:            _id,
		employeeNum:   _employeeNum,
		departmentNum: _departmentNum,
		firstName:     _firstName,
		lastName:      _lastName,
	}, nil
}

func (e *Employee) ID() uint {
	return e.id
}

func (e *Employee) EmployeeNum() string {
	return e.employeeNum
}

func (e *Employee) DepartmentNum() uint {
	return e.departmentNum
}

func (e *Employee) FirstName() string {
	return e.firstName
}

func (e *Employee) LastName() string {
	return e.lastName
}

// EmployeeSearchCriteria 従業員一覧検索条件
type EmployeeSearchCriteria struct {
	// 従業員番号
	employeeNum string
	// 所属部署番号
	departmentNum uint
	// 名前
	firstName string
	// 苗字
	lastName string
}

// NewEmployeeSearchCriteria コンストラクタ
func NewEmployeeSearchCriteria(_employeeNum string, _departmentNum uint, _firstName string, _lastName string) *EmployeeSearchCriteria {
	return &EmployeeSearchCriteria{
		employeeNum:   _employeeNum,
		departmentNum: _departmentNum,
		firstName:     _firstName,
		lastName:      _lastName,
	}
}

func (c *EmployeeSearchCriteria) EmployeeNum() string {
	return c.employeeNum
}

func (c *EmployeeSearchCriteria) DepartmentNum() uint {
	return c.departmentNum
}

func (c *EmployeeSearchCriteria) FirstName() string {
	return c.firstName
}

func (c *EmployeeSearchCriteria) LastName() string {
	return c.lastName
}
