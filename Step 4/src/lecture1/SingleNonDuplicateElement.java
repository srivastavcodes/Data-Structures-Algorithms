package lecture1;

public class SingleNonDuplicateElement {
    public static void main(String[] args) {

        int[] vals = {1, 1, 2, 3, 3, 4, 4, 8, 8};
        int result = singleNonDuplicate(vals);
        System.out.print(result);
    }

    private static int singleNonDuplicate(int[] vals) {
        int len = vals.length, low = 1, high = len - 2;

        if (len == 1) return vals[0];
        if (vals[0] != vals[1]) {
            return vals[0];
        }
        if (vals[len - 1] != vals[len - 2]) {
            return vals[len - 1];
        }
        while (low <= high) {
            int centre = (low + high) / 2;
            if (vals[centre] != vals[centre + 1] && vals[centre] != vals[centre - 1]) {
                return vals[centre];
            }
            if (centre % 2 == 0 && vals[centre] == vals[centre + 1] ||
                centre % 2 == 1 && vals[centre] == vals[centre - 1]) {
                low = centre + 1;
            } else {
                high = centre - 1;
            }
        }
        return -1;
    }
}