package revision.medium;

import java.util.Arrays;

public class MergeSortRevision {
    public static void main(String[] args) {

        int[] vals = {4, 3, 2, 5, 0, 4, 1};
        int[] res = mergeSortAlgo(vals);
        System.out.print(Arrays.toString(res));
    }

    public static int[] mergeSortAlgo(int[] vals) {
        if (vals.length <= 1) return vals;

        int mid = vals.length / 2;
        int[] left = mergeSortAlgo(Arrays.copyOfRange(vals, 0, mid));
        int[] right = mergeSortAlgo(Arrays.copyOfRange(vals, mid, vals.length));
        return sortAndMerge(left, right);
    }

    private static int[] sortAndMerge(int[] left, int[] right) {
        int[] sorted = new int[left.length + right.length];
        int i = 0, j = 0, k = 0;

        while (i < left.length && j < right.length) {
            if (left[i] > right[j]) {
                sorted[k++] = right[j++];
            } else {
                sorted[k++] = left[i++];
            }
        }
        for (int l = i; l < left.length; l++) {
            sorted[k++] = left[l];
        }
        for (int l = j; l < right.length; l++) {
            sorted[k++] = right[l];
        }
        return sorted;
    }
}







