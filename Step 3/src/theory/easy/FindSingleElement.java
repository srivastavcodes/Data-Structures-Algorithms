package theory.easy;

public class FindSingleElement {
    public static void main(String[] args) {

        int[] arr = {1, 1, 2, 3, 3, 4, 4};
        int res = getSingleElement(arr);
        System.out.print(res);
    }

    public static int getSingleElement(int[] arr) {
        int xor = 0;
        for (int i : arr) xor = xor ^ i;
        return xor;
    }
}
