package theory.medium;

import java.util.Arrays;
import java.util.HashSet;
import java.util.Set;

public class LongestConsecutiveSequence {
    public static void main(String[] args) {

        int[] vals = {5, 8, 3, 2, 1, 4};
        int res = longestSuccessiveElements(vals);
        System.out.print(res);
    }

    @SuppressWarnings("unused")
    public static int longestSuccessiveElements(int[] vals) {
        int len = vals.length, longest = 1;
        if (len == 0) return 0;
        Set<Integer> values = new HashSet<>();

        for (int i = 0; i < len; i++) {
            values.add(vals[i]);
        }
        for (Integer val : values) {
            if (!values.contains(val - 1)) {
                int count = 1, cur = val;
                while (values.contains(cur + 1)) {
                    cur += 1;
                    count += 1;
                }
                longest = Math.max(longest, count);
            }
        }
        return longest;
    }

    @SuppressWarnings("unused")
    public static int longestConsecutiveSequence(int[] vals) {
        Arrays.sort(vals); // O(nLog(n))
        int longest = 1, count = 0, lastSmallest = Integer.MIN_VALUE;

        for (int i = 0; i < vals.length; i++) {
            if (vals[i] - 1 == lastSmallest) {
                count++;
                lastSmallest = vals[i];
            } else if (vals[i] != lastSmallest) {
                count++;
                lastSmallest = vals[i];
            }
            longest = Math.max(longest, count);
        }
        return longest;
    }
}







