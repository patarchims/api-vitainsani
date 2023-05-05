package user

type ApiUser struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Ket      string `json:"ket"`
}

func (ApiUser) TableName() string {
	return "rekam.api_user"
}
