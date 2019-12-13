package entity

// User table struct
type User struct {
	ID         int    `json:"id" db:"id"`
	Email      string `json:"email" db:"email"`
	Pwd        string `json:"pwd" db:"pwd"`
	Name       string `json:"name" db:"name"`
	Identity   string `json:"identity" db:"identity"`
	Active     bool   `json:"active" db:"active"`
	CreatDate  string `json:"creatDate" db:"creat_date"`
	UpdateDate string `json:"updateDate" db:"update_date"`
}
