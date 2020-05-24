package prng64

const truncatedBound float32 = 2.0

func (prng *Prng64) TruncNorm(value []float32) {
	size := len(value)

	i := 0
	for {
		r1, r2 := prng.boxMullerTransform()

		if r1 <= truncatedBound && r1 >= -truncatedBound {
			value[i] = r1
			i++
			if i == size {
				break
			}
		}

		if r2 <= truncatedBound && r2 >= -truncatedBound {
			value[i] = r2
			i++
			if i == size {
				break
			}
		}
	}
}
