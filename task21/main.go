package main

//Реализовать паттерн «адаптер» на любом примере.

type PayPal struct{}

func (p *PayPal) MakePayment(amount float32) bool {
	return true
}

type Amazon struct{}

func (a *Amazon) PayAmazon(amount float32) bool {
	return true
}

type PaymentGateway interface {
	ProcessPayment(amount float32) bool
}

type PayPalAdapter struct {
	PayPal *PayPal
}

func (p *PayPalAdapter) ProcessPayment(amount float32) bool {
	return p.PayPal.MakePayment(amount)
}

type AmazonAdapter struct {
	Amazon *Amazon
}

func (a *AmazonAdapter) ProcessPayment(amount float32) bool {
	return a.Amazon.PayAmazon(amount)
}

func main() {
	paymentGateway := &PayPalAdapter{
		PayPal: &PayPal{},
	}
	paymentGateway2 := &AmazonAdapter{
		Amazon: &Amazon{},
	}
	paymentGateway.ProcessPayment(100)
	paymentGateway2.ProcessPayment(100)
}
