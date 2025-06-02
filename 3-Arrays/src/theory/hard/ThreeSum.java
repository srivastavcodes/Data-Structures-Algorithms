package theory.hard;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

class ThreeSum {
    public static void main(String[] args) {

        int[] vals = {-1, -1, 2, 0, 1};
        int len = vals.length;
        List<List<Integer>> result = triplet(len, vals);
        System.out.print(result);
    }

/*    public static List<List<Integer>> triplet(int len, int[] vals) {
        Set<List<Integer>> valueSet = new HashSet<>();
        for (int i = 0; i < len; i++) {

            Set<Integer> intSet = new HashSet<>();
            for (int j = i + 1; j < len; j++) {
                int k = -(vals[i] + vals[j]);

                if (intSet.contains(k)) {
                    List<Integer> vallist = Arrays.asList(vals[i], vals[j], k);
                    Collections.sort(vallist);
                    valueSet.add(vallist);
                }
                intSet.add(vals[j]);
            }
        }
        return valueSet.stream().toList();
    }*/

    public static List<List<Integer>> triplet(int len, int[] vals) {
        List<List<Integer>> reslist = new ArrayList<>();

        List<Integer> val = new ArrayList<>(Arrays.stream(vals).boxed().toList());
        Collections.sort(val);

        for (int i = 0; i < len; i++) {
            if (i > 0 && val.get(i).equals(val.get(i - 1))) continue;

            int j = i + 1, k = len - 1;
            while (j < k) {
                int sum = val.get(i) + val.get(j) + val.get(k);

                if (sum < 0) j += 1;
                else if (sum > 0) {
                    k -= 1;
                } else {
                    List<Integer> rslt = List.of(val.get(i), val.get(j), val.get(k));
                    reslist.add(rslt);
                    j += 1;
                    k -= 1;
                    while (j < k && val.get(j).equals(val.get(j - 1))) j++;
                    while (j < k && val.get(k).equals(val.get(k + 1))) k--;
                }
            }
        }
        return reslist;
    }
}









