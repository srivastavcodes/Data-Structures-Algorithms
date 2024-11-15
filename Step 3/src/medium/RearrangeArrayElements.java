package medium;

import java.util.Arrays;

public class RearrangeArrayElements {
    public static void main(String[] args) {

        int[] vals = {3, 1, -2, -5, 3, -2};
        int[] result = alternateNumbers(vals);
        System.out.print(Arrays.toString(result));
    }

    public static int[] alternateNumbers(int[] vals) {
        int[] result = new int[vals.length];
        int posIndex = 0, negIndex = 1;

        for (int val : vals) {
            if (val > 0) {
                result[posIndex] = val;
                posIndex += 2;
            } else {
                result[negIndex] = val;
                negIndex += 2;
            }
        }
        return result;
    }
}







