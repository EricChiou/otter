package userpo

// table name
const Table string = "user"

// pk name
const PK string = "user.id"

// column name
const (
	ID          string = "user.id"
	Acc         string = "user.acc"
	Pwd         string = "user.pwd"
	Name        string = "user.name"
	RoleCode    string = "user.role_code"
	Status      string = "user.status"
	CreatedDate string = "user.created_date"
	UpdatedDate string = "user.updated_date"
)

// Entity user table entity
type Entity struct {
	ID          int    `json:"id,omitempty"`
	Acc         string `json:"acc,omitempty"`
	Pwd         string `json:"pwd,omitempty"`
	Name        string `json:"name,omitempty"`
	RoleCode    string `json:"roleCode,omitempty"`
	Status      string `json:"status,omitempty"`
	CreatedDate string `json:"creatDate,omitempty"`
	UpdatedDate string `json:"updateDate,omitempty"`
}
