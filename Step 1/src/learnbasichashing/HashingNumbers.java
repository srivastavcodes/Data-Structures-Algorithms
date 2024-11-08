package learnbasichashing;

public class HashingNumbers {
    public static void main(String[] args) {

        int[] arr = {1, 2, 1, 5, 12, 5, 5, 2};
        int digit = 1;
        int maxLength = 20;
        preHash(arr, digit, maxLength);
    }

    static void preHash(int[] arr, int digit, int maxLen) {
        int[] hashArr = new int[(int) 1e8 + 1];
        for (int i : arr) {
            hashArr[i] += 1;
        }
        System.out.print(hashArr[digit]);
    }
}
