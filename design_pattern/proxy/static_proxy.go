package proxy

type IUser interface {
	Login(username, password string) error
}

type User struct {
}

func (u *User) Login(username, password string) error {
	return nil
}

// userProxy
type UserProxy struct {
	user *User
}

func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{user: user}
}

// 实现相同的接口
func (p *UserProxy) Login(username, password string) error {
	if err := p.user.Login(username, password); err != nil {
		return err
	}
	return nil
}
