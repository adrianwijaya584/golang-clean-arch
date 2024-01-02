package userdto

type CreateUserDto struct {
	Name string `validate:"required,min=5"`
}
