package lecture2;

public class SmallestDivisor {
    public static void main(String[] args) {

        int[] vals = {8, 4, 2, 3};
        int limit = 10;
        int output = smallestDivisor(vals, limit);
        System.out.print(output);
    }

    private static int smallestDivisor(int[] vals, int limit) {
        int low = 1, high = 1000000000;
        while (low <= high) {
            int centre = (low + high) / 2;
            if (divisor(vals, centre) <= limit) {
                high = centre - 1;
            } else {
                low = centre + 1;
            }
        }
        return low;
    }

    private static double divisor(int[] vals, int divisor) {
        double sum = 0;
        for (double val : vals) {
            sum += Math.ceil(val / (double) divisor);
        }
        return sum;
    }
}