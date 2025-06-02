package theory.lecture1;

import java.util.Arrays;

class LowerUpperBound {
    public static void main(String[] args) {

        int[] vals = {3, 4, 4, 7, 8, 10};
        int len = vals.length;
        int k = 8;
        System.out.print(lowerBound(vals, len, k));
        System.out.print("\n" + upperBound(vals, len, k));
        System.out.print("\n" + searchInsert(vals, k));
        System.out.print("\n" + Arrays.toString(getFloorAndCeil(vals, len, k)));
    }

    private static int lowerBound(int[] vals, int len, int k) {
        int low = 0, high = len - 1, result = len;

        while (low <= high) {
            int mid = (low + high) / 2;

            if (vals[mid] >= k) {
                result = mid;
                high = mid - 1;
            } else low = mid + 1;
        }
        return result;
    }

    private static int upperBound(int[] vals, int len, int k) {
        int low = 0, high = len - 1, result = len;

        while (low <= high) {
            int mid = (low + high) / 2;

            if (vals[mid] > k) {
                result = mid;
                high = mid - 1;
            } else low = mid + 1;
        }
        return result;
    }

    private static int searchInsert(int[] vals, int k) {
        int len = vals.length, low = 0, high = len - 1, result = len;

        while (low <= high) {
            int mid = (low + high) / 2;

            if (vals[mid] >= k) {
                result = mid;
                high = mid - 1;
            } else low = mid + 1;
        }
        return result;
    }

    private static int[] getFloorAndCeil(int[] vals, int len, int x) {
        int low = 0, high = len - 1;
        int[] result = {-1, -1};

        while (low <= high) {
            int mid = (low + high) / 2;

            if (vals[mid] <= x) {
                result[0] = vals[mid];
                low = mid + 1;
            }
            if (vals[mid] >= x) {
                result[1] = vals[mid];
                high = mid - 1;
            }
        }
        return result;
    }
}









