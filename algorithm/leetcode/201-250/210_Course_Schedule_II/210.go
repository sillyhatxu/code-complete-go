package _10_Course_Schedule_II

func findOrder(numCourses int, prerequisites [][]int) []int {
	courseList := make([]int, numCourses, numCourses)
	var res []int
	for i := 0; i < len(prerequisites); i++ {
		courseList[prerequisites[i][0]]++
	}
	var queue []int
	for i := 0; i < len(courseList); i++ {
		if courseList[i] == 0 {
			queue = append(queue, i)
			res = append(res, i)
		}
	}
	for len(queue) != 0 {
		finishCourse := queue[0]
		queue = queue[1:]
		for _, pair := range prerequisites {
			if courseList[pair[0]] == 0 {
				continue
			}
			if pair[1] == finishCourse {
				courseList[pair[0]]--
			}
			if courseList[pair[0]] == 0 {
				queue = append(queue, pair[0])
				res = append(res, pair[0])
			}
		}
	}
	for i := 0; i < len(courseList); i++ {
		if courseList[i] != 0 {
			return []int{}
		}
	}
	return res
}
