package accs

type Account struct {
	login    string `json:"login"`
	password string `json:"password"`
}

func Registration(log, pass string) {
	Accounts := getData()
	temp := Account{log, pass}
	Accounts.Users = append(Accounts.Users, temp)

	saveDataToFile(Accounts)
}

func Authorization(log, pass string) *Account {
	Accounts := getData()
	for i := 0; i < len(Accounts.Users); i++ {
		if Accounts.Users[i].login == log && Accounts.Users[i].password == pass {
			return &Accounts.Users[i]
		}
	}
	//return
	return &Account{}
}
