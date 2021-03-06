package userpo

// table name
const Table string = "user"

// pk name
const PK string = "id"

// column name
const (
	ID          string = "id"
	Acc         string = "acc"
	Pwd         string = "pwd"
	Name        string = "name"
	RoleCode    string = "role_code"
	Status      string = "status"
	CreatedDate string = "created_date"
	UpdatedDate string = "updated_date"
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
