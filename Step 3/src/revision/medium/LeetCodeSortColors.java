package revision.medium;

import java.util.Arrays;

public class LeetCodeSortColors {
    public static void main(String[] args) {

        int[] vals = {1, 2, 0};
        sortColors(vals);
        System.out.print(Arrays.toString(vals));
    }

    private static void sortColors(int[] vals) {
        int left = 0, centre = 0, right = vals.length - 1;
        while (centre <= right) {
            if (vals[centre] == 0) {
                swapVals(vals, left++, centre++);
            } else if (vals[centre] == 1) {
                centre += 1;
            } else {
                swapVals(vals, centre, right--);
            }
        }
    }

    private static void swapVals(int[] vals, int i, int j) {
        int val = vals[i];
        vals[i] = vals[j];
        vals[j] = val;
    }
}







