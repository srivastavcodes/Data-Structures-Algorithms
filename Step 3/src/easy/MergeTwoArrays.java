package easy;

import java.util.ArrayList;
import java.util.List;

public class MergeTwoArrays {
    public static void main(String[] args) {

        int[] arr = {1, 2, 3, 4, 5};
        int[] secArr = {2, 3, 5, 7, 9};
        List<Integer> merged = sortedArray(arr, secArr);
        System.out.println(merged);
    }

    public static List<Integer> sortedArray(int[] arr, int[] secArr) {
        List<Integer> result = new ArrayList<>();
        int i = 0;
        int j = 0;

        while (i < arr.length && j < secArr.length) {
            if (arr[i] <= secArr[j]) {
                if (result.isEmpty() || result.get(result.size() - 1) != arr[i]) {
                    result.add(arr[i]);
                }
                i++;
            } else {
                if (result.isEmpty() || result.get(result.size() - 1) != secArr[j]) {
                    result.add(secArr[j]);
                }
                j++;
            }
        }
        for (int k = i; k < arr.length; k++) {
            if (result.isEmpty() || result.get(result.size() - 1) != arr[i]) {
                result.add(arr[i]);
            }
            i++;
        }
        for (int k = j; k < secArr.length; k++) {
            if (result.isEmpty() || result.get(result.size() - 1) != secArr[j]) {
                result.add(secArr[j]);
            }
            j++;
        }
        return result;
    }
}







