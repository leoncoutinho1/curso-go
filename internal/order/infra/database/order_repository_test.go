package database

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/leoncoutinho1/curso_go/internal/order/entity"
	"github.com/stretchr/testify/suite"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("mysql", os.Getenv("USER")+":"+os.Getenv("PASS")+"@/"+os.Getenv("DATABASE"))
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("mysql", os.Getenv("USER")+":"+os.Getenv("PASS")+"@/"+os.Getenv("DATABASE"))
	suite.NoError(err)
	db.Exec("DROP TABLE orders")
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)
	var orderResult entity.Order
	err = suite.Db.QueryRow("select id, price, tax, final_price from orders where id = ?", order.Id).Scan(&orderResult.Id, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)
	suite.NoError(err)
	suite.Equal(order.Id, orderResult.Id)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}
