package theory.medium;

import java.util.ArrayList;
import java.util.Arrays;

public class RearrangeUnequalArrayElements {
    public static void main(String[] args) {

        int[] vals = {3, 1, -2, -5, 3, -2};
        int[] result = alternateUnequalArray(vals);
        System.out.print(Arrays.toString(result));
    }

    public static int[] alternateUnequalArray(int[] vals) {
        ArrayList<Integer> posArr = new ArrayList<>();
        ArrayList<Integer> negArr = new ArrayList<>();

        for (int val : vals) {
            if (val > 0) posArr.add(val);
            else negArr.add(val);
        }

        if (posArr.size() > negArr.size()) {
            for (int i = 0; i < negArr.size(); i++) {
                vals[2 * i] = posArr.get(i);
                vals[2 * i + 1] = negArr.get(i);
            }
            int index = negArr.size() * 2;
            for (int i = negArr.size(); i < posArr.size(); i++) {
                vals[index++] = posArr.get(i);
            }
        } else {
            for (int i = 0; i < posArr.size(); i++) {
                vals[2 * i] = posArr.get(i);
                vals[2 * i + 1] = negArr.get(i);
            }
            int index = posArr.size() * 2;
            for (int i = posArr.size(); i < negArr.size(); i++) {
                vals[index++] = negArr.get(i);
            }
        }
        return vals;
    }
}
