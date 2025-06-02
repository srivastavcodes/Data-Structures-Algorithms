package sorting_i;

import java.util.Arrays;

public class RecursiveInsertionSort {
    public static void main(String[] args) {

        int[] arr = {13, 46, 24, 52, 20, 9, 18};
        int[] sortedArr = insertionSort(arr, 0, arr.length);
        System.out.print(Arrays.toString(sortedArr));
    }

    private static int[] insertionSort(int[] arr, int left, int right) {
        if (left >= right) return arr;

        int j = left;
        while (j > 0 && arr[j - 1] > arr[j]) {
            swapElements(arr, j - 1, j);
            j--;
        }
        return insertionSort(arr, left + 1, right);
    }

    private static void swapElements(int[] arr, int left, int right) {
        int hv = arr[left];
        arr[left] = arr[right];
        arr[right] = hv;
    }
}
