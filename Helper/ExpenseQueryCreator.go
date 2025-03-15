package Helper

import (
	"ExpenceTraker/Helper/Caption"
	HelperType "ExpenceTraker/Helper/Type"
)

func AddExpenseQueryGenerator(exp Expense) string {
	BaseQuery := "INSERT INTO Expense_Detail (Name, ExpenseDate, Is_Visible, Cost, Type ) VALUES  ('"

	BaseQuery = BaseQuery + exp.Name + "','" + exp.Date + "',true,'" + exp.Cost + "','"

	switch exp.Category {
	case Caption.Electri_Bill:
		BaseQuery = BaseQuery + HelperType.Electri_Bill + "')"

	case Caption.Groceries:
		BaseQuery = BaseQuery + HelperType.Groceries + "')"

	case Caption.Petrol_Cost:
		BaseQuery = BaseQuery + HelperType.Petrol_Cost + "')"

	case Caption.Phone_Bill:
		BaseQuery = BaseQuery + HelperType.Phone_Bill + "')"

	case Caption.Savings:
		BaseQuery = BaseQuery + HelperType.Savings + "')"

	case Caption.Maintainance:
		BaseQuery = BaseQuery + HelperType.Maintainance + "')"

	case Caption.Outing_Cost:
		BaseQuery = BaseQuery + HelperType.Outing_Cost + "')"

	case Caption.House_Help:
		BaseQuery = BaseQuery + HelperType.House_Help + "')"

	case Caption.Tax:
		BaseQuery = BaseQuery + HelperType.Tax + "')"

	case Caption.Mislaneous:
		BaseQuery = BaseQuery + HelperType.Mislaneous + "')"
	}

	return BaseQuery 

}

func GetExpenseQueryGenerator(exp GetExpenseCost) string {
	BaseQuery := "select SUM(Cost) from Expense_Detail where Is_Visible=1 AND ExpenseDate <= '"

	BaseQuery = BaseQuery + exp.EndDate + "' AND ExpenseDate >= '" + exp.StartDate + "'"

	return BaseQuery
}

func UpdateExpenseQueryGenerator(exp UpdateExpense) string {
	BaseQuery := "Update Expense_Detail set Name='" + exp.Name + "', Cost=" + exp.Cost + ", ExpenseDate='" + exp.Date + "', Is_Visible=" + exp.IsVisible + " where ID=" + exp.ID

	return BaseQuery
}

func DeleteQueryGenerator(exp RemoveExpense) string {
	BaseQuery := "Update Expense_Detail set Is_Visible = false where ID=" + exp.ID

	return BaseQuery
}
