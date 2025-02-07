package lecture1;

import java.util.Arrays;

public class CountOcc {
    public static void main(String[] args) {

        int[] vals = {1, 1, 1, 2, 2, 3, 3};
        int length = vals.length;
        int k = 1;
        int[] result = firstAndLastPosition(vals, length, k);
        int countOccurrences = count(vals, length, k);
        System.out.print(countOccurrences);
        System.out.print("\n" + Arrays.toString(result));
    }

    private static int[] firstAndLastPosition(int[] vals, int length, int k) {
        int start = startingIndex(vals, length, k);
        if (start == -1) return new int[]{-1, -1};
        int last = lastIndex(vals, length, k);
        return new int[]{start, last};
    }

    private static int startingIndex(int[] vals, int length, int k) {
        int start = 0, last = length - 1, result = -1;

        while (start <= last) {
            int centre = (start + last) / 2;

            if (vals[centre] == k) {
                result = centre;
                last = centre - 1;
            } else if (vals[centre] <= k) {
                start = centre + 1;
            } else last = centre - 1;
        }
        return result;
    }

    private static int lastIndex(int[] vals, int length, int k) {
        int start = 0, last = length - 1, result = -1;

        while (start <= last) {
            int centre = (start + last) / 2;

            if (vals[centre] == k) {
                result = centre;
                start = centre + 1;
            } else if (vals[centre] <= k) {
                start = centre + 1;
            } else last = centre - 1;
        }
        return result;
    }

    public static int count(int[] vals, int length, int k) {
        int start = startingIndex(vals, length, k);
        if (start == -1) return -1;
        int last = lastIndex(vals, length, k);
        return last - start + 1;
    }
}