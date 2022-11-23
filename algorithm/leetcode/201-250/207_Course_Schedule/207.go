package _07_Course_Schedule

//TODO
func canFinish(numCourses int, prerequisites [][]int) bool {
	courseNum := make([]int, numCourses, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		courseNum[prerequisites[i][0]]++
	}
	var queue []int
	for i := 0; i < len(courseNum); i++ {
		if courseNum[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		current := queue[0]
		queue = append(queue[1:])
		for _, pair := range prerequisites {
			if courseNum[pair[0]] == 0 {
				continue
			}
			if pair[1] == current {
				courseNum[pair[0]]--
			}
			if courseNum[pair[0]] == 0 {
				queue = append(queue, pair[0])
			}
		}
	}
	for i := 0; i < numCourses; i++ {
		if courseNum[i] != 0 {
			return false
		}
	}
	return true
}
