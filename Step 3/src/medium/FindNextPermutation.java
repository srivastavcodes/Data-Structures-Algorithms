package medium;

import java.util.Arrays;
import java.util.Collections;
import java.util.List;

public class FindNextPermutation {
    public static void main(String[] args) {

        List<Integer> vals = Arrays.asList(1, 3, 2);
        List<Integer> result = nextGreaterPermutation(vals);
        System.out.print(result);
    }

    public static List<Integer> nextGreaterPermutation(List<Integer> vals) {
        int index = -1, size = vals.size();

        for (int i = size - 2; i >= 0; i--) {
            if (vals.get(i) < vals.get(i + 1)) {
                index = i;
                break;
            }
        }
        if (index == -1) {
            Collections.reverse(vals);
            return vals;
        }
        for (int i = size - 1; i > index; i--) {
            if (vals.get(i) > vals.get(index)) {

                int swap = vals.get(i);
                vals.set(i, vals.get(index));
                vals.set(index, swap);
                break;
            }
        }
        List<Integer> sublist = vals.subList(index + 1, size);
        Collections.reverse(sublist);
        return vals;
    }
}







