package revision.medium;

import java.util.HashSet;
import java.util.Set;

public class LeetCodeLongestConsecutiveElement {
    public static void main(String[] args) {

        int[] vals = {100, 4, 200, 1, 3, 2};
        int res = longestConsecutive(vals);
        System.out.print(res);
    }

    private static int longestConsecutive(int[] vals) {
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
}