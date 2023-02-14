package accs

type Account struct {
	login    string `json:"login"`
	password string `json:"password"`
}

func Registration(login, password string) {
	Accounts := getDataAboutAccs()
	Accounts = append(Accounts, Account{login, password})
	saveAccs(Accounts)
}

func Authorization(log, pass string) *Account {
	Accounts := getDataAboutAccs()
	for i := 0; i < len(Accounts); i++ {
		if Accounts[i].login == log && Accounts[i].password == pass {
			return &Accounts[i]
		}
	}
	//return
	return &Account{}
}
