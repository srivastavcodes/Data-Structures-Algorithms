package theory.lecture2;

public class MedianSortedArray {
    public static void main(String[] args) {
        int[] vals1 = {2, 4, 6};
        int[] vals2 = {1, 3, 5};
        double output = median(vals1, vals2);
        System.out.print(output);
    }

    public static double median(int[] vals1, int[] vals2) {
        int len1 = vals1.length, len2 = vals2.length;
        if (len1 > len2) {
            return median(vals2, vals1);
        }
        int len = len1 + len2, left = (len1 + len2 + 1) / 2;
        int low = 0, high = len1;
        while (low <= high) {
            int centre1 = (low + high) / 2;
            int centre2 = left - centre1;

            // calculating l1, l2
            int l1 = (centre1 > 0) ? vals1[centre1 - 1] : Integer.MIN_VALUE;
            int l2 = (centre2 > 0) ? vals2[centre2 - 1] : Integer.MIN_VALUE;
            // calculating r1,r2
            int r1 = (centre1 < len1) ? vals1[centre1] : Integer.MAX_VALUE;
            int r2 = (centre2 < len2) ? vals2[centre2] : Integer.MAX_VALUE;

            if (l1 <= r2 && l2 <= r1) {
                if (len % 2 == 1) return Math.max(l1, l2);
                else {
                    return ((double) Math.max(l1, l2) + Math.min(r1, r2)) / 2.0;
                }
            } else if (l1 > r2) {
                high = centre1 - 1;
            } else low = centre1 + 1;
        }
        return 0;
    }
}
