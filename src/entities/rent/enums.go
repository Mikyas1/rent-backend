package rent

import "database/sql/driver"

type RentRequestStatus string
type RentStatus string
type PaymentType string

func (c *RentRequestStatus) Scan(value interface{}) error {
	*c = RentRequestStatus(value.(string))
	return nil
}

func (c RentRequestStatus) Value() (driver.Value, error) {
	return string(c), nil
}

func (c *RentStatus) Scan(value interface{}) error {
	*c = RentStatus(value.(string))
	return nil
}

func (c RentStatus) Value() (driver.Value, error) {
	return string(c), nil
}

func (c *PaymentType) Scan(value interface{}) error {
	*c = PaymentType(value.(string))
	return nil
}

func (c PaymentType) Value() (driver.Value, error) {
	return string(c), nil
}

const (
	Pending  RentRequestStatus = "PENDING"
	Accepted RentRequestStatus = "ACCEPTED"
	Rejected RentRequestStatus = "REJECTED"

	OnGoing   RentStatus = "ONGOING"
	Completed RentStatus = "COMPLETED"

	Weekly   PaymentType = "WEEKLY"
	Monthly  PaymentType = "MONTHLY"
	Annually PaymentType = "ANNUALLY"
)

var AllRentRequestStatues []RentRequestStatus = []RentRequestStatus{Pending, Accepted, Rejected}
var AllRentStatues []RentStatus = []RentStatus{OnGoing, Completed}
var AllPaymentTypes []PaymentType = []PaymentType{Weekly, Monthly, Annually}
