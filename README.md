# ExpenceTraker
This is an Expense Traker API Build using go 

This API , apart from the authentication routes  will consist of 3 routes . 

These routes are as follows:
  1) "/GetExpense"
  2) "/UpdateExpence"
  3) "/DeleteExpence"
  4) "/AddExpense"

The request body to the above api must of the following types , respectively

/GetExpense --


  {
    "StartDate":""
    "EndDate":""
  }

/UpdateExpence --


  {
    "ExpenseID":""    
  }

/DeleteExpence --


  {
    "ExpenseID":""
  }

  Also the Sensitive information such as connection strings are directly hard coded instead of using envoirment variables / Key vaults 
  

/AddExpense --
  {
  "ExpenseName" : ""
  "ExpenseType" : ""
    "Date":""    
  }
