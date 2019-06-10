package knapsack

// The first element "0" is just the dummy element.
var value = []int{
	0,
	60,
	100,
	120,
}

// The first element "0" is just the dummy element.
var weight = []int{
	0,
	10,
	20,
	30,
}

func Max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}

func Knapsack(index int, capacity int) int {
	// If index or capacity
	if index == 0 || capacity == 0 {
		return 0
	}

	result := 0

	if weight[index] > capacity {
		result = Knapsack(index-1, capacity)
	} else {
		tmp1 := Knapsack(index-1, capacity)
		tmp2 := value[index] + Knapsack(index-1, capacity-weight[index])

		result = Max(tmp1, tmp2)
	}

	return result
}
