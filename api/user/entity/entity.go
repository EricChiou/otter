package entity

// User table struct
type User struct {
	Col         UserCol
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Pwd         string `json:"pwd"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Active      bool   `json:"active"`
	CreatedDate string `json:"creatDate"`
	UpdatedDate string `json:"updateDate"`
}

// Table get table name
func (user *User) Table() string {
	return "user"
}

// UserCol get user table column name
type UserCol struct{}

// PK get table pk column name
func (user *UserCol) PK() string {
	return "id"
}

// ID get ID column name
func (user *UserCol) ID() string {
	return "id"
}

// Email get Email column name
func (user *UserCol) Email() string {
	return "email"
}

// Pwd get Pwd column name
func (user *UserCol) Pwd() string {
	return "pwd"
}

// Name get Name column name
func (user *UserCol) Name() string {
	return "name"
}

// Role get role column name
func (user *UserCol) Role() string {
	return "role"
}

// Active get Active column name
func (user *UserCol) Active() string {
	return "active"
}

// CreatedDate get CreatedDate column name
func (user *UserCol) CreatedDate() string {
	return "created_date"
}

// UpdatedDate get UpdatedDate column name
func (user *UserCol) UpdatedDate() string {
	return "updated_date"
}
