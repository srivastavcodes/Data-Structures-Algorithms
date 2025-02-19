package revision.revision1;

public class LeetCodeMinimumInSortedArray {
    public static void main(String[] args) {

        int[] vals = {1, 2, 3, 4, 5, 6};
        int result = findMinimum(vals);
        System.out.print(result);
    }

    private static int findMinimum(int[] vals) {
        int low = 0, high = vals.length - 1, output = Integer.MAX_VALUE;

        while (low <= high) {
            int centre = (low + high) / 2;

            if (vals.length == 1) return vals[0];

            if (vals[low] <= vals[centre]) {
                if (output > vals[low]) {
                    output = vals[low];
                }
                low = centre + 1;
            } else {
                if (output > vals[centre]) {
                    output = vals[centre];
                }
                high = centre - 1;
            }
        }
        return output;
    }
}