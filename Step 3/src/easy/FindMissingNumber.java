package easy;

public class FindMissingNumber {
    public static void main(String[] args) {

        int[] arr = {1, 2, 3, 4, 6};
        int missingNumSum = missingNumberSum(arr, arr.length + 1);
        int missingNum = missingNumber(arr, arr.length);

        System.out.print(missingNum + " " + missingNumSum);
    }

    public static int missingNumberSum(int[] arr, int len) {
        int sum = len * (len + 1) / 2;
        int result = 0;
        for (int i = 0; i < len - 1; i++) {
            result += arr[i];
        }
        return sum - result;
    }

    public static int missingNumber(int[] arr, int len) {
        int xor1 = 0, xor2 = 0;

        for (int i = 0; i < len - 1; i++) {
            xor2 = xor2 ^ arr[i];
            xor1 = xor1 ^ (i + 1);
        }
        xor1 = xor1 ^ len;
        return xor1 ^ xor2;
    }
}
