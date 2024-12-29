class Solution {
    public int maxArea(int[] height) {
        int n = height.length;
        int i = 0; 
        int j = n-1;

        int maxArea = 0;
        while (i < j) {
            int area = min(height[i], height[j]) * (j-i);
            maxArea = max(maxArea, area);

            if (height[i] < height[j]) {
                i++;
            } else {
                j--;
            }
        }
        return maxArea;
    }

    public int min(int a, int b) {
        if (a < b) {
            return a;
        }
        return b;
    }

    public int max(int a, int b) {
        if (a > b) {
            return a;
        }
        return b;
    }
}