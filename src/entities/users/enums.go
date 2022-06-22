package users

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type UserStatus string

func (us *UserStatus) Scan(value interface{}) error {
	*us = UserStatus(value.(string))
	return nil
}

func (us UserStatus) Value() (driver.Value, error) {
	return string(us), nil
}

func (us UserStatus) String() string {
	return string(us)
}

type AdminRole string

func (ar *AdminRole) Scan(value interface{}) error {
	*ar = AdminRole(value.(string))
	return nil
}

func (ar AdminRole) Value() (driver.Value, error) {
	return string(ar), nil
}

func (ar AdminRole) String() string {
	return string(ar)
}

func (ar AdminRole) StringWithQuot() string {
	return "\"" + ar.String() + "\""
}

type AdminRoles []AdminRole

func (ar AdminRoles) Value() (driver.Value, error) {
	if ar == nil {
		return nil, nil
	}
	if n := len(ar); n > 0 {
		s := "{"
		s = s + ar[0].StringWithQuot()
		for i := 1; i < n; i++ {
			s = s + "," + ar[i].StringWithQuot()
		}
		return s + "}", nil
	}
	return "{}", nil
}

func (ar *AdminRoles) Scan(value interface{}) error {
	switch value := value.(type) {
	case string:
		var res AdminRoles
		value = value[1 : len(value)-1] // Remove '{' and '}'
		sts := strings.Split(value, ",")
		for _, st := range sts {
			py := AdminRole(st)
			res = append(res, py)
		}
		*ar = res
		return nil
	case nil:
		*ar = nil
		return nil
	}
	return fmt.Errorf("cannot convert %T to AdminRoles", value)
}

const (
	Active       UserStatus = "ACTIVE"
	Suspended    UserStatus = "SUSPENDED"
	Blocked      UserStatus = "BLOCKED"
	AdminUserCE  AdminRole  = "ADMINUSERCE"
	RegularUserE AdminRole  = "REGULARUSERE"
)

var AllAdminRoles AdminRoles = AdminRoles{AdminUserCE, RegularUserE}
