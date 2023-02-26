package usecase

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/leoncoutinho1/curso_go/internal/order/entity"
	"github.com/leoncoutinho1/curso_go/internal/order/infra/database"
	"github.com/stretchr/testify/suite"
)

type CalculateFinalPriceUseCaseTestSuite struct {
	suite.Suite
	OrderRepository entity.OrderRepositoryInterface
	Db              *sql.DB
}

func (suite *CalculateFinalPriceUseCaseTestSuite) SetupSuite() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("mysql", os.Getenv("USER")+":"+os.Getenv("PASS")+"@/"+os.Getenv("DATABASE"))
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
	suite.OrderRepository = database.NewOrderRepository(db)
}

func (suite *CalculateFinalPriceUseCaseTestSuite) TearDownTest() {
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
	suite.Run(t, new(CalculateFinalPriceUseCaseTestSuite))
}

func (suite *CalculateFinalPriceUseCaseTestSuite) TestCalculateFinalPrice() {
	order, err := entity.NewOrder("1", 10, 2)
	suite.NoError(err)
	order.CalculateFinalPrice()
	calculateFinalPriceInput := OrderInputDTO{
		Id:    order.Id,
		Price: order.Price,
		Tax:   order.Tax,
	}
	calculateFinalPriceUseCase := NewCalculateFinalPriceUseCase(suite.OrderRepository)
	output, err := calculateFinalPriceUseCase.Execute(calculateFinalPriceInput)
	suite.NoError(err)
	suite.Equal(order.Id, output.Id)
	suite.Equal(order.Price, output.Price)
	suite.Equal(order.Tax, output.Tax)
	suite.Equal(order.FinalPrice, output.FinalPrice)
}
