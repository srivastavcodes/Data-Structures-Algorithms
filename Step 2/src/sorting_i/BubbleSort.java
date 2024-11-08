package sorting_i;

import java.util.Arrays;

public class BubbleSort {
    public static void main(String[] args) {

        int[] arr = {13, 46, 24, 52, 20, 9};
        System.out.print("\n" + Arrays.toString(selectSort(arr)));
    }

    private static int[] selectSort(int[] arr) {
        int swaps = 0;

        for (int i = arr.length - 1; i > 0; i--) {
            swaps = compareAndSwap(arr, i, swaps);
            if (swaps < 1) break;

            System.out.print("ran ");
        }
        return arr;
    }

    private static int compareAndSwap(int[] arr, int i, int swaps) {
        for (int j = 0; j < i; j++) {
            if (arr[j] > arr[j + 1]) {
                swapElements(arr, j, j + 1);
                swaps++;
            }
        }
        return swaps;
    }

    private static void swapElements(int[] arr, int left, int right) {
        int hv = arr[left];
        arr[left] = arr[right];
        arr[right] = hv;
    }
}
