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

func isSubsequence1(s string, t string) bool {
	s_i, t_i := 0, 0
	for s_i < len(s) && t_i < len(t) {
		if s[s_i] == t[t_i] {
			s_i++
		}
		t_i++
	}
	return s_i == len(s)
}
