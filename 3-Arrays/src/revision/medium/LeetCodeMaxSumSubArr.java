package revision.medium;

public class LeetCodeMaxSumSubArr {
    public static void main(String[] args) {

        int[] vals = {2, -1, 4};
        int maxSum = maxSubArr(vals);
        System.out.print(maxSum);
    }

    private static int maxSubArr(int[] vals) {
        int sum = 0, maxSum = Integer.MIN_VALUE;
        for (int i = 0; i < vals.length; i++) {
            if (sum < 0) sum = 0;
            sum += vals[i];
            if (sum > maxSum) maxSum = sum;
        }
        return maxSum;
    }
}
