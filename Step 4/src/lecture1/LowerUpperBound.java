package lecture1;

class LowerUpperBound {
    public static void main(String[] args) {

        int[] vals = {1, 2, 3, 4, 4, 5, 6, 7, 7, 8, 8, 9};
        int len = vals.length;
        int k = 2;
        System.out.print(lowerBound(vals, len, k));
        System.out.print("\n" + upperBound(vals, len, k));
        System.out.print(searchInsert(vals, k));
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
}
