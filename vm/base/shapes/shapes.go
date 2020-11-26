package shapes

// Check OpTBROAD for definition.
func IsBroadcastable(src, dest []int) bool {
	rankSrc := len(src)
	rankDest := len(dest)

	if rankSrc > rankDest {
		return false
	}

	realRankForSrc := rankSrc
	for i := 0; i < rankSrc; i++ {
		if src[i] != 1 {
			break
		}
		realRankForSrc--
	}

	if realRankForSrc == 0 {
		return true // all 1s.
	}

	// After right alignments, all dim sizes must agree.
	for i := 0; i < realRankForSrc; i++ {
		if src[rankSrc-1-i] != dest[rankDest-1-i] {
			return false
		}
	}
	return true
}

func IsScalar(s []int) bool { return len(s) == 1 && s[0] ==1 }
func IsValid(s []int) bool { return len(s) > 0 && s[0] ==1 }
