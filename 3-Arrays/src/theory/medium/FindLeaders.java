package theory.medium;

import java.util.ArrayList;
import java.util.List;

public class FindLeaders {
    public static void main(String[] args) {

        int[] vals = {1, 2, 3, 2};
        List<Integer> leaders = superiorElements(vals);
        System.out.print(leaders);
    }

    public static List<Integer> superiorElements(int[] vals) {
        List<Integer> leaders = new ArrayList<>();

        int max = Integer.MIN_VALUE;
        int len = vals.length - 1;

        for (int i = len; i >= 0; i--) {
            if (vals[i] > max) {
                max = vals[i];
                leaders.add(vals[i]);
            }
        }
        // If asked sorted order - Collections.sort(leaders);
        // If asked preserved order(it'll be reversed because we store the last ones first)
        // - Collections.reverse(leaders);
        return leaders;
    }
}







