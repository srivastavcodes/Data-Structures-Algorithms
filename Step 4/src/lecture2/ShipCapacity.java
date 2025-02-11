package lecture2;

public class ShipCapacity {
    public static void main(String[] args) {
        int[] vals = {1, 2, 3, 1, 1};
        int days = 4;
        int output = leastWeightCapacity(vals, days);
        System.out.print(output);
    }

    private static int leastWeightCapacity(int[] vals, int days) {
        int low = Integer.MIN_VALUE, high = 0;
        for (int weight : vals) {
            high += weight;
            low = Math.max(low, weight);
        }
        while (low <= high) {
            int centre = (low + high) / 2;
            if (totalDays(vals, centre) <= days) {
                high = centre - 1;
            } else low = centre + 1;
        }
        return low;
    }

    private static int totalDays(int[] vals, int cap) {
        int days = 1, load = 0;

        for (int weight : vals) {
            if (weight + load > cap) {
                days += 1;
                load = weight;
            } else load += weight;
        }
        return days;
    }
}