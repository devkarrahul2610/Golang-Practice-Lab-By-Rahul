package main

// OrderService uses FileLogger
type OrderService struct {
	FileLogger
}

func (o OrderService) ProcessOrder(orderId string) {
	o.Log("Processing Order" + orderId) // BaseLogger.Log()
	o.LogToFile("order" + orderId + "stored in file")
}

// PaymentService uses CloudLogger
type PaymentService struct {
	CloudLogger
}

func (p PaymentService) ProcessPayment(paymentId string) {
	p.Log("processing payment" + paymentId)
	p.LogToCloud("payment" + paymentId + "send to cloud")
}
