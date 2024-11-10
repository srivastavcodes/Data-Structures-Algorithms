package Easy;

public class FindLargestElement {
    public static void main(String[] args) {

        int[] arr = {4, 7, 8, 6, 7, 6};
        int largestElement = getLargest(arr, arr.length - 1);
        System.out.print(largestElement);
    }

    private static int getLargest(int[] arr, int len) {
        int largest = arr[0];

        for (int i = 0; i < len; i++) {
            if (largest < arr[i]) {
                largest = arr[i];
            }
        }
        return largest;
    }
}
