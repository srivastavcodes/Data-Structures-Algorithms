package sorting_ii;

import java.util.Arrays;

public class MergeSort {
    public static void main(String[] args) {

        int[] unsorted = {13, 46, 24, 52, 20, 9, 18};
        int[] sortedArr = mergeSort(unsorted);
        System.out.print(Arrays.toString(sortedArr));
    }

    private static int[] mergeSort(int[] unsorted) {
        if (unsorted.length == 1) {
            return unsorted;
        }
        int mid = unsorted.length / 2;
        int[] left = mergeSort(Arrays.copyOfRange(unsorted, 0, mid));
        int[] right = mergeSort(Arrays.copyOfRange(unsorted, mid, unsorted.length));
        return sortArr(left, right);
    }

    private static int[] sortArr(int[] left, int[] right) {
        int[] result = new int[left.length + right.length];
        int i = 0, j = 0, k = 0;

        while (i < left.length && j < right.length) {
            if (left[i] <= right[j]) {
                result[k++] = left[i++];
            } else {
                result[k++] = right[j++];
            }
        }
        for (int l = i; l < left.length; l++) {
            result[k++] = left[i++];
        }
        for (int r = j; r < right.length; r++) {
            result[k++] = right[j++];
        }
        return result;
    }
}







