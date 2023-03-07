package request

// O validator é utilizao pelo gin
// na hora dded fazer o Bind ele  faz a validação se eu utilizar o campo binding
// casos eu queira utilizar o calidator puro eu uso a tag validate
type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%&*"`
	Name     string `json:"name" binding:"required,min=4,max=100"`
	Age      int8   `json:"age" binding:"required,min=2,max=140"`
}
