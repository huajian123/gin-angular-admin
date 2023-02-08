package sysDto

type UserDto struct {
	Name      string `json:"name" binding:"required"`
	Telephone string `json:"telephone" binding:"required"`
	Email     string `json:"email" binding:"email"`
	Password  string `json:"password"`
}

//func UserVoToDto(v vo.UserVO) UserDto {
//	return UserDto{
//		Name:      v.Name,
//		Telephone: v.Telephone,
//		Email:     v.Email,
//		Password:  v.Password,
//	}
//}
