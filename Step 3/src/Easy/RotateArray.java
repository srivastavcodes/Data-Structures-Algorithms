package Easy;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class RotateArray {
    public static void main(String[] args) {

        List<Integer> arrList = Arrays.asList(1, 2, 3, 4, 5, 6, 7);
        ArrayList<Integer> arr = new ArrayList<>(arrList);
        List<Integer> reversedArr = rotateArray(arr, 3);
        System.out.println(reversedArr);
    }

    public static ArrayList<Integer> rotateArray(ArrayList<Integer> arr, int k) {
        getReversedElements(arr, 0, k);
        getReversedElements(arr, k, arr.size());
        return getReversedElements(arr, 0, arr.size());
    }

    private static ArrayList<Integer> getReversedElements(ArrayList<Integer> arr, int left, int right) {
        while (left < right) {
            swapListElements(arr, left++, right--);
        }
        return arr;
    }

    private static void swapListElements(ArrayList<Integer> arr, int left, int right) {
        int hv = arr.get(left);
        arr.set(left, arr.get(right - 1));
        arr.set(right - 1, hv);
    }
}
