package sysEntity

type UserEntity struct {
	ID        uint   `json:"id" gorm:"primary_key;not null;unique"`
	Name      string `json:"name" gorm:"varchar(20);not null"`
	Telephone string `json:"telephone" gorm:"varchar(11);not null"`
	Email     string `json:"email" gorm:"varchar(255)"`
	Password  string `json:"password" gorm:"varchar(255)"`
}
