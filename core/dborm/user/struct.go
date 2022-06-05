package user

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SecretInput struct {
	UserID    uint   `json:"userId"`
	SecretId  string `json:"secretId" binding:"required"`
	SecretKey string `json:"secretKey" binding:"required"`
	Describe  string `json:"describe" binding:"required"`
}
