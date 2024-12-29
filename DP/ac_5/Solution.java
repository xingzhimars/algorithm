import java.util.HashSet;
import java.util.Set;

class Solution {
    public String longestPalindrome(String s) {
        int n = s.length();
        boolean[][] dp = new boolean[n][n];

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (i >= j) {
                    dp[i][j] = true;
                }
            }
        }

        int maxLen = 0;
        int start = 0;
        int end = 0;
        for (int k = 1; k < n; k++) {
            for (int i = 0; i + k < n; i++) {
                int j = i + k;
                if (s.charAt(i) == s.charAt(j)) {
                    dp[i][j] = dp[i + 1][j - 1];
                } else {
                    dp[i][j] = false;
                }
                if (dp[i][j]) {
                    if (k + 1 > maxLen) {
                        maxLen = k + 1;
                        start = i;
                        end = j;
                    }
                }
            }
        }
        return s.substring(start, end + 1);
    }
}