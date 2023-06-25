package josephus

func Josephus(items []interface{}, k int) []interface{} {
	queue := append([]interface{}{}, items...)
	result := make([]interface{}, 0, len(items))

	index := 0
	for len(queue) > 0 {
		index = (index + k - 1) % len(queue)
		result = append(result, queue[index])
		queue = append(queue[:index], queue[index+1:]...)
	}

	return result
}
