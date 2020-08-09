package main

/**
 * 有效的括号：https://leetcode-cn.com/problems/valid-parentheses/
 */

func isValid(s string) bool {
	m := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	lps := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		p := s[i]
		if _, ok := m[p]; ok {
			lps = append(lps, p)
		} else {
			if len(lps) > 0 && p == m[lps[len(lps)-1]] {
				lps = lps[:len(lps)-1]
			} else {
				return false
			}
		}
	}
	if len(lps) == 0 {
		return true
	}
	return false
}
