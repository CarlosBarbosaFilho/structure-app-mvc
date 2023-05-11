package response

type UserResponseLogin struct {
	ID       uint
	Name     string
	Email    string `gorm:"unique"`
	Password string
}
