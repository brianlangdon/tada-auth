package users

type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}

type UnableToCreateUserError struct{}

func (m *UnableToCreateUserError) Error() string {
	return "unable to create user"
}
