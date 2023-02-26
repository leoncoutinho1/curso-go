package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/leoncoutinho1/curso_go/internal/order/infra/database"
	"github.com/leoncoutinho1/curso_go/internal/order/usecase"
	"github.com/leoncoutinho1/curso_go/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("mysql", os.Getenv("USER")+":"+os.Getenv("PASS")+"@/"+os.Getenv("DATABASE"))
	if err != nil {
		panic(err)
	}
	repository := database.NewOrderRepository(db)
	uc := usecase.CalculateFinalPriceUseCase{
		OrderRepository: repository,
	}

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out)
	for msg := range out {
		var inputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}
		outputDTO, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println(outputDTO)
		time.Sleep(500 * time.Millisecond)
	}
}
