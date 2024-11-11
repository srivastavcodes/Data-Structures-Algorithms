package Easy;

public class RemoveDuplicates {
    public static void main(String[] args) {

        int[] arr = {1, 2, 2, 3, 3, 3, 4, 4, 5, 5};
        int len = removeDuplicates(arr, arr.length);
        System.out.print(len);
    }

    private static int removeDuplicates(int[] arr, int len) {
        int left = 0;

        for (int right = 1; right < len; right++) {
            if (arr[left] != arr[right]) {
                arr[left + 1] = arr[right];
                left++;
            }
        }
        return left + 1;
    }
}
