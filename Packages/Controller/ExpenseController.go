package Controller

import (
	"ExpenceTraker/Helper"
	Model "ExpenceTraker/Packages/Models"
	Utility "ExpenceTraker/Packages/Utilities"
	"encoding/json"
	"net/http"
)

func AddExpenceControl(writer http.ResponseWriter, Req *http.Request) {
	IsValidCredentials, _ := IsloggedIN(Req)

	if IsValidCredentials != true {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	Response := Helper.GenericResponse{}

	Exp := &Helper.Expense{}
	Utility.ParseBody(Req, Exp)

	IsAdded, err := Model.AddExpense(*Exp)

	if err != nil {
		ErrorResponse(writer, Response, err)
		return
	}

	if IsAdded == true {
		validOperationResponse(writer, Response, "SignUp was successfull")
		return
	} else {
		InvalidOperationResponse(writer, Response, "SignUp was Unsuccessfull")

		return
	}

}

func DeleteExpenseControl(writer http.ResponseWriter, Req *http.Request) {
	IsValidCredentials, _ := IsloggedIN(Req)

	if IsValidCredentials != true {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	Response := Helper.GenericResponse{}
	Exp := &Helper.RemoveExpense{}
	Utility.ParseBody(Req, Exp)

	IsDeleted, err := Model.DeleteExpense(*Exp)

	if err != nil {
		ErrorResponse(writer, Response, err)
		return
	}

	if IsDeleted == true {
		validOperationResponse(writer, Response, "SignUp was successfull")
		return
	} else {
		InvalidOperationResponse(writer, Response, "SignUp was Unsuccessfull")

		return
	}

}

func GetExpenseControl(writer http.ResponseWriter, Req *http.Request) {
	IsValidCredentials, _ := IsloggedIN(Req)

	if IsValidCredentials != true {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	Response := Helper.CostResponse{}
	Exp := &Helper.GetExpenseCost{}
	Utility.ParseBody(Req, Exp)

	TotalCost, err := Model.GetExpense(*Exp)

	if err != nil {
		Response.Message = "Unable to Process"
		res, _ := json.Marshal(Response)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		writer.Write(res)
		return
	}

	Response.Message = "Succesfully"
	Response.Cost = TotalCost
	Response.EndDate = Exp.EndDate
	Response.StartDate = Exp.StartDate

	res, _ := json.Marshal(Response)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
	return

}

func UpdateExpenseControl(writer http.ResponseWriter, Req *http.Request) {
	IsValidCredentials, _ := IsloggedIN(Req)

	if IsValidCredentials != true {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	Response := Helper.GenericResponse{}
	Exp := &Helper.UpdateExpense{}
	Utility.ParseBody(Req, Exp)

	UpdatedExpense, err := Model.UpdateExpense(*Exp)

	if err != nil {
		Response.Message = "Unable to Process"
		res, _ := json.Marshal(Response)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		writer.Write(res)
		return
	}

	res, _ := json.Marshal(UpdatedExpense)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
	return

}
