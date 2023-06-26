package kata

// LCS returns the longest subsequence common to s1 and s2
func LCS(s1, s2 string) string {
	m := len(s1)
	n := len(s2)

	// Create a 2D slice to store the lengths of common subsequences.
	// dp[i][j] will hold the length of the LCS between s1[0:i] and s2[0:j].
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// Fill in the dp table.
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// If the characters match, the length of the current LCS is one greater than
			// the length of the LCS without these characters.
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// If the characters don't match, the length of the current LCS is the
				// maximum length of LCS by either excluding seq1[i] or s2[j].
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// Backtrack to find the longest common subsequence.
	lcs := ""
	i := m
	j := n
	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			// If the characters match, they are part of the LCS.
			lcs = string(s1[i-1]) + lcs
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			// If the length of LCS by excluding s1[i] is greater than the length of
			// LCS by excluding s2[j], move to the previous row.
			i--
		} else {
			// Otherwise, move to the previous column.
			j--
		}
	}

	return lcs
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
