package sysDto

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
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
