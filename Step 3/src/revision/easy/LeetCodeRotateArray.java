package revision.easy;

import java.util.Arrays;

public class LeetCodeRotateArray {
    public static void main(String[] args) {

        int[] vals = {1, 2, 3, 4, 5, 6, 7};
        int k = 3;
        rotateArray(vals, k);
        System.out.println(Arrays.toString(vals));
    }

    public static void rotateArray(int[] vals, int con) {
        if (vals.length <= 1) return;
        int k = con % vals.length;

        reverseArray(vals, 0, vals.length - k - 1);
        reverseArray(vals, vals.length - k, vals.length - 1);
        reverseArray(vals, 0, vals.length - 1);
    }

    private static void reverseArray(int[] vals, int i, int j) {
        while (i < j) swapVals(vals, i++, j--);
    }

    private static void swapVals(int[] vals, int i, int j) {
        int val = vals[i];
        vals[i] = vals[j];
        vals[j] = val;
    }
}







