package accs

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func Registration(log, pass string) {
	Accounts := GetData()
	temp := Account{
		Login:    log,
		Password: pass,
	}
	Accounts.Users = append(Accounts.Users, temp)

	saveDataToFile(&Accounts)
	//fmt.Fprintf(w, "Registration successful for user: %s", log)
}

func Authorization(log, pass string) (*Account, bool) {
	Accounts := GetData()
	for i := 0; i < len(Accounts.Users); i++ {
		//fmt.Println(Accounts.Users[i])
		if Accounts.Users[i].Login == log && Accounts.Users[i].Password == pass {
			return &Accounts.Users[i], true
		}
	}
	//return
	return &Account{}, false
}
