package lecture2;

public class RoseGarden {
    public static void main(String[] args) {

        int[] vals = {1, 2, 1, 2, 7, 2, 2, 3, 1};
        int r = 3, b = 2;
        int output = roseGarden(vals, r, b);
        System.out.print(output);
    }

    private static int roseGarden(int[] vals, int r, int b) {
        int low = 1, high = 1000000000;
        if ((r * b) > vals.length) return -1;
        while (low <= high) {
            int centre = (low + high) / 2;
            if (isPossible(vals, centre, r, b)) {
                high = centre - 1;
            } else low = centre + 1;
        }
        return low;
    }

    private static boolean isPossible(int[] vals, int day, int r, int b) {
        int count = 0, bokeh = 0;

        for (int val : vals) {
            if (val <= day) {
                count++;
            } else {
                bokeh += count / r;
                count = 0;
            }
        }
        bokeh += count / r;
        return bokeh >= b;
    }
}