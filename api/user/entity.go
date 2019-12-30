package user

// Table user table name
const Table string = "user"

// Col user table column name
var Col col = col{
	PK:          "id",
	ID:          "id",
	Email:       "email",
	Pwd:         "pwd",
	Name:        "name",
	Role:        "role",
	Active:      "active",
	CreatedDate: "created_date",
	UpdatedDate: "updated_date",
}

// Entity user table entity
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

type col struct {
	PK          string
	ID          string
	Email       string
	Pwd         string
	Name        string
	Role        string
	Active      string
	CreatedDate string
	UpdatedDate string
}
