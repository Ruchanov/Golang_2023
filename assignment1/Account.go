package accs

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func Registration(log, pass string) {
	Accounts := getData()
	temp := Account{
		Login:    log,
		Password: pass,
	}
	Accounts.Users = append(Accounts.Users, temp)

	saveDataToFile(&Accounts)
}

func Authorization(log, pass string) (*Account, bool) {
	Accounts := getData()
	for i := 0; i < len(Accounts.Users); i++ {
		//fmt.Println(Accounts.Users[i])
		if Accounts.Users[i].Login == log && Accounts.Users[i].Password == pass {
			return &Accounts.Users[i], true
		}
	}
	//return
	return &Account{}, false
}
