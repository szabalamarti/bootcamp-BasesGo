package taxes

func GetTax(salary float64) (tax float64) {
	if salary > 150000 {
		tax = salary * 0.27
	} else if salary > 50000 {
		tax = salary * 0.17
	} else {
		tax = 0.0
	}
	return tax
}
