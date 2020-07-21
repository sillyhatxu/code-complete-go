package x_solutions_chapter01_PalindromePermutation

//TODO
func PalindromePermutation(input string) bool {
	palindrome := ""
	for i := len(input) - 1; i >= 0; i-- {
		palindrome += string(input[i])
	}
	return palindrome == input
}
