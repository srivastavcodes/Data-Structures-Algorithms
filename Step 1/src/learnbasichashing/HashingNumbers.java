package learnbasichashing;

public class HashingNumbers {
    public static void main(String[] args) {

        int[] arr = {1, 2, 1, 5, 12, 5, 5, 2};
        int digit = 1;
        System.out.print(preHash(arr, digit));
    }

    static int preHash(int[] arr, int digit) {
        int[] hashArr = new int[(int) 1e3];
        for (int i : arr) {
            hashArr[i]++;
        }
        return hashArr[digit];
    }
}
