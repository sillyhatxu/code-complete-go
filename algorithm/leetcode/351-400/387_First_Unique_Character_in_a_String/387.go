package _87_First_Unique_Character_in_a_String

func firstUniqChar(s string) int {
	f := make([]int, 26, 26)
	for _, t := range s {
		f[t-'a']++
	}
	for i, t := range s {
		if f[t-'a'] == 1 {
			return i
		}
	}
	return -1
}
