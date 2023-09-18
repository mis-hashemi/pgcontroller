package entity

const (
	FieldFirstName   = "first_name"
	FieldLastName    = "last_name"
	FieldPhoneNumber = "phone_number"
)

type User struct {
	ID          string
	FirstName   string
	LastName    string
	Password    string
	PhoneNumber string
}
