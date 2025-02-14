package Revision2;

public class MakeBouquetsLeetCode {
    public static void main(String[] args) {
        int[] vals = {1000000000, 1000000000};
        int bouquets = 1, adjFlowers = 1;
        int output = minDays(vals, bouquets, adjFlowers);
        System.out.print(output);
    }

    private static int findMax(int[] vals) {
        int maxElement = Integer.MIN_VALUE;
        for (int val : vals) {
            maxElement = Math.max(maxElement, val);
        }
        return maxElement;
    }

    private static int minDays(int[] vals, int bouquets, int adjFlowers) {
        long count = (long) bouquets * adjFlowers;
        int low = 1, high = findMax(vals);
        if (count > vals.length) {
            return -1;
        }
        while (low <= high) {
            int centre = (low + high) / 2;
            if (isValid(vals, centre, bouquets, adjFlowers)) {
                high = centre - 1;
            } else low = centre + 1;
        }
        return low;
    }

    private static boolean isValid(int[] vals, int centre, int bouquets, int adjFlowers) {
        int count = 0, maxBouquets = 0;
        for (int val : vals) {
            if (val <= centre) {
                count++;
                if (count == adjFlowers) {
                    count = 0;
                    maxBouquets++;
                }
            } else count = 0;
        }
        return maxBouquets >= bouquets;
    }
}