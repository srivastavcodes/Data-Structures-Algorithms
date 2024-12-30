package revision.hard;

import java.util.ArrayList;
import java.util.List;

class MajorityElement_II {
    public static void main(String[] args) {

        int[] vals = {1, 1, 1, 4, 5, 6};
        List<Integer> rslt = majorityElement(vals);
        System.out.print(rslt);
    }

/*
    static List<Integer> majorityElement(int[] vals) {
        int len = vals.length;
        List<Integer> intList = new ArrayList<>();
        Map<Integer, Integer> intMap = new HashMap<>();

        for (int i = 0; i < len; i++) {
            intMap.put(vals[i], intMap.getOrDefault(vals[i], 0) + 1);

            if (intMap.get(vals[i]) > len / 3 && !intList.contains(vals[i])) {
                intList.add(vals[i]);
            }
            if (intList.size() == 2) break;
        }
        return intList;
    }
*/

    public static List<Integer> majorityElement(int[] vals) {
        int countOne = 0, countTwo = 0;
        int val1 = Integer.MIN_VALUE, val2 = Integer.MIN_VALUE;

        for (int i = 0; i < vals.length; i++) {
            if (countOne == 0 && val2 != vals[i]) {
                countOne = 1;
                val1 = vals[i];
            } else if (countTwo == 0 && val1 != vals[i]) {
                countTwo = 1;
                val2 = vals[i];
            } else if (val1 == vals[i]) {
                countOne += 1;
            } else if (val2 == vals[i]) {
                countTwo += 1;
            } else {
                countOne -= 1;
                countTwo -= 1;
            }
        }
        List<Integer> intList = new ArrayList<>();
        countOne = 0;
        countTwo = 0;
        for (int i = 0; i < vals.length; i++) {
            if (vals[i] == val1) countOne++;
            if (vals[i] == val2) countTwo++;
        }
        int min = vals.length / 3;
        if (countOne > min) intList.add(val1);
        if (countTwo > min) intList.add(val2);

        return intList;
    }
}









