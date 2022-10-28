package _92_Is_Subsequence

func isSubsequence(s string, t string) bool {
	if s == "" {
		return false
	}
	i := 0
	current := s[i]
	for j := range t {
		if current == t[j] {
			i++
			if i >= len(s) {
				return true
			}
			current = s[i]
		}
	}
	return false
}
