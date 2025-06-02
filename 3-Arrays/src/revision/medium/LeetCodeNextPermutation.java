package revision.medium;

import java.util.Arrays;

public class LeetCodeNextPermutation {
    public static void main(String[] args) {

        int[] vals = {3, 2, 1};
        nextPermutation(vals);
        System.out.print(Arrays.toString(vals));
    }

    public static void nextPermutation(int[] vals) {
        int index = -1, len = vals.length - 1;

        for (int i = len - 1; i >= 0; i--) {
            if (vals[i + 1] > vals[i]) {
                index = i;
                break;
            }
        }
        if (index == -1) {
            reverseArray(vals, 0, len);
            return;
        }
        for (int i = len; i >= 0; i--) {
            if (vals[index] < vals[i]) {
                swapVals(vals, index, i);
                break;
            }
        }
        reverseArray(vals, index + 1, len);
    }

    private static void reverseArray(int[] vals, int i, int len) {
        while (i < len) swapVals(vals, i++, len--);
    }

    private static void swapVals(int[] vals, int i, int j) {
        int val = vals[i];
        vals[i] = vals[j];
        vals[j] = val;
    }
}









