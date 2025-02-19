package theory.lecture2;

import java.util.Arrays;

public class AggressiveCows {
    public static void main(String[] args) {
        int[] vals = {0, 3, 4, 7, 10, 9};
        int cows = 4;
        int output = aggressiveCows(vals, cows);
        System.out.print(output);
    }

    public static int aggressiveCows(int[] stalls, int cows) {
        Arrays.sort(stalls);
        int length = stalls.length;
        int low = 0, high = stalls[length - 1] - stalls[0];

        while (low <= high) {
            int center = (low + high) / 2;
            if (weCanPlace(stalls, center, cows)) {
                low = center + 1;
            } else high = center - 1;
        }
        return high;
    }

    private static boolean weCanPlace(int[] stalls, int center, int cows) {
        int cowCount = 1, lastStall = stalls[0];
        for (int i = 1; i < stalls.length; i++) {
            if (stalls[i] - lastStall >= center) {
                cowCount++;
                lastStall = stalls[i];
            }
            if (cowCount >= cows) return true;
        }
        return false;
    }
}