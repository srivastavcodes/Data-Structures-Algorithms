package Revision1;

public class LeetCodeRotatedSortedArray {
    public static void main(String[] args) {

        int[] vals = {4, 5, 6, 7, 0, 1, 2};
        int k = 0;
        int output = searchValue(vals, k);
        System.out.print(output);
    }

    private static int searchValue(int[] vals, int k) {
        int low = 0, high = vals.length - 1;

        while (low <= high) {
            int centre = (low + high) / 2;
            if (vals[centre] == k) return centre;

            if (vals[low] <= vals[centre]) {
                if (vals[low] <= k && k <= vals[centre]) {
                    high = centre - 1;
                } else {
                    low = centre + 1;
                }
            } else {
                if (vals[centre] <= k && k <= vals[high]) {
                    low = centre + 1;
                } else {
                    high = centre - 1;
                }
            }
        }
        return -1;
    }
}