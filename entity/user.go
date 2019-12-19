package entity

// User table struct
type User struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Pwd         string `json:"pwd"`
	Name        string `json:"name"`
	Identity    string `json:"identity"`
	Active      bool   `json:"active"`
	CreatedDate string `json:"creatDate"`
	UpdatedDate string `json:"updateDate"`
}

// Table get table name
func (user *User) Table() string {
	return "user"
}

// PK get table pk column name
func (user *User) PK() string {
	return "id"
}

// IDCol get ID column name
func (user *User) IDCol() string {
	return "id"
}

// EmailCol get Email column name
func (user *User) EmailCol() string {
	return "email"
}

// PwdCol get Pwd column name
func (user *User) PwdCol() string {
	return "pwd"
}

// NameCol get Name column name
func (user *User) NameCol() string {
	return "name"
}

// IdentityCol get Identity column name
func (user *User) IdentityCol() string {
	return "identity"
}

// ActiveCol get Active column name
func (user *User) ActiveCol() string {
	return "active"
}

// CreatedDateCol get CreatedDate column name
func (user *User) CreatedDateCol() string {
	return "created_date"
}

// UpdatedDateCol get UpdatedDate column name
func (user *User) UpdatedDateCol() string {
	return "updated_date"
}
