package sorting_i;

import java.util.Arrays;

public class InsertionSort {
    public static void main(String[] args) {

        int[] arr = {13, 46, 24, 52, 20, 9};
        System.out.print("\n" + Arrays.toString(insertionSort(arr)));
    }

    private static int[] insertionSort(int[] arr) {
        for (int i = 0; i < arr.length; i++) {
            compareAndSwap(arr, i);
        }
        return arr;
    }

    private static void compareAndSwap(int[] arr, int i) {
        for (int j = i; j > 0; j--) {
            if (arr[j - 1] > arr[j]) {
                swapElements(arr, j - 1, j);
            }
        }
    }

    private static void swapElements(int[] arr, int left, int right) {
        int hv = arr[left];
        arr[left] = arr[right];
        arr[right] = hv;
    }
}
