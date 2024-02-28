package requests

type Login struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
