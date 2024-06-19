package main

func accountBalanceAfterPurchase(purchaseAmount int) int {
	r := (100-purchaseAmount) % 10
	if r <= 5 && r > 0{
		return 100 - (purchaseAmount/10+1)*10 
	} else {
		return 100 - (purchaseAmount/10)*10
	}
}

func main() {

}