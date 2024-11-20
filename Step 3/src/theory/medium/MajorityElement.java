package theory.medium;

import java.util.HashMap;
import java.util.Map;

public class MajorityElement {
    public static void main(String[] args) {

        int[] vals = {2, 2, 1, 3, 3, 2, 2};
        int majorityE = majorityElement(vals);
        System.out.print(majorityE);
    }

    public static int majorityElement(int[] vals) {
        HashMap<Integer, Integer> intMap = new HashMap<>();
        for (int i : vals) {
            intMap.put(i, intMap.getOrDefault(i, 0) + 1);
        }
        for (Map.Entry<Integer, Integer> val : intMap.entrySet()) {
            if (val.getValue() > intMap.size() / 2) {
                return val.getKey();
            }
        }
        return -1;
    }
}
