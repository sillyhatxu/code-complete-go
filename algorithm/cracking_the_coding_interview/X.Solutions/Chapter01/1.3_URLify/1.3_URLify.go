package x_solutions_chapter01_URlify

//TODO
func URLify(str []byte, trueLength int) []byte {
	spaceCount, index, i := 0, 0, 0
	for ; i < trueLength; i++ {
		if str[i] == ' ' {
			spaceCount++
		}
	}
	index = trueLength + spaceCount*2
	//if trueLength < len(str) {
	//	str[trueLength] = '0'
	//}
	for i := trueLength - 1; i >= 0; i-- {
		if str[i] == ' ' {
			str[index-1] = '0'
			str[index-2] = '2'
			str[index-3] = '%'
			index = index - 3
		} else {
			str[index-1] = str[i]
			index--
		}
	}
	return str
}
