package easy;

public class MaxConsecutiveOne {
    public static void main(String[] args) {

        int[] arr = {1, 1, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1};
        int max = consecutiveOnes(arr.length, arr);
        System.out.println(max);
    }

    public static int consecutiveOnes(int len, int[] arr) {
        int count = 0;
        int result = 0;

        for (int i = 0; i < len; i++) {
            if (arr[i] == 1) {
                count += 1;
            } else count = 0;
            if (count > result) result = count;
        }
        return result;
    }
}
