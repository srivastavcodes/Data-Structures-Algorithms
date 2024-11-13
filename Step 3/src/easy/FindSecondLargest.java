package easy;

import java.util.Arrays;

public class FindSecondLargest {
    public static void main(String[] args) {

        int[] arr = {4, 7, 8, 5, 7, 6};
        int[] elmts = getSecondOrderElements(arr, arr.length - 1);
        System.out.println(Arrays.toString(elmts));
    }

    private static int[] getSecondOrderElements(int[] arr, int len) {
        int secLargest = secondLargestOptimal(arr, len);
        int secSmallest = secondSmallestOptimal(arr, len);

        return new int[] {secLargest, secSmallest};
    }
    
    private static int secondLargestOptimal(int[] arr, int len) {
        int largest = arr[0];
        int secLargest = -1;

        for (int i = 0; i < len; i++) {
            if (largest < arr[i]) {
                secLargest = largest;
                largest = arr[i];
            } else if (arr[i] < largest && arr[i] > secLargest) {
                secLargest = arr[i];
            }
        }
        return secLargest;
    }

    private static int secondSmallestOptimal(int[] arr, int len) {
        int smallest = arr[0];
        int secSmallest = Integer.MAX_VALUE;

        for (int i = 0; i < len; i++) {
            if (smallest > arr[i]) {
                secSmallest = smallest;
                smallest = arr[i];
            } else if (arr[i] > smallest && arr[i] < secSmallest) {
                secSmallest = arr[i];
            }
        }
        return secSmallest;
    }

/*
    private static int getSecondLargest(int[] arr, int len) {
        int lar = arr[0];
        int secLar = -1;

        for (int i = 0; i < len; i++) {
            if (lar < arr[i]) lar = arr[i];
        }

        for (int i = 0; i < len; i++) {
            if (arr[i] != lar && secLar < arr[i]) {
                secLar = arr[i];
            }
        }
        return secLar;
    }
*/
}







