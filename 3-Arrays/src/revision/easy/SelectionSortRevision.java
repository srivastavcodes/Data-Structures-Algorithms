package revision.easy;

import java.util.Arrays;

public class SelectionSortRevision {
    public static void main(String[] args) {

        int[] vals = {4, 3, 2, 5, 4, 1};
        sortArray(vals);
        System.out.print(Arrays.toString(vals));
    }

    private static void sortArray(int[] vals) {
        int len = vals.length;

        for (int i = 0; i < len - 1; i++) {
            for (int j = i + 1; j < len; j++) {
                if (vals[i] > vals[j]) swapVals(vals, i, j);
            }
        }
    }

    private static void swapVals(int[] vals, int i, int j) {
        int val = vals[i];
        vals[i] = vals[j];
        vals[j] = val;
    }
}







