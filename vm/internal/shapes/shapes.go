package shapes

import "fmt"

// The only valid definition of scalar.
func IsScalar(s []int) bool { return len(s) == 1 && s[0] == 1 }

// Dims must be non-empty positive array.
func IsValid(dims []int) error {
	if len(dims) == 0 {
		return fmt.Errorf("dims cannot be empty for shape.")
	}
	for _, d := range dims {
		if d <= 0 {
			return fmt.Errorf("all dims must be positive, but got: %v", dims)
		}
	}
	return nil
}

// Check OpTBROAD for definition.
func IsBroadcastable(src, dest []int) bool {
	rankSrc := len(src)
	rankDest := len(dest)

	if rankSrc > rankDest {
		return false
	}

	realRankForSrc := effectiveRank(src)

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

// Finds the output shape based on the binary operands, l and r.
//
// Examples include:
//   [1, 1]         [1]       -> [1, 1]
//   [1]            [1, 1]    -> [1, 1]
//   [2]            [3, 2]    -> [3, 2]
//   [2, 1]         [3, 2, 1] -> [3, 2, 1]
//   [1, 2, 1]      [3, 2, 1] -> [3, 2, 1]
//   [1, 1, 2, 1]   [3, 2, 1] -> [1, 3, 2, 1]
func OutputShapeForBinaryBroadcastingOp(l, r []int) ([]int, error) {
	// The algorithrm is simple.
	//
	// 1. [find effective rank]: For both operands, start from the first not-1 dim on the left, and
	//    count the dimes on the right to it, including itself.
	// 2. [righ alignment]: right aligning both operands. up to the minmal of the effective ranks, if
	//    both dimes are same, fill it. otherwise, reports an error.
	// 3. [fill the rest of dims]: fills the dims from the one with larger effective rank.
	// 4. [fill ones]: fills ones until the max of the original rank is reached.
	rankL := len(l)
	rankR := len(r)
	maxRank := rankL
	if maxRank < rankR {
		maxRank = rankR
	}

	efRankL := effectiveRank(l)
	efRankR := effectiveRank(r)
	efMaxRank := efRankL
	efMinRank := efRankR
	efMaxRankOperand := l
	efMaxRankOperandRank := rankL
	if efMaxRank < efRankR {
		efMaxRank = efRankR
		efMinRank = efRankL
		efMaxRankOperand = r
		efMaxRankOperandRank = rankR
	}

	dims := make([]int, maxRank)
	i := 0
	for i < efMinRank {
		dl := l[rankL-1-i]
		dr := r[rankR-1-i]
		if dl != dr {
			return nil, fmt.Errorf(
				"shapes are not compatible for binary op supporting broadcasting: lhs: %v, rhs %v", l, r)
		}
		dims[maxRank-1-i] = dl
		i++
	}
	for i < efMaxRank {
		dims[maxRank-1-i] = efMaxRankOperand[efMaxRankOperandRank-1-i]
		i++
	}
	for i < maxRank {
		dims[maxRank-1-i] = 1
		i++
	}
	return dims, nil
}

// -----------------------------------------------------------------------------
// helper methods.
// -----------------------------------------------------------------------------
func effectiveRank(dims []int) int {
	rank := len(dims)
	realRank := rank
	for i := 0; i < rank; i++ {
		if dims[i] != 1 {
			break
		}
		realRank--
	}
	return realRank
}
