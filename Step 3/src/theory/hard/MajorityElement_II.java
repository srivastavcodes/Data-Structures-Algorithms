package theory.hard;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

class MajorityElement_II {
    public static void main(String[] args) {

        int[] vals = {2, 2, 0, 1, 2, 1};
        List<Integer> result = majorityElement(vals);
        System.out.print(result);
    }

//    public static List<Integer> majorityElement(int[] vals) {
//        Map<Integer, Integer> intMap = new HashMap<>();
//        List<Integer> intList = new ArrayList<>();
//
//        int len = vals.length, min = len / 3 + 1;
//
//        for (int i = 0; i < len; i++) {
//            intMap.put(vals[i], intMap.getOrDefault(vals[i], 0) + 1);
//
//            if (intMap.get(vals[i]) == min) intList.add(vals[i]);
//            if (intList.size() == 2) break;
//        }
//        return intList;
//    }

    public static List<Integer> majorityElement(int[] vals) {
        int count1 = 0, count2 = 0;
        int elmnt1 = Integer.MIN_VALUE, elmnt2 = Integer.MIN_VALUE;

        for (int i = 0; i < vals.length; i++) {
            if (count1 == 0 && elmnt2 != vals[i]) {
                count1 = 1;
                elmnt1 = vals[i];
            } else if (count2 == 0 && elmnt1 != vals[i]) {
                count2 = 1;
                elmnt2 = vals[i];
            } else if (vals[i] == elmnt1) {
                count1++;
            } else if (vals[i] == elmnt2) {
                count2++;
            } else {
                count1 -= 1;
                count2 -= 1;
            }
        }

        List<Integer> intList = new ArrayList<>();
        count1 = 0;
        count2 = 0;

        for (int i = 0; i < vals.length; i++) {
            if (vals[i] == elmnt1) count1++;
            if (vals[i] == elmnt2) count2++;
        }
        int mini = (vals.length / 3) + 1;
        if (count1 >= mini) intList.add(elmnt1);
        if (count2 >= mini) intList.add(elmnt2);

        Collections.sort(intList);
        return intList;
    }
}









