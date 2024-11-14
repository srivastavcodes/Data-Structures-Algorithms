package medium;

import java.util.Arrays;

public class KadanesAlgorithm {
    public static void main(String[] args) {

        int[] vals = {-2, -3, 4, -1, -2, 1, 5, -3};

        long maxSum = maxSubarraySum(vals, vals.length);
        int[] maxSumSubarr = returnMaxSumSubarray(vals);

        System.out.print(maxSum);
        System.out.println(Arrays.toString(maxSumSubarr));
    }

    public static long maxSubarraySum(int[] vals, int len) {
        long sum = 0, maxSum = Long.MIN_VALUE;

        for (int i = 0; i < len; i++) {
            sum += vals[i];
            if (sum > maxSum) maxSum = sum;
            if (sum < 0) sum = 0;
        }
        if (maxSum < 0) return 0;
        return maxSum;
    }

    public static int[] returnMaxSumSubarray(int[] vals) {
        long sum = 0, maxSum = Long.MIN_VALUE;
        int left = -1, start = -1, last = -1;

        for (int i = 0; i < vals.length; i++) {
            if (sum == 0) left = i;
            sum += vals[i];

            if (sum > maxSum) {
                maxSum = sum;
                start = left;
                last = i;
            }
            if (sum < 0) sum = 0;
        }

        int[] result = new int[last - start + 1];
        System.arraycopy(vals, start, result, 0, last - start + 1);
        return result;
    }
}







