package lecture1;

public class FindPeakElement {
    public static void main(String[] args) {

        int[] vals = {1, 2, 1, 2, 1};
        int result = peakElement(vals);
        System.out.print(result);
    }

    private static int peakElement(int[] vals) {
        int length = vals.length, low = 1, high = length - 2;

        if (length == 1) return 0;
        if (vals[0] > vals[1]) {
            return 0;
        }
        if (vals[length - 1] > vals[length - 2]) {
            return length - 1;
        }
        while (low <= high) {
            int center = (low + high) / 2;

            if (vals[center] > vals[center + 1] && vals[center] > vals[center - 1]) {
                return center;
            } else if (vals[center] > vals[center - 1]) {
                low = center + 1;
            } else {
                high = center - 1;
            }
        }
        return -1;
    }
}