package _42_Valid_Anagram

func isAnagram(s string, t string) bool {
	if len(s) == 0 && len(t) == 0 {
		return true
	} else if len(s) == 0 || len(t) == 0 {
		return false
	} else if len(s) != len(t) {
		return false
	}
	verify := [26]int{}
	for i := range s {
		verify[s[i]-'a']++
		verify[t[i]-'a']--
	}
	for _, temp := range verify {
		if temp < 0 {
			return false
		}
	}
	return true
}
