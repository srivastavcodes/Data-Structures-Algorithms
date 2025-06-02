package revision.revision2;

import java.util.Arrays;

public class AggressiveCows {
    public static void main(String[] args) {
        int[] vals = {0, 3, 4, 7, 10, 9};
        int cows = 4;
        int output = aggressiveCows(vals, cows);
        System.out.print(output);
    }

    public static int aggressiveCows(int[] vals, int cows) {
        int len = vals.length;
        Arrays.sort(vals);
        int left = 0, right = vals[len - 1] - vals[0];
        while (left <= right) {
            int center = (left + right) / 2;
            if (valid(vals, center, cows)) {
                left = center + 1;
            } else {
                right = center - 1;
            }
        }
        return right;
    }

    private static boolean valid(int[] vals, int center, int cows) {
        int count = 1, lastStall = vals[0];
        for (int i = 1; i < vals.length; i++) {
            if (vals[i] - lastStall >= center) {
                count++;
                lastStall = vals[i];
            }
        }
        return count >= cows;
    }
}