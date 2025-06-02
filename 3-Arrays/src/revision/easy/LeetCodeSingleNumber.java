package revision.easy;

import java.util.HashMap;
import java.util.Map;

public class LeetCodeSingleNumber {
    public static void main(String[] args) {

        int[] vals = {4, 1, 2, 1, 2};
        int sin = singleNumber(vals);
        System.out.print(sin);
    }

    private static int singleNumber(int[] vals) {
        Map<Integer, Integer> intMap = new HashMap<>();
        for (int i = 0; i < vals.length; i++) {
            intMap.put(vals[i], intMap.getOrDefault(vals[i], 0) + 1);
        }
        for (Map.Entry<Integer, Integer> entry : intMap.entrySet()) {
            if (entry.getValue() == 1) {
                return entry.getKey();
            }
        }
        return -1;
    }
}







