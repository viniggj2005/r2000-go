package utils

func calculateChecksumFromFrame(frame []byte) int {
	byteForChecksum := frame[:len(frame)-1]
	return CalculateChecksum(byteForChecksum)
}
func sum(nums []byte) int {
	total := 0
	for _, n := range nums {
		total += int(n)
	}
	return total
}

func CalculateChecksum(frameMiddleContent []byte) int {
	somaBytes := sum(frameMiddleContent)
	checksum := (256 - (somaBytes % 256)) % 256
	return checksum
}

func ValidateChecksum(frame []byte) bool {
	if len(frame) == 0 {
		return false
	}
	calculatedChecksum := calculateChecksumFromFrame(frame)
	frameChecksum := frame[len(frame)-1]

	return byte(calculatedChecksum) == frameChecksum
}
