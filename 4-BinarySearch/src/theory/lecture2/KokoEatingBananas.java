package theory.lecture2;

public class KokoEatingBananas {
    public static void main(String[] args) {
        int[] vals = {805306368, 805306368, 805306368};
        int hrs = 1000000000;
        int result = minimumRateToEatBananas(vals, hrs);
        System.out.print(result);
    }

    private static double totalHours(int[] vals, int hrs) {
        double hours = 0;

        for (double val : vals) {
            hours += Math.ceil(val / (double) hrs);
        }
        return hours;
    }

    private static int minimumRateToEatBananas(int[] vals, int hrs) {
        int low = 1, high = 1000000000;

        while (low <= high) {
            int centre = (low + high) / 2;
            int hours = (int) totalHours(vals, centre);

            if (hours <= hrs) {
                high = centre - 1;
            } else {
                low = centre + 1;
            }
        }
        return low;
    }
}