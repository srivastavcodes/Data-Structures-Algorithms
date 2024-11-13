package Easy;

public class LongestSubArrNatural {
    public static void main(String[] args) {

        int[] arr = {1, 2, 3, 1, 1, 1, 1};
        long k = 3;
        int subArrLen = longestSubarrayWithSumK(arr, k);
        System.out.print(subArrLen);
    }

    public static int longestSubarrayWithSumK(int[] arr, long k) {
        int left = 0, right = 0, maxLen = 0;
        long sum = arr[0];

        while (right < arr.length) {
            while (left <= right && sum > k) {
                sum -= arr[left];
                left += 1;
            }
            if (sum == k) {
                maxLen = Math.max(maxLen, right - left + 1);
            }
            right += 1;
            if (right < arr.length) sum += arr[right];
        }
        return maxLen;
    }
}







