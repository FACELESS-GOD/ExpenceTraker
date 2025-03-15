package Controller

import (
	"ExpenceTraker/Helper"
	Model "ExpenceTraker/Packages/Models"
	Utility "ExpenceTraker/Packages/Utilities"
	"encoding/json"
	"net/http"
)

func SignUp(writer http.ResponseWriter, Req *http.Request) {
	Response := Helper.GenericResponse{}

	CurrUser := &Helper.User{}

	Utility.ParseBody(Req, CurrUser)

	err, IsSaved := Model.DoSignUp(*CurrUser)

	if err != nil {
		ErrorResponse(writer, Response, err)
		return
	}

	if IsSaved == false {
		InvalidOperationResponse(writer, Response, "SignUp was Unsuccessfull")
		return
	} else {
		validOperationResponse(writer, Response, "SignUp was successfull")
		return
	}

}

func Login(writer http.ResponseWriter, Req *http.Request) {
	Response := Helper.GenericResponse{}

	Cred := &Helper.Credentials{}

	Utility.ParseBody(Req, Cred)

	IsPresentInRedis, err := IsRedis(*Cred)

	if err != nil {
		ErrorResponse(writer, Response, err)
		return
	}

	if IsPresentInRedis == true {
		Token, err := Model.CreateToken(Cred.UserName)

		if err != nil {
			ErrorResponse(writer, Response, err)
			return
		}

		writer.Header().Set("Authorization", "Bearer "+Token)
		validOperationResponse(writer, Response, "SignUp was successfull")
		return
	}

	Isvalid, err := Model.IsValidCredentials(*Cred)

	if err != nil {
		ErrorResponse(writer, Response, err)
		return
	}

	if Isvalid == false {
		InvalidOperationResponse(writer, Response, "SignUp was Unsuccessfull")
		return
	} else {
		Token, err := Model.CreateToken(Cred.UserName)

		if err != nil {
			ErrorResponse(writer, Response, err)
			return
		}

		writer.Header().Set("Authorization", "Bearer "+Token)
		validOperationResponse(writer, Response, "SignUp was successfull")
		return
	}

}

func IsRedis(CurrCred Helper.Credentials) (bool, error) {
	IsValid, err := Model.IsPresentInRedis(CurrCred)
	if err != nil {
		return false, err
	}

	if IsValid == true {
		return true, nil
	} else {
		return false, nil
	}

}

func validOperationResponse(writer http.ResponseWriter, Response Helper.GenericResponse, message string) {
	Response.Message = message
	res, _ := json.Marshal(Response)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func InvalidOperationResponse(writer http.ResponseWriter, Response Helper.GenericResponse, message string) {
	Response.Message = message
	res, _ := json.Marshal(Response)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	writer.Write(res)
}

func ErrorResponse(writer http.ResponseWriter, Response Helper.GenericResponse, err error) {
	Response.Message = "Unable to Process"
	res, _ := json.Marshal(Response)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	writer.Write(res)

}
