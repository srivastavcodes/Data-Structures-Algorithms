package theory.easy;

import java.util.HashMap;
import java.util.Map;

public class LongestSubArrInteger {
    public static void main(String[] args) {

        int[] arr = {1, 2, 1, 0, 1};
        int k = 4;
        int subArrLen = getLongestSubArray(arr, k);
        System.out.print(subArrLen);
    }

    public static int getLongestSubArray(int[] arr, int k) {
        Map<Integer, Integer> preSumMap = new HashMap<>();
        int sum = 0, maxLength = 0;

        for (int i = 0; i < arr.length; i++) {
            sum = sum + arr[i];
            if (sum == k) maxLength = Math.max(maxLength, i + 1);

            int isPresent = sum - k;
            if (preSumMap.containsKey(isPresent)) {
                maxLength = Math.max(maxLength, i - preSumMap.get(isPresent));
            }
            if (!preSumMap.containsKey(sum)) preSumMap.put(sum, i);
        }
        return maxLength;
    }
}







