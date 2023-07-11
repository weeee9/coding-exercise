package score

func sortTheStudents(score [][]int, k int) [][]int {
	return quickSortTheStudents(score, k)
}

func quickSortTheStudents(score [][]int, k int) [][]int {
	if len(score) < 2 {
		return score
	} else {
		pivot := score[0]
		less := [][]int{}
		greater := [][]int{}

		for _, student := range score[1:] {
			if student[k] <= pivot[k] {
				less = append(less, student)
			} else {
				greater = append(greater, student)
			}
		}

		result := [][]int{}
		result = append(result, sortTheStudents(greater, k)...)
		result = append(result, pivot)
		result = append(result, sortTheStudents(less, k)...)

		return result
	}
}

func insertionSortTheStudents(score [][]int, k int) [][]int {
	if len(score) == 1 {
		return score
	}

	for i := 1; i < len(score); i++ {
		student := score[i]
		j := i - 1

		for j >= 0 && score[j][k] < student[k] {
			score[j+1] = score[j]
			j--
		}

		score[j+1] = student
	}

	return score
}
