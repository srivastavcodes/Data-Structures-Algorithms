package revision.easy;

public class LeetCodeMaxConsecutiveOnes {
    public static void main(String[] args) {

        int[] vals = {1, 0, 1, 1, 0, 1};
        int res = findMaxConsecutiveOnes(vals);
        System.out.print(res);
    }

    private static int findMaxConsecutiveOnes(int[] vals) {
        int count = 0, largest = Integer.MIN_VALUE;

        for (int i = 0; i < vals.length; i++) {
            if (vals[i] != 0) {
                count += 1;
            } else {
                largest = Math.max(largest, count);
                count = 0;
            }
        }
        return Math.max(largest, count);
    }
}
