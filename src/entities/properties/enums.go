package properties

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type Country string
type Region string
type City string
type Feature string
type Features []Feature
type PropertyType string
type PropertyStatus string

func (c *Country) Scan(value interface{}) error {
	*c = Country(value.(string))
	return nil
}

func (c Country) Value() (driver.Value, error) {
	return string(c), nil
}

func (c Country) String() string {
	return string(c)
}

func (c *Region) Scan(value interface{}) error {
	*c = Region(value.(string))
	return nil
}

func (c Region) Value() (driver.Value, error) {
	return string(c), nil
}

func (c Region) String() string {
	return string(c)
}

func (c *City) Scan(value interface{}) error {
	*c = City(value.(string))
	return nil
}

func (c City) Value() (driver.Value, error) {
	return string(c), nil
}

func (c City) String() string {
	return string(c)
}

func (c *Feature) Scan(value interface{}) error {
	*c = Feature(value.(string))
	return nil
}

func (c Feature) Value() (driver.Value, error) {
	return string(c), nil
}

func (c Feature) String() string {
	return string(c)
}

func (c Feature) StringWithQuot() string {
	return "\"" + c.String() + "\""
}

func (ar Features) Value() (driver.Value, error) {
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

func (ar *Features) Scan(value interface{}) error {
	switch value := value.(type) {
	case string:
		var res Features
		value = value[1 : len(value)-1] // Remove '{' and '}'
		sts := strings.Split(value, ",")
		for _, st := range sts {
			py := Feature(st)
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

func (c *PropertyType) Scan(value interface{}) error {
	*c = PropertyType(value.(string))
	return nil
}

func (c PropertyType) Value() (driver.Value, error) {
	return string(c), nil
}

func (c PropertyType) String() string {
	return string(c)
}

func (c *PropertyStatus) Scan(value interface{}) error {
	*c = PropertyStatus(value.(string))
	return nil
}

func (c PropertyStatus) Value() (driver.Value, error) {
	return string(c), nil
}

func (c PropertyStatus) String() string {
	return string(c)
}

const (
	Ethiopia Country = "ETHIOPIA"
	Sudan    Country = "SUDAN"
	Eritrea  Country = "ERITREA"

	AddisAbaba City = "ADDISABABA"
	Baherdar   City = "BAHERDAR"
	Adama      City = "ADAMA"
	Hawasa     City = "HAWASA"

	Oromia      Region = "OROMIA"
	Amhara      Region = "AMHARA"
	AddisAbabaR Region = "ADDISABABA"

	AirCondition Feature = "AIRCONDITION"
	Garden       Feature = "GARDEN"
	Pool         Feature = "POOL"
	Balcony      Feature = "BALCONY"
	WaterTank    Feature = "WATERTANK"
	Generator    Feature = "GENERATOR"
	Security     Feature = "SECURITY"
	Internet     Feature = "INTERNET"
	WaterPump    Feature = "WATERPUMP"
	Garage       Feature = "GARAGE"

	PendingApproval PropertyStatus = "PENDINGAPPROVAL"
	Approved        PropertyStatus = "APPROVED"
	Rented          PropertyStatus = "RENTED"
	Rejected        PropertyStatus = "REJECTED"

	Apartment PropertyType = "APARTMENT"
	Villa     PropertyType = "VILLA"
	Studio    PropertyType = "STUDIO"
)

var AllFeatures Features = Features{AirCondition, Garden, Pool, Balcony, WaterTank, Generator, Security, Internet, WaterPump, Garage}
var AllRegions []Region = []Region{Oromia, Amhara, AddisAbabaR}
var AllCities []City = []City{Baherdar, Adama, Hawasa, AddisAbaba}
var AllCountries []Country = []Country{Ethiopia, Eritrea, Sudan}
var AllPropertyStatus []PropertyStatus = []PropertyStatus{PendingApproval, Approved, Rented, Rejected}
var AllPropertyTypes []PropertyType = []PropertyType{Apartment, Villa, Studio}
