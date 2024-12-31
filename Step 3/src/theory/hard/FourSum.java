package theory.hard;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class FourSum {

    public static void main(String[] args) {

        int[] vals = {1, 2, 1, 0, 1};
        int goal = 5;
        List<List<Integer>> reslist = fourSum(vals, goal);
        System.out.print(reslist);
    }

    public static List<List<Integer>> fourSum(int[] vals, int goal) {
        List<List<Integer>> reslist = new ArrayList<>();
        Arrays.sort(vals);
        int len = vals.length;

        for (int i = 0; i < len - 3; i++) {
            if (i > 0 && vals[i] == vals[i - 1]) continue;

            for (int j = i + 1; j < len; j++) {
                if (j != i + 1 && vals[j] == vals[j - 1]) continue;

                int k = j + 1, l = len - 1;
                while (k < l) {
                    int sum = vals[i] + vals[j] + vals[k] + vals[l];

                    if (sum < goal) k += 1;
                    else if (sum > goal) {
                        l -= 1;
                    } else {
                        reslist.add(List.of(vals[i] + vals[j] + vals[k] + vals[l]));
                        k++;
                        l--;
                        while (k < l && vals[k] == vals[k - 1]) k++;
                        while (k < l && vals[l] == vals[l + 1]) l--;
                    }
                }
            }
        }
        return reslist;
    }
}









