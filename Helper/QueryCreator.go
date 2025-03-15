package Helper

func IsUserExistsQuery(SignUpUser User, CurrCredentials Credentials, IsSignUp bool) (string, error) {
	query := "SELECT COUNT(*) AS USER FROM SEC_LOGIN where UserName = '"
	if IsSignUp == true {
		query = query + SignUpUser.UserName + "'"
	} else {
		query = query + CurrCredentials.UserName + "'"
	}

	return query, nil
}

func SignUpQueryCreator(SignUpUser User) string {
	query := "INSERT INTO SEC_LOGIN (UserName, FirstName, LastName, Password) VALUES  ('" + SignUpUser.UserName + "','" + SignUpUser.FirstName + "','" + SignUpUser.LastName + "','" + SignUpUser.Password + "')"
	return query
}

func CredentialsQueryCreator(CurrCredentials Credentials) string {
	query := "SELECT COUNT(*) AS USER FROM SEC_LOGIN where UserName = '" + CurrCredentials.UserName + "' and Password='" + CurrCredentials.Password + "'"
	return query
}
