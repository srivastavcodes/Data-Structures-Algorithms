package revision.hard;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class LeetCodeThreeSum {
    public static void main(String[] args) {

        int[] vals = {-1, 0, 1, 2, -1, -4};
        List<List<Integer>> reslist = threeSum(vals);
        System.out.print(reslist);
    }

    public static List<List<Integer>> threeSum(int[] vals) {
        List<List<Integer>> reslist = new ArrayList<>();
        Arrays.sort(vals);
        int len = vals.length;

        for (int i = 0; i < len - 2; i++) {
            if (i > 0 && vals[i] == vals[i - 1]) continue;

            int j = i + 1, k = len - 1;
            while (j < k) {

                int sum = vals[i] + vals[j] + vals[k];
                if (sum < 0) j += 1;
                else if (sum > 0) {
                    k -= 1;
                } else {
                    reslist.add(List.of(vals[i], vals[j], vals[k]));
                    j++;
                    k--;
                    while (j < k && vals[j] == vals[j - 1]) j++;
                    while (j < k && vals[k] == vals[k + 1]) k--;
                }
            }
        }
        return reslist;
    }
}









