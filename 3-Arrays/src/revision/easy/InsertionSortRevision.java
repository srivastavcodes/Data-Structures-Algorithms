package revision.easy;

import java.util.Arrays;

public class InsertionSortRevision {
    public static void main(String[] args) {

        int[] vals = {4, 3, 2, 5, 0, 4, 1};
        sortArray(vals);
        System.out.print(Arrays.toString(vals));
    }

    public static void sortArray(int[] vals) {
        for (int i = 0; i < vals.length; i++) {
            int val = i;
            while (val > 0 && vals[val - 1] > vals[val]) {
                swapVals(vals, val - 1, val--);
            }
        }
    }

    private static void swapVals(int[] vals, int i, int j) {
        int val = vals[i];
        vals[i] = vals[j];
        vals[j] = val;
    }
}







