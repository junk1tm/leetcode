// https://leetcode.com/problems/add-binary
package leetcode

// 0067. Add Binary [Easy]
//
// Given two binary strings a and b, return their sum as a binary string.
//
// Constraints:
//
//    1 <= a.length, b.length <= 104
//    a and b consist only of '0' or '1' characters.
//    Each string does not contain leading zeros except for the zero itself.
//
func AddBinary(a, b string) string {
	l1, l2 := len(a), len(b)
	l := max(l1+1, l2+1)
	b1, b2 := binary(a, l), binary(b, l)

	carry := byte(0)
	sum := make([]rune, l)
	for i := l - 1; i >= 0; i-- {
		d := b1[i] + b2[i] + carry
		carry = d / 2
		sum[i] = rune(d%2) + '0'
	}

	if sum[0] == '0' {
		return string(sum[1:])
	}
	return string(sum)
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func binary(s string, size int) []byte {
	diff := size - len(s)
	b := make([]byte, size)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '1' {
			b[i+diff] = 1
		}
	}
	return b
}
