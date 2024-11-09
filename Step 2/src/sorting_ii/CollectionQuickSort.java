package sorting_ii;

import java.util.Arrays;
import java.util.List;

public class CollectionQuickSort {
    public static void main(String[] args) {

        List<Integer> arr = Arrays.asList(13, 46, 24, 52, 20, 9, 18);
        List<Integer> sortedArr = quickSort(arr);
        System.out.println(sortedArr);
    }

    private static List<Integer> quickSort(List<Integer> arr) {
        sortingAlgo(arr, 0, arr.size() - 1);
        return arr;
    }

    private static void sortingAlgo(List<Integer> arr, int low, int high) {
        if (low < high) {
            int pIndex = getPartition(arr, low, high);
            sortingAlgo(arr, low, pIndex - 1);
            sortingAlgo(arr, pIndex + 1, high);
        }
    }

    private static int getPartition(List<Integer> arr, int low, int high) {
        int pivot = arr.get(low);
        int left = low, right = high;

        for (int i = left; i < right; i++) {
            while (arr.get(left) <= pivot && left <= high - 1) {
                left += 1;
            }
            while (arr.get(right) > pivot && right >= low + 1) {
                right -= 1;
            }
            if (left < right) swapListElements(arr, left, right);
        }
        swapListElements(arr, low, right);
        return right;
    }

    private static void swapListElements(List<Integer> arr, int left, int right) {
        int hv = arr.get(left);
        arr.set(left, arr.get(right));
        arr.set(right, hv);
    }
}







