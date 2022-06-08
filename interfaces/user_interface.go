package interfaces

type User struct {
	ID   string
	Name string
}

type UserMethods interface {
	GetUserById(id string)
	InsertUser(name string)
}

func (u *User) GetUserById(id string)  {}
func (u *User) InsertUser(name string) {}

type UserMock struct{}

func (u *UserMock) GetUserById(id string)  {}
func (u *UserMock) InsertUser(name string) {}
