package revision.revision1;

public class FindPeakElementLeetCode {
    public static void main(String[] args) {

        int[] vals = {1, 2, 1, 3, 5, 6, 4};
        int result = findPeakElement(vals);
        System.out.print(result);
    }

    private static int findPeakElement(int[] vals) {
        int low = 1, high = vals.length - 2, length = vals.length;

        if (length == 1) return 0;
        if (vals[0] > vals[1]) {
            return 0;
        } else if (vals[length - 1] > vals[length - 2]) {
            return length - 1;
        }
        while (low <= high) {
            int centre = (low + high) / 2;

            if (vals[centre] > vals[centre + 1] &&
                vals[centre] > vals[centre - 1]) {
                return centre;
            } else if (vals[centre] > vals[centre - 1]) {
                low = centre + 1;
            } else {
                high = centre - 1;
            }
        }
        return Integer.MIN_VALUE;
    }
}