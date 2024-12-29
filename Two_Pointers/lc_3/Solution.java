import java.util.HashSet;
import java.util.Set;

class Solution {
    public int lengthOfLongestSubstring(String s) {
        int n = s.length();
        Set<Character> set = new HashSet<>();
        int i = 0;
        int j = 0;

        int maxLen = 0;
        while (i <= j && j < n) {
            if (!set.contains(s.charAt(j))) {
                set.add(s.charAt(j));
                j++;

                if (j - i > maxLen) {
                    maxLen = j - i;
                }
            } else {
                set.remove(s.charAt(i));
                i++;
            }
        }
        return maxLen;
    }
}