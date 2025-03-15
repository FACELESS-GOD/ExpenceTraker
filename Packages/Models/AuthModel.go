package Model

import (
	"ExpenceTraker/Helper"
	Utility "ExpenceTraker/Packages/Utilities"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
)

var IsAnyError bool
var mutex sync.Mutex

func DoSignUp(User Helper.User) (error, bool) {
	IsAnyError = false

	IsValid, err := IsUserExists(User, Helper.Credentials{}, true)

	if err != nil {
		return err, false
	}

	if IsValid != true {

		var waitGroup sync.WaitGroup

		waitGroup.Add(1)
		go AddUser(&waitGroup, User)

		waitGroup.Add(1)
		go AddCredentials(&waitGroup, User)

		waitGroup.Wait()

		if IsAnyError == true {
			return nil, false
		} else {
			return nil, true
		}

	} else {
		return nil, false
	}
}

func AddUser(WaitGroup *sync.WaitGroup, User Helper.User) {
	defer WaitGroup.Done()
	query := Helper.AddUserQueryCreator(User)

	IsSignUpUser, err := Utility.DatabaseInstace.Query(query)

	if err != nil {
		mutex.Lock()
		IsAnyError = true
		mutex.Unlock()
		return
	}

	for IsSignUpUser.Next() {
	}

}

func AddCredentials(WaitGroup *sync.WaitGroup, User Helper.User) {
	defer WaitGroup.Done()

	query := Helper.SignUpQueryCreator(User)

	IsSignUpUser, err := Utility.DatabaseInstace.Query(query)

	if err != nil {
		mutex.Lock()
		IsAnyError = true
		mutex.Unlock()
		return
	}
	for IsSignUpUser.Next() {
	}

	errRedis := Utility.RedisInstance.Set(Utility.Ctx, User.UserName, User.Password, 0).Err()

	if errRedis != nil {
		mutex.Lock()
		IsAnyError = true
		mutex.Unlock()
		return
	}

}

func IsValidCredentials(Credentials Helper.Credentials) (bool, error) {
	var CurrCount Helper.GenericCountResponse

	IsValid, err := IsUserExists(Helper.User{}, Credentials, false)

	if err != nil {
		return false, err
	}

	if IsValid != true {
		query := Helper.CredentialsQueryCreator(Credentials)

		IsLoginUser, err := Utility.DatabaseInstace.Query(query)

		if err != nil {
			return false, err
		}

		for IsLoginUser.Next() {
			err := IsLoginUser.Scan(&CurrCount.Count)
			if err != nil {
				return false, err
			}
		}

		if CurrCount.Count > 0 {
			return true, nil
		} else {
			return false, nil
		}

	} else {
		return false, err
	}
}

func IsUserExists(User Helper.User, Credentials Helper.Credentials, IsSignUp bool) (bool, error) {
	var CountResponse Helper.GenericCountResponse

	isUserExistsQuery, err := Helper.IsUserExistsQuery(User, Credentials, true)

	if err != nil {
		return false, err
	}

	IsValid, err := Utility.DatabaseInstace.Query(isUserExistsQuery)

	if err != nil {
		return false, err
	}

	for IsValid.Next() {
		err := IsValid.Scan(&CountResponse.Count)
		if err != nil {
			return false, err
		}
	}

	if CountResponse.Count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func IsValid(Req *http.Request) (bool, error) {
	tokenString := Req.Header.Get("Authorization")
	tokenString = tokenString[len("Bearer "):]
	isValid, err := VerifyToken(tokenString)

	if err != nil {
		return false, err
	}
	return isValid, nil
}

func CreateToken(UserName string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"UserName": UserName,
			"exp":      time.Now().Add(time.Hour * 1).Unix(),
		})
	tokenString, err := token.SignedString(Helper.SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(TokenString string) (bool, error) {

	token, err := jwt.Parse(
		TokenString,
		func(token *jwt.Token) (interface{}, error) {
			return Helper.SecretKey, nil
		})

	if err != nil {
		log.Printf("Invalid Token")
		return false, err
	}

	if token.Valid != true {
		return false, err
	}
	return true, nil
}

func IsPresentInRedis(Credentials Helper.Credentials) (bool, error) {
	PasswordSavedInRedis, err := Utility.RedisInstance.Get(Utility.Ctx, Credentials.UserName).Result()

	switch err {
	case nil:
		if PasswordSavedInRedis == Credentials.Password {
			return true, nil
		} else {
			return false, nil
		}

	case redis.Nil:
		return false, nil

	default:
		return false, err

	}

}
