package revision.medium;

import java.util.HashMap;
import java.util.Map;

class LeetCodeSubArrSumEqualsK {
    public static void main(String[] args) {

        int[] vals = {1, 2, 3};
        int k = 2, res = subarraySum(vals, k);
        System.out.print(res);
    }

    public static int subarraySum(int[] vals, int k) {
        Map<Integer, Integer> integerMap = new HashMap<>();
        int cont = 0, preSum = 0;

        integerMap.put(0, 1);
        for (int i = 0; i < vals.length; i++) {
            preSum += vals[i];
            int remove = preSum - k;
            cont += integerMap.getOrDefault(remove, 0);
            integerMap.put(preSum, integerMap.getOrDefault(preSum, 0) + 1);
        }
        return cont;
    }
}









