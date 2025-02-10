package Revision1;

public class LeetCodeSingleElementInSorted {
    public static void main(String[] args) {

        int[] vals = {1, 1, 2};
        int result = singleNonDuplicate(vals);
        System.out.print(result);
    }

    private static int singleNonDuplicate(int[] vals) {
        int low = 0, high = vals.length - 1, len = vals.length;

        if (vals.length == 1) return vals[0];

        if (vals[0] != vals[1]) {
            return vals[0];
        } else if (vals[len - 1] != vals[len - 2]) {
            return vals[len - 1];
        }

        while (low <= high) {
            int centre = (low + high) / 2;

            if (vals[centre] != vals[centre - 1] && vals[centre] != vals[centre + 1]) {
                return vals[centre];
            }
            if (centre % 2 == 0 && vals[centre] == vals[centre + 1] ||
                centre % 2 != 0 && vals[centre] == vals[centre - 1]) {
                low = centre + 1;
            } else {
                high = centre - 1;
            }
        }
        return -1;
    }
}