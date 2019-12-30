package user

const (
	// Table user table name
	Table string = "user"

	// PK user table pk column name
	PK string = "id"
	// ID user table id column name
	ID string = "id"
	// Email user table email column name
	Email string = "email"
	// Pwd user table pwd column name
	Pwd string = "pwd"
	// Name user table name column name
	Name string = "name"
	// Role user table role column name
	Role string = "role"
	// Active user table active column name
	Active string = "active"
	// CreatedDate user table created date column name
	CreatedDate string = "created_date"
	// UpdatedDate user table updated date column name
	UpdatedDate string = "updated_date"
)

// Entity table struct
type Entity struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Pwd         string `json:"pwd"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Active      bool   `json:"active"`
	CreatedDate string `json:"creatDate"`
	UpdatedDate string `json:"updateDate"`
}
