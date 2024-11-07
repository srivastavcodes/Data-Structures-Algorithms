package learnbasicrecursion;

import java.util.Arrays;

public class ReverseAnArray {
    public static void main(String[] args) {

        int[] arr = {2, 3, 4, 5, 6, 8, 10};
        System.out.print(Arrays.toString(reverseArray(arr)));
    }

    private static int[] reverseArray(int[] arr) {
        int n = arr.length;
        return getReversedArray(arr, 0, n);
    }

    private static int[] getReversedArray(int[] arr, int i, int n) {
        if (i >= n / 2) {
            return arr;
        }
        swapElements(arr, i, n);
        return getReversedArray(arr, i + 1, n);
    }

    private static void swapElements(int[] arr, int i, int n) {
        int temp = arr[i];
        arr[i] = arr[n - i - 1];
        arr[n - i - 1] = temp;
    }
}







