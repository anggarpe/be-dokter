package models

type User struct {
	ID       string  `gorm:"primary_key;unique;not null" json:"id"`
	Name     string `gorm:"not null" json:"name" binding:"required"`
	Email    string `gorm:"type:varchar(100);not null;unique_index" json:"email"`
	Password string `json:"password"`
	Address  string `gorm:"type:text" json:"address"`
	Phone    string `json:"phone"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type user []User

//func (usr *User) TableName() string {
//	return "user"
//}
//
//func (usr User) ToString() string {
//	return fmt.Sprintf("id: %d\nname: %s\nemail: %s", usr.ID, usr.Name, usr.Email)
//}
