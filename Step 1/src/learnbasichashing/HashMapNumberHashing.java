package learnbasichashing;

import java.util.HashMap;
import java.util.Map;

public class HashMapNumberHashing {
    public static void main(String[] args) {

        int[] arr = {1, 2, 1, 5, 12, 5, 5, 2};
        int[] digits = {1, 12, 5, 11};
        preHash(arr, digits);
    }

    private static void preHash(int[] arr, int[] digits) {
        Map<Integer, Integer> hashArr = new HashMap<>();
        for (int i : arr) {
            hashArr.put(i, hashArr.getOrDefault(i, 0) + 1);
        }
        for (int dig : digits) {
            System.out.print(hashArr.get(dig) + " ");
        }
    }
}







