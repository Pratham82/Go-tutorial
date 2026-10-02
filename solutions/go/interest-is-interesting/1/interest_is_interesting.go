package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var interestRate float32

	switch {
	case balance < 0:
		interestRate = 3.213

	case balance >= 0 && balance < 1000:
		interestRate = 0.5

	case balance >= 1000 && balance < 5000:
		interestRate = 1.621

	case balance >= 5000:
		interestRate = 2.475
	}
	return interestRate

}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	yearsTook := 0
	// predictedBalance := balance
	// 	for {
	// 		yearsTook++
	// 		predictedBalance = AnnualBalanceUpdate(predictedBalance)
	//
	// 		if predictedBalance >= targetBalance {
	// 			break
	// 		}
	// 	}

	// for predictedBalance < targetBalance {
	// 	yearsTook++
	// 	predictedBalance = AnnualBalanceUpdate(predictedBalance)
	// }

	for balance < targetBalance {
		yearsTook++
		balance = AnnualBalanceUpdate(balance)
	}

	return yearsTook
}
