package service

type UserService struct {
	
}

func (u *UserService) isValidUser(account string, password string) error{
	params := map[string]string{"account": account}
	RemoteService("UserService.GetUserByAccount", params, )
}