package theory.lecture2;

public class KthPositiveInteger {
    public static void main(String[] args) {
        int[] vals = {2, 4, 5, 7};
        int length = vals.length;
        int k = 3;
        int output = missingK(vals, length, k);
        System.out.print(output);
    }

    private static int missingK(int[] vals, int length, int k) {
        int low = 0, high = length - 1;
        while (low <= high) {
            int center = (low + high) / 2;
            int missing = vals[center] - (center + 1);
            if (missing < k) {
                low = center + 1;
            } else high = center - 1;
        }
        return low + k;
    }
}