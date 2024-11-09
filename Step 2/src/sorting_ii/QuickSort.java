package sorting_ii;

import java.util.Arrays;

public class QuickSort {
    public static void main(String[] args) {

        int[] arr = {13, 46, 24, 52, 20, 9, 18};
        int[] sortedArr = quickSort(arr, 0, arr.length - 1);
        System.out.print(Arrays.toString(sortedArr));
    }

    private static int[] quickSort(int[] arr, int low, int high) {
        if (low < high) {
            int pIndex = getPartition(arr, low, high);
            quickSort(arr, low, pIndex - 1);
            quickSort(arr, pIndex + 1, high);
        }
        return arr;
    }

    private static int getPartition(int[] arr, int low, int high) {
        int pivot = arr[low];
        int left = low, right = high;

        while (left < right) {
            while (arr[left] <= pivot && left <= high - 1) {
                left += 1;
            }
            while (arr[right] > pivot && right >= low + 1) {
                right -= 1;
            }
            if (left < right) {
                swapElements(arr, left, right);
            }
        }
        swapElements(arr, low, right);
        return right;
    }

    private static void swapElements(int[] arr, int left, int right) {
        int hv = arr[left];
        arr[left] = arr[right];
        arr[right] = hv;
    }
}







