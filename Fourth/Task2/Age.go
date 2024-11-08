package Task2

func AverageAge(people map[string]int) float64 {
	if len(people) == 0 {
		return 0
	}

	totalAge := 0
	for _, age := range people {
		totalAge += age
	}

	return float64(totalAge) / float64(len(people))
}
