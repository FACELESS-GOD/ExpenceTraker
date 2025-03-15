package Model

import (
	"ExpenceTraker/Helper"
	Utility "ExpenceTraker/Packages/Utilities"
	"strconv"
)

/*
 Expense_Detail(

	Name varchar(255),
    Cost float,
	ExpenseDate date,
    Is_Visible boolean,
    Type int,

);
*/

func AddExpense(Exp Helper.Expense) (bool, error) {

	query := Helper.AddExpenseQueryGenerator(Exp)

	IsExpenseAdded, err := Utility.DatabaseInstace.Query(query)

	if err != nil {
		return false, err
	}
	var GenericCountResponse Helper.GenericIDResponse
	for IsExpenseAdded.Next() {
		err := IsExpenseAdded.Scan(&GenericCountResponse.ID)
		if err != nil {
			return false, nil
		}
	}
	

	return true, nil

}

func GetExpense(Exp Helper.GetExpenseCost) (int64, error) {
	query := Helper.GetExpenseQueryGenerator(Exp)

	execute, err := Utility.DatabaseInstace.Query(query)

	if err != nil {
		return 0, err
	}
	var GenericCountResponse Helper.GenericCountResponseFloat
	for execute.Next() {
		err := execute.Scan(&GenericCountResponse.Count)
		if err != nil {
			return 0, nil
		}
	}

	return int64(GenericCountResponse.Count), nil

}

func UpdateExpense(Exp Helper.UpdateExpense) (Helper.Expense, error) {
	query := Helper.UpdateExpenseQueryGenerator(Exp)

	var NewExpense Helper.Expense
	var DbExpense Helper.DBExpense

	execute, err := Utility.DatabaseInstace.Query(query)

	if err != nil {
		return Helper.Expense{}, err
	}

	for execute.Next() {
	}

	query = "Select * from Expense_Detail where Id = " + Exp.ID

	execute, error := Utility.DatabaseInstace.Query(query)

	if error != nil {
		return Helper.Expense{}, err
	}

	for execute.Next() {
		error := execute.Scan(&DbExpense.ID, &DbExpense.Name, &DbExpense.Cost, &DbExpense.Date, &DbExpense.IsVisible, &DbExpense.Category)
		if error != nil {
			return Helper.Expense{}, err
		}
	}

	NewExpense.Name = DbExpense.Name
	NewExpense.Date = DbExpense.Date

	NewExpense.Cost = strconv.FormatFloat(float64(DbExpense.Cost), 'f', -1, 32)
	NewExpense.Category = strconv.Itoa(DbExpense.Category)

	return NewExpense, nil

}

func DeleteExpense(Exp Helper.RemoveExpense) (bool, error) {
	query := Helper.DeleteQueryGenerator(Exp)

	execute, err := Utility.DatabaseInstace.Query(query)

	if err != nil {
		return false, err
	}

	for execute.Next() {
	}

	return true, nil
}
