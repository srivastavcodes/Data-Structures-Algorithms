package learnbasichashing;

import java.util.Arrays;

public class HashingNumbers {
    public static void main(String[] args) {

        int[] arr = {8, 9};
        frequencyCount(arr, arr.length);
        System.out.println(Arrays.toString(arr));
    }

    public static void frequencyCount(int[] arr, int N) {
        int[] hashArr = new int[N];
        for (int i : arr) {
            if (i > N) continue;
            hashArr[i - 1]++;
        }
        System.arraycopy(hashArr, 0, arr, 0, hashArr.length);
    }
}
