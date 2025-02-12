package Revision2;

public class ShipCapacityLeetCode {
    public static void main(String[] args) {
        int[] vals = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
        int days = 5;
        int output = shipWithinDays(vals, days);
        System.out.print(output);
    }

    public static int shipWithinDays(int[] vals, int days) {
        int low = findRange(vals)[0], high = findRange(vals)[1];
        while (low <= high) {
            int centre = (low + high) / 2;
            if (isValid(vals, centre, days)) {
                high = centre - 1;
            } else low = centre + 1;
        }
        return low;
    }

    private static boolean isValid(int[] vals, int centre, int days) {
        int load = 0, dayTaken = 1;
        for (int val : vals) {
            if (load + val > centre) {
                dayTaken += 1;
                load = val;
            } else load += val;
        }
        return dayTaken <= days;
    }

    private static int[] findRange(int[] vals) {
        int low = Integer.MIN_VALUE, high = 0;
        for (int val : vals) {
            high += val;
            if (val > low) low = val;
        }
        return new int[]{low, high};
    }
}