package accs

type Account struct {
	login    string
	password string
	money    int
}

func Registration(login, password string) {
	db := initDB()
	db.Query("INSERT INTO accounts Values(%s, %s)", login, password)
	//b := Account{login, password, 10000}
	//Accounts = getDataAboutAccs()
	//Accounts = append(Accounts, b)

}

//func Authorization(log, pass string) *Account {
//	//for i := 0; i < len(Accounts); i++ {
//	//	if Accounts[i].login == login && Accounts[i].password == password {
//	//		return true, i
//	//	}
//	//}
//	//return false, -1
//	db := initDB()
//	db.Query("SELECT * from accounts WHERE login = log AND pass = password")
//}
