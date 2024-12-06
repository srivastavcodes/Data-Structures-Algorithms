package revision.medium;

import java.util.Arrays;

public class LeetCodeReArrangeBySign {
    public static void main(String[] args) {

        int[] vals = {39, -8, -28, 46, -16, -21, -12, 49, 14, -46, 22, 9};
        int[] res = rearrangeArray(vals);
        System.out.println(Arrays.toString(res));
    }

    public static int[] rearrangeArray(int[] vals) {
        int[] res = new int[vals.length];
        int posIndex = 0, negIndex = 1;

        for (int i = 0; i < vals.length; i++) {
            if (vals[i] < 0) {
                res[negIndex] = vals[i];
                negIndex += 2;
            } else {
                res[posIndex] = vals[i];
                posIndex += 2;
            }
        }
        return res;
    }
}
