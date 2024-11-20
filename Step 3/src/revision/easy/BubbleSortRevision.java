package revision.easy;

import java.util.Arrays;

public class BubbleSortRevision {
    public static void main(String[] args) {

        int[] vals = {4, 3, 2, 5, 0, 4, 1};
        sortArray(vals);
        System.out.print(Arrays.toString(vals));
    }

    private static void sortArray(int[] vals) {
        int swaps = -1;
        for (int i = vals.length - 1; i > 0; i--) {
            for (int j = 0; j < i; j++) {
                if (vals[i] < vals[j]) {
                    swapVals(vals, i, j);
                    swaps += 1;
                }
            }
            if (swaps < 0) break;
        }
    }

    private static void swapVals(int[] vals, int i, int j) {
        int val = vals[i];
        vals[i] = vals[j];
        vals[j] = val;
    }
}







