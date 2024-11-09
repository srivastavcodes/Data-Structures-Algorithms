package sorting_i;

import java.util.Arrays;

public class RecursiveSelectionSort {
    public static void main(String[] args) {

        int[] arr = {13, 46, 24, 52, 20, 9, 18};
        int[] sortedArr = selectionSort(arr, 0, arr.length - 1);
        System.out.print(Arrays.toString(sortedArr));
    }

    private static int[] selectionSort(int[] arr, int left, int right) {
        if (left >= right) return arr;

        int minIndex = left;
        for (int i = left + 1; i <= right; i++) {
            if (arr[minIndex] > arr[i]) minIndex = i;
        }
        swapElements(arr, left, minIndex);
        return selectionSort(arr, left + 1, right);
    }

    private static void swapElements(int[] arr, int left, int right) {
        int hv = arr[left];
        arr[left] = arr[right];
        arr[right] = hv;
    }
}
