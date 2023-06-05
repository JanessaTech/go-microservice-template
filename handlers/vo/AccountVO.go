package vo

type AccountVO struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
