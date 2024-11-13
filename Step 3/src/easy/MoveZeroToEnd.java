package easy;

import java.util.Arrays;

public class MoveZeroToEnd {
    public static void main(String[] args) {

        int[] arr = {0, 0, 0, 1};
        int[] res = moveZeros(arr, arr.length - 1);
        System.out.println(Arrays.toString(res));
    }

    private static int[] moveZeros(int[] arr, int len) {
        int right = -1;

        for (int i = 0; i <= len; i++) {
            if (arr[i] == 0) {
                right = arr[i];
                break;
            }
        }
        // if no zero found.
        if (right == -1) return arr;

        for (int i = 0; i <= len; i++) {
            if (arr[i] != 0) {
                swapElements(arr, i, right);
                right += 1;
            }
        }
        return arr;
    }

    private static void swapElements(int[] arr, int left, int right) {
        int hvar = arr[left];
        arr[left] = arr[right];
        arr[right] = hvar;
    }
}







