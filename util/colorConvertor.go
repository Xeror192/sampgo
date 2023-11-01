package util

func HexToInt(hex string) int {
	cur := 1
	res := 0

	for i := len(hex); i > 0; i-- {
		if hex[i-1] < 58 {
			res += cur * (int(hex[i-1]) - 48)
		} else {
			res += cur * (int(hex[i-1]) - 55)
		}
		cur = cur * 16
	}

	return res
}
