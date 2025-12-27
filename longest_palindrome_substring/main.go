package main

import (
	"fmt"
)

// for {
// 	orig := string(subSr)

// 	if orig == string(rev) {
// 		return orig
// 	} else {
// 		if endIdx > len(sr) {
// 			startIdx = 0
// 			endIdx = len(subSr) - 1
// 		}

// 		subSr = sr[startIdx:endIdx]
// 		rev = slices.Clone(subSr)
// 		slices.Reverse(rev)
// 		startIdx++
// 		endIdx++
// 	}
// }

// for i := endIdx; i >= 0; i-- {
// 	rev[len(subSb)-1-i] = subSb[i]

// 	if i > 0 {
// 		continue
// 	}

// 	orig := unsafe.String(unsafe.SliceData(subSb), len(subSb))
// 	reversed := unsafe.String(unsafe.SliceData(rev), len(rev))

// 	if orig == reversed {
// 		return orig
// 	} else {
// 		if endIdx > len(sb) {
// 			startIdx = 0
// 			endIdx = len(subSb) - 1
// 		}

// 		subSb = sb[startIdx:endIdx]
// 		i = len(subSb) // due to i-- on next step
// 		rev = make([]byte, len(subSb))
// 		startIdx++
// 		endIdx++
// 	}
// }

// func isPalindrome(s string, l, r int) bool {
// 	for l < r {
// 		if s[l] != s[r] {
// 			return false
// 		}
// 		l++
// 		r--
// 	}

// 	return true
// }

// func longestPalindrome(s string) string {
// 	currLen := len(s)
// 	startIdx := 0
// 	endIdx := currLen - 1

// 	for {
// 		if endIdx == len(s) {
// 			startIdx = 0
// 			currLen--
// 			endIdx = currLen - 1
// 		}

// 		if isPalindrome(s, startIdx, endIdx) {
// 			return s[startIdx : endIdx+1]
// 		}

// 		startIdx++
// 		endIdx++
// 	}
// }

func expand(s string, l, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return l, r
}

func longestPalindrome(s string) string {
	start, end := len(s), -1
	for i := range len(s) {
		l, r := expand(s, i, i)
		if r-l > end-start {
			start = l
			end = r
		}

		l, r = expand(s, i, i+1)
		if r-l > end-start {
			start = l
			end = r
		}
	}

	return s[start+1 : end]
}

func main() {
	// fmt.Println(isPalindrome("abcba", 1, 3))
	fmt.Println(longestPalindrome("bcaxcba"))
	// fmt.Println(longestPalindrome("jcwwnkwiajicysmdueefqjnrokunucidhgkswbgjkkrujkxkxeanrpjvpliomxztalhmvcldnqmkslkprhgtwlnsnygbzdvtdbsdzsdjggzgmhogknpfhtgjmclrkgfqdbiagwrqqcnagosnqrnpapxfrtrhzlyhszdtgkqggmewqmwugrbufiwfvtjhczqgcwpcyjioeacggiwyinpkyxrpxhglrtojgjmmswxnvirfsrbhmpqgjyyagjqfwkqkjkumywvgfutmiwihwnnhbfxcijaoiyxbdnrckexqfhsmmxflaneccyaoqoxfbaylcvvpfafsikebzcdufvhluldguwsmrtjaljpcjestranfrvgvytozxpcvnwquhnsxlmzkehwopgxvifajmrlwqiqylgxibnypxwpkggxevyfoxywfsfpjbzfxxysgxgwxrzkwdqlkrpajlltzqfqshdokibakkhydizsgwbygqulljqmtxkwepitsukwjlrrmsjptwabtlqytprkkuxtqdmptctkfabxsohrfrqvrbjfbppfkpthosveoppiywkkuoasefviegormlqkqlbhnhblkmglxcbszblfipsyumcrjrkrnzplkveznbtdbtlcptgswhiqsjugqrvujkzuwvoxbjremyxqqzxkgerhgtidsefyemtmstsznvgohexdgetqbhrxaomzsamapxhjibfvtbquhowyrwyxthpwvmfyyqsyibemnfbwkddtyoijzwfxhossylygxmnznpegtgvlrsreepkrcdgbujkghrgtsxwlvxrgrqxnvgqkppbkrxjupjfjcsfzepdemaulfetn"))
}
