package yenc

func getParts(in *[]byte) *[][8]byte {
	var parts [][8]byte

	var part [8]byte
	var i int

	for _, b := range *in {
		part[i] = b
		i++

		if i == 8 {
			parts = append(parts, part)
			i = 0
		}
	}

	if i != 0 {
		for i < 8 {
			part[i] = 0
			i++
		}
		parts = append(parts, part)
	}

	return &parts
}
