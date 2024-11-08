package sorting_i;

import java.util.Arrays;

public class SelectionSort {
    public static void main(String[] args) {

        int[] arr = {13, 46, 24, 52, 20, 9};
        System.out.print(Arrays.toString(selectSort(arr)));
    }

    private static int[] selectSort(int[] arr) {
        for (int i = 0; i < arr.length - 1; i++) {
            compareAndSwap(arr, i);
        }
        return arr;
    }

    private static void compareAndSwap(int[] arr, int i) {
        for (int j = i + 1; j < arr.length; j++) {
            if (arr[i] > arr[j]) {
                swapElements(arr, i, j);
            }
        }
    }

    private static void swapElements(int[] arr, int left, int right) {
        int hv = arr[left];
        arr[left] = arr[right];
        arr[right] = hv;
    }
}







