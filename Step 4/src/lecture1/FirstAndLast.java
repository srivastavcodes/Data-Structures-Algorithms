package lecture1;

import java.util.ArrayList;
import java.util.Arrays;

class FirstAndLast {
    public static void main(String[] args) {

        ArrayList<Integer> valList = new ArrayList<>(Arrays.asList(0, 1, 1, 5));
        int length = valList.size();
        int x = 1;
        System.out.println("\n" + Arrays.toString(firstAndLastPosition(valList, length, x)));
    }

    private static int[] firstAndLastPosition(ArrayList<Integer> vals, int len, int k) {
        int lb = lowerBound(vals, len, k);
        if (lb == len || !vals.get(lb).equals(k)) return new int[]{-1, -1};
        return new int[]{lb, upperBound(vals, len, k) - 1};
    }

    private static int lowerBound(ArrayList<Integer> vals, int len, int k) {
        int low = 0, high = len - 1, result = len;

        while (low <= high) {
            int mid = (low + high) / 2;

            if (vals.get(mid) >= k) {
                result = mid;
                high = mid - 1;
            } else low = mid + 1;
        }
        return result;
    }

    private static int upperBound(ArrayList<Integer> vals, int len, int k) {
        int low = 0, high = len - 1, result = len;

        while (low <= high) {
            int mid = (low + high) / 2;

            if (vals.get(mid) > k) {
                result = mid;
                high = mid - 1;
            } else low = mid + 1;
        }
        return result;
    }
}
