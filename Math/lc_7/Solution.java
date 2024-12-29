class Solution {
    public int reverse(int x) {
        long y = 0;

        boolean flag = true;
        if (x < 0) {
            flag = !flag;
            x = -x;
        }

        while (x > 0) {
            y = y * 10 + x % 10;
            x /= 10;
        }
        if (!flag) {
            y = -y;
        }

        if (y > 0 && y > Integer.MAX_VALUE) {
            return 0;
        }
        if (y < 0 && y < Integer.MIN_VALUE) {
            return 0;
        }
        return (int) y;
    }
}