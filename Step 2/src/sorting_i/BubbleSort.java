package sorting_i;

import java.util.Arrays;

public class BubbleSort {
    public static void main(String[] args) {

        int[] arr = {13, 46, 24, 52, 20, 9};
        System.out.print(Arrays.toString(selectSort(arr)));
    }

    private static int[] selectSort(int[] arr) {
        int swapped = 0;
        for (int i = arr.length - 1; i > 0; i--) {
            swapped = compareAndSwap(arr, i, swapped);
            if (swapped < 1) {
                break;
            }
            System.out.println("runs");
        }
        return arr;
    }

    private static int compareAndSwap(int[] arr, int i, int swapped) {
        for (int j = 0; j < i; j++) {
            if (arr[j] > arr[j + 1]) {
                swapElements(arr, j, j + 1);
                swapped++;
            }
        }
        return swapped;
    }

    private static void swapElements(int[] arr, int left, int right) {
        int hv = arr[left];
        arr[left] = arr[right];
        arr[right] = hv;
    }
}
