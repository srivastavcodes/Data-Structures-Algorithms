package revision.medium;

import java.util.Arrays;

public class QuickSortRevision {
    public static void main(String[] args) {

        int[] vals = {4, 3, 2, 5, 0, 4, 1};
        int[] res = quickSortAlgo(vals, 0, vals.length - 1);
        System.out.print(Arrays.toString(res));
    }

    private static int[] quickSortAlgo(int[] vals, int low, int high) {
        if (low < high) {
            int pIndex = getPartition(vals, low, high);
            quickSortAlgo(vals, low, pIndex - 1);
            quickSortAlgo(vals, pIndex + 1, high);
        }
        return vals;
    }

    private static int getPartition(int[] vals, int low, int high) {
        int pivot = vals[low];
        int left = low, right = high;

        while (left < right) {
            while (left < high && vals[left] <= pivot) {
                left += 1;
            }
            while (right > low && vals[right] > pivot) {
                right -= 1;
            }
            if (left < right) swapVals(vals, left, right);
        }
        swapVals(vals, low, right);
        return right;
    }

    private static void swapVals(int[] vals, int i, int j) {
        int val = vals[i];
        vals[i] = vals[j];
        vals[j] = val;
    }
}







