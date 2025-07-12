package interfaces

import "fmt"


type paymenter interface {
	Pay(amount float64) string
}

type PaymentGateway struct {
	gateway paymenter
}

type BkashPayment struct {
}

func (b *BkashPayment) Pay(amount float64) string {
	addedCharge := (18.5/1000) * amount
	amount += addedCharge
	return fmt.Sprintf("Paid %.2f using Bkash", amount)
}

type NagadPayment struct {
}
func (n *NagadPayment) Pay(amount float64) string {
	addedCharge := (15.0 /1000) * amount
	amount += addedCharge
	return fmt.Sprintf("Paid %.2f using Nagad", amount)
}

func (pg *PaymentGateway) ProcessPayment(amount float64) string {
	return pg.gateway.Pay(amount)
}

func Interfaces() { 
	// payments := []paymenter{
	// 	&BkashPayment{},
	// 	&NagadPayment{},
	// }

	// for _, payment := range payments {
	// 	fmt.Println(payment.Pay(100.50))
	// }

	paymentAmount := 20000.0;
	pg := PaymentGateway{gateway: &BkashPayment{}}
	fmt.Println(pg.ProcessPayment(paymentAmount))

	pg.gateway = &NagadPayment{}
	fmt.Println(pg.ProcessPayment(paymentAmount))

	fmt.Println("All payments processed successfully!")
}