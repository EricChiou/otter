package entity

// User table struct
type User struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Pwd        string `json:"pwd"`
	Name       string `json:"name"`
	Identity   string `json:"identity"`
	Active     bool   `json:"active"`
	CreatDate  string `json:"creatDate"`
	UpdateDate string `json:"updateDate"`
}