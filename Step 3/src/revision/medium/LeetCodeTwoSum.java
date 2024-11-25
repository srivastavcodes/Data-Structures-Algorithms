package revision.medium;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class LeetCodeTwoSum {
    public static void main(String[] args) {

        int[] vals = {0, 1, 3, 11, 5, 8};
        int k = 9;
        int[] indices = getIndices(vals, k);
        System.out.print(Arrays.toString(indices));
    }

    private static int[] getIndices(int[] vals, int k) {
        Map<Integer, Integer> indexMap = new HashMap<>();

        for (int i = 0; i < vals.length; i++) {
            if (indexMap.containsKey(k - vals[i])) {
                return new int[]{indexMap.get(k - vals[i]), i};
            }
            indexMap.put(vals[i], i);
        }
        return new int[]{};
    }
}







