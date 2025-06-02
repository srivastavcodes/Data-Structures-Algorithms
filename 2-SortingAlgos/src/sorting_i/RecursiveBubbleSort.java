package sorting_i;

import java.util.Arrays;

public class RecursiveBubbleSort {
    public static void main(String[] args) {

        int[] unsorted = {13, 46, 24, 52, 20, 9, 18};
        int[] sortedArr = bubbleSort(unsorted, 0, unsorted.length - 1);
        System.out.print(Arrays.toString(sortedArr));
    }

    private static int[] bubbleSort(int[] arr, int left, int right) {
        if (right <= 0) return arr;

        if (left < right) {
            if (arr[left] > arr[left + 1]) swapElements(arr, left, left + 1);
            bubbleSort(arr, left + 1, right);
        }
        return bubbleSort(arr, 0, right - 1);
    }

    private static void swapElements(int[] arr, int left, int right) {
        int hv = arr[left];
        arr[left] = arr[right];
        arr[right] = hv;
    }
}








