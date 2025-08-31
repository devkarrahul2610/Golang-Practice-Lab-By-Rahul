package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the composition and interfaces demo.")

	// OrderService logs to File
	orderSVC := OrderService{

		FileLogger{
			BaseLogger: BaseLogger{ServiceName: "OrderService"},
			FilePath:   "/var/log/orders.log",
		},
	}

	orderSVC.ProcessOrder("ORD123")

	// PaymentService logs to Cloud
	paymentSvc := PaymentService{
		CloudLogger{
			BaseLogger:   BaseLogger{ServiceName: "PaymentService"},
			CloudService: "AWS CloudWatch",
		},
	}
	paymentSvc.ProcessPayment("PAY456")
}
