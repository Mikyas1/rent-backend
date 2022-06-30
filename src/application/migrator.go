package application

import (
	"rent/src/datasource"
	"rent/src/entities/message"
	"rent/src/entities/properties"
	"rent/src/entities/rent"
	"rent/src/entities/users"
)

func SetUpTable(Db datasource.Database) {
	// user status
	_ = Db.ProvideDb().Exec("DROP TYPE IF EXISTS user_status;")
	_ = Db.ProvideDb().Exec("CREATE TYPE user_status AS ENUM (" +
		"'ACTIVE'," +
		"'SUSPENDED'," +
		"'BLOCKED');")

	// admin roles
	_ = Db.ProvideDb().Exec("DROP TYPE IF EXISTS admin_roles;")
	_ = Db.ProvideDb().Exec("CREATE TYPE admin_roles AS ENUM (" +
		"'ADMINUSERCE'," +
		"'REGULARUSERE');")

	// countries
	_ = Db.ProvideDb().Exec("DROP TYPE IF EXISTS countries;")
	_ = Db.ProvideDb().Exec("CREATE TYPE countries AS ENUM (" +
		"'ETHIOPIA'," +
		"'SUDAN'," +
		"'ERITREA');")

	// cities
	_ = Db.ProvideDb().Exec("DROP TYPE IF EXISTS cities;")
	_ = Db.ProvideDb().Exec("CREATE TYPE cities AS ENUM (" +
		"'ADDISABABA'," +
		"'BAHERDAR'," +
		"'ADAMA'," +
		"'HAWASA');")

	// regions
	_ = Db.ProvideDb().Exec("DROP TYPE IF EXISTS regions;")
	_ = Db.ProvideDb().Exec("CREATE TYPE regions AS ENUM (" +
		"'OROMIA'," +
		"'AMHARA'," +
		"'ADDISABABA');")

	// Feature
	_ = Db.ProvideDb().Exec("DROP TYPE IF EXISTS features;")
	_ = Db.ProvideDb().Exec("CREATE TYPE features AS ENUM (" +
		"'AIRCONDITION'," +
		"'GARDEN'," +
		"'POOL'," +
		"'BALCONY'," +
		"'WATERTANK'," +
		"'GENERATOR'," +
		"'SECURITY'," +
		"'INTERNET'," +
		"'WATERPUMP'," +
		"'GARAGE');")

	// PropertyStatus
	_ = Db.ProvideDb().Exec("DROP TYPE IF EXISTS property_status;")
	_ = Db.ProvideDb().Exec("CREATE TYPE property_status AS ENUM (" +
		"'PENDINGAPPROVAL'," +
		"'APPROVED'," +
		"'REJECTED'," +
		"'RENTED');")

	// PropertyType
	_ = Db.ProvideDb().Exec("DROP TYPE IF EXISTS property_types;")
	_ = Db.ProvideDb().Exec("CREATE TYPE property_types AS ENUM (" +
		"'APARTMENT'," +
		"'VILLA'," +
		"'STUDIO');")

	// RentRequestStatus
	_ = Db.ProvideDb().Exec("DROP TYPE IF EXISTS rent_request_status;")
	_ = Db.ProvideDb().Exec("CREATE TYPE rent_request_status AS ENUM (" +
		"'PENDING'," +
		"'ACCEPTED'," +
		"'REJECTED');")

	// RentStatus
	_ = Db.ProvideDb().Exec("DROP TYPE IF EXISTS rent_status;")
	_ = Db.ProvideDb().Exec("CREATE TYPE rent_status AS ENUM (" +
		"'ONGOING'," +
		"'COMPLETED');")

	// PaymentType
	_ = Db.ProvideDb().Exec("DROP TYPE IF EXISTS payment_types;")
	_ = Db.ProvideDb().Exec("CREATE TYPE payment_types AS ENUM (" +
		"'WEEKLY'," +
		"'MONTHLY'," +
		"'ANNUALLY');")

}

func DropTables(db datasource.Database) error {
	return db.DropTables(
		users.User{},
		users.AdminUser{},
		properties.Property{},
		properties.Address{},
		rent.RentRequest{},
		rent.Rent{},
		message.Conversation{},
		message.Message{},
	)
}
func CreateTables(db datasource.Database) error {
	return db.AutoMigrate(
		users.User{},
		users.AdminUser{},
		properties.Property{},
		properties.Address{},
		rent.RentRequest{},
		rent.Rent{},
		message.Conversation{},
		message.Message{},
	)
}
