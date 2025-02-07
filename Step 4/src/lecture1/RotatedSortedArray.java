package lecture1;

import java.util.ArrayList;
import java.util.Arrays;

public class RotatedSortedArray {
    public static void main(String[] args) {

        ArrayList<Integer> vals = new ArrayList<>(Arrays.asList(12, 15, 18, 2, 4));
        int len = vals.size(), k = 2;
        System.out.print(search(vals, len, k));
    }

    private static int search(ArrayList<Integer> vals, int len, int k) {
        int low = 0, high = len - 1;

        while (low <= high) {
            int centre = (low + high) / 2;
            if (vals.get(centre) == k) return centre;

            if (vals.get(low) <= vals.get(centre)) {
                if (vals.get(low) <= k && k <= vals.get(centre)) {
                    high = centre - 1;
                } else {
                    low = centre + 1;
                }
            } else {
                if (vals.get(centre) <= k && k <= vals.get(high)) {
                    low = centre + 1;
                } else {
                    high = centre - 1;
                }
            }
        }
        return -1;
    }
}