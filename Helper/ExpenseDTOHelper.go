package Helper

type Expense struct {
	Name     string `json : "name"`
	Cost     string `json : "cost"`
	Date     string `json : "date"`
	Category string `json : "category"`
}

type GetExpenseCost struct {
	StartDate string `json : "startdate"`
	EndDate   string `json : "enddate"`
}

type UpdateExpense struct {
	ID        string `json : "ID"`
	Name      string `json : "name"`
	Cost      string `json : "cost"`
	Date      string `json : "date"`
	Category  string `json : "category"`
	IsVisible string `json : "isvisible"`
}

type RemoveExpense struct {
	ID string `json : "ID"`
}

type DBExpense struct {
	ID        int     `json : "ID"`
	Name      string  `json : "name"`
	Cost      float32 `json : "cost"`
	Date      string  `json : "date"`
	Category  int     `json : "category"`
	IsVisible bool    `json : "isvisible"`
}

type CostResponse struct {
	StartDate string
	EndDate   string
	Cost      int64
	Message   string
}
