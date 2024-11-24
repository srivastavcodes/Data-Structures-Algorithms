package revision.easy;

import java.util.HashMap;
import java.util.Map;

public class GFGLongestSubarrWithSumK {
    public static void main(String[] args) {

        int[] vals = {10, 5, 2, 7, 1, 9};
        int k = 15;
        int len = lengthOfLongestSubarr(vals, k);
        System.out.print(len);
    }

    // For both positives and negatives
    @SuppressWarnings("unused")
    private static int lenOfLongestSubarr(int[] vals, int k) {
        Map<Integer, Integer> preSum = new HashMap<>();
        int sum = 0, maxLen = 0;

        for (int i = 0; i < vals.length; i++) {
            sum += vals[i];
            if (sum == k) maxLen = Math.max(maxLen, i + 1);

            int reqVal = sum - k;
            if (preSum.containsKey(reqVal)) {
                maxLen = Math.max(maxLen, i - preSum.get(reqVal));
            }
            if (!preSum.containsKey(sum)) preSum.put(sum, i);
        }
        return maxLen;
    }

    // For positives only
    private static int lengthOfLongestSubarr(int[] vals, int k) {
        int left = 0, right = 0, sum = 0, maxLen = 0;

        while (right < vals.length) {
            while (left <= right && sum > k) sum -= vals[left++];
            if (sum == k) maxLen = Math.max(maxLen, right - left);
            sum += ++right < vals.length ? vals[right] : 0;
        }
        return maxLen;
    }
}







