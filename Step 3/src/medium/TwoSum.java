package medium;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class TwoSum {
    public static void main(String[] args) {

        int[] arr = {2, 6, 5, 8, 11};
        int k = 14;
        int[] result = twoSum(arr, k);
        System.out.println(Arrays.toString(result));
    }

    public static int[] twoSum(int[] arr, int k) {
        Map<Integer, Integer> integerMap = new HashMap<>();

        for (int i = 0; i < arr.length; i++) {
            int cur = arr[i];
            int re = k - cur;

            if (integerMap.containsKey(re)) {
                return new int[]{integerMap.get(re), i};
            }
            integerMap.put(arr[i], i);
        }
        return new int[]{};
    }
}







