package revision1;

import java.util.Arrays;

public class LeetCodeFindFirstAndLast {
    public static void main(String[] args) {

        int[] vals = {1, 2, 3, 4, 5, 5, 6, 7, 8, 9};
        int k = 5;
        int[] result = searchRange(vals, k);
        int count = countFreq(vals, k);
        System.out.print(Arrays.toString(result));
        System.out.print("\n" + count);
    }

    private static int[] searchRange(int[] vals, int k) {
        int start = findFirstIndex(vals, k);
        if (start == -1) return new int[]{-1, -1};
        int last = findLastIndex(vals, k);
        return new int[]{start, last};
    }

    public static int countFreq(int[] vals, int k) {
        int start = findFirstIndex(vals, k);
        if (start == -1) return 0;
        int last = findLastIndex(vals, k);
        return last - start + 1;
    }

    private static int findLastIndex(int[] vals, int k) {
        int low = 0, high = vals.length - 1, output = -1;

        while (low <= high) {
            int center = (low + high) / 2;

            if (vals[center] == k) {
                output = center;
                low = center + 1;
            } else if (vals[center] <= k) {
                low = center + 1;
            } else {
                high = center - 1;
            }
        }
        return output;
    }

    private static int findFirstIndex(int[] vals, int k) {
        int low = 0, high = vals.length - 1, output = -1;

        while (low <= high) {
            int center = (low + high) / 2;

            if (vals[center] == k) {
                output = center;
                high = center - 1;
            } else if (vals[center] <= k) {
                low = center + 1;
            } else {
                high = center - 1;
            }
        }
        return output;
    }
}