package kata

func CountBits(num uint) int {
	numOfBits := 0
	for num > 0 {
		numOfBits += int(num & 0x1)
		num >>= 1
	}
	return numOfBits
}
