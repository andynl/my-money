package controller

import (
	"github.com/andynl/my-money/model"
	"github.com/gofiber/fiber/v2"
	"time"
)

type TransactionData struct {
	Summary      summary        `json:"summary"`
	Transactions []transaction `json:"transactions"`
}

type summary struct {
	BalanceLastMonth int `json:"balance_last_month"`
	BalanceThisMonth int `json:"balance_this_month"`
}

type transaction struct {
	ID			string `json:"id"`
	Type 		string `json:"type"`
	CategoryID 	string `json:"category_id"`
	Name 		string `json:"name"`
	Amount 		float64 `json:"amount"`
	CreatedAt 	time.Time `json:"created_at"`
}

var transactions = []transaction{
	{ID: "1", Type: "Income", CategoryID: "1",Name: "Gaji Bulan September 2021", Amount: 45000000, CreatedAt: time.Now()},
	{ID: "2", Type: "Income", CategoryID: "2", Name: "Jual Tanah Warisan", Amount: 100000000,  CreatedAt: time.Now()},
	{ID: "3", Type: "Income", CategoryID: "3", Name: "Arisan", Amount: 15000000, CreatedAt: time.Now()},
	{ID: "4", Type: "Outcome", CategoryID: "4", Name: "Iphone 12 Pro", Amount: 13000000, CreatedAt: time.Now()},
	{ID: "5", Type: "Outcome", CategoryID: "5", Name: "Baju Bagus", Amount: 1000000, CreatedAt: time.Now()},
	{ID: "6", Type: "Outcome", CategoryID: "6", Name: "Beli Ayam Geprek", Amount: 10000, CreatedAt: time.Now()},
}


var incomes = []transaction{
	{ID: "1", Type: "Income", CategoryID: "1",Name: "Gaji Bulan September 2021", Amount: 45000000, CreatedAt: time.Now()},
	{ID: "2", Type: "Income", CategoryID: "2", Name: "Jual Tanah Warisan", Amount: 100000000,  CreatedAt: time.Now()},
	{ID: "3", Type: "Income", CategoryID: "3", Name: "Arisan", Amount: 15000000, CreatedAt: time.Now()},
}

var outcomes = []transaction{
	{ID: "4", Type: "Outcome", CategoryID: "4", Name: "Iphone 12 Pro", Amount: 13000000, CreatedAt: time.Now()},
	{ID: "5", Type: "Outcome", CategoryID: "5", Name: "Baju Bagus", Amount: 1000000, CreatedAt: time.Now()},
	{ID: "6", Type: "Outcome", CategoryID: "6", Name: "Beli Ayam Geprek", Amount: 10000, CreatedAt: time.Now()},
}

var summaries = summary{BalanceLastMonth: 34000000, BalanceThisMonth: 35000000}

type TransactionController struct{}

func NewTransactionController() TransactionController {
	return TransactionController{}
}

func (controller *TransactionController) Route(app *fiber.App) {
	app.Get("/api/transactions", controller.Get)
	app.Get("/api/transactions/income", controller.Income)
	app.Get("/api/transactions/outcome", controller.Outcome)
	app.Get("/api/v2/transactions", controller.Get2)
}

func (controller *TransactionController) Get(c *fiber.Ctx) error {
	// params
	// limit
	// type: income, expense
	// sort by: highest, lowest

	//msg := fmt.Sprintf("limit %s - type %s", c.QueryParser("limit"), c.Query("type"))
	//return c.SendString(msg)
	//limit := c.Query("limit")

 	//trxData := TransactionData{summaries, transactions}
	return c.JSON(transactions)
}

func (controller *TransactionController) Get2(c *fiber.Ctx) error {
	trxData := TransactionData{summaries, transactions}
	return c.JSON(model.WebResponse{
		Code: 200,
		Status: "OK",
		Data: trxData,
	})
}

func (controller *TransactionController) Income(c *fiber.Ctx) error {
	return c.JSON(incomes)
}

func (controller *TransactionController) Outcome(c *fiber.Ctx) error {
	return c.JSON(outcomes)
}