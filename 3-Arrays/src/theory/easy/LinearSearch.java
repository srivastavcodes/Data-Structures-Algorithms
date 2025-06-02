package theory.easy;

public class LinearSearch {
    public static void main(String[] args) {

        int[] arr = {4, 1, 54, 3, 9, 4};
        int n = 54;
        int got = linearSearch(arr, n, arr.length);
        System.out.println(got);
    }

    private static int linearSearch(int[] arr, int num, int len) {
        for (int i = 0; i < len; i++) {
            if (arr[i] == num) {
                return i;
            }
        }
        return -1;
    }
}
