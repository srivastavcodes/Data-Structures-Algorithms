package theory.medium;

import java.util.HashMap;
import java.util.Map;

public class TotalSubArrayWithSumK {
    public static void main(String[] args) {

        int[] vals = {3, 1, 2, 4};
        int k = 6;
        int result = getSubarraysWithSumK(vals, k);
        System.out.print(result);
    }

    public static int getSubarraysWithSumK(int[] vals, int k) {
        Map<Integer, Integer> integerMap = new HashMap<>();
        integerMap.put(0, 1);
        int prefixSum = 0, count = 0;

        for (int i = 0; i < vals.length; i++) {
            prefixSum += vals[i];
            int left = prefixSum - k;
            if (integerMap.get(left) != null) count += integerMap.get(left);
            integerMap.put(prefixSum, integerMap.getOrDefault(prefixSum, 0) + 1);
        }
        return count;
    }
}







