package revision.medium;

import java.util.ArrayList;

public class GFGFindLeaders {
    public static void main(String[] args) {

        int[] vals = {16, 17, 4, 3, 5, 2};
        ArrayList<Integer> leaders = leaders(vals);
        System.out.println(leaders);
    }

    public static ArrayList<Integer> leaders(int[] vals) {
        ArrayList<Integer> leaders = new ArrayList<>();
        int largest = Integer.MIN_VALUE, len = vals.length - 1;

        for (int i = len; i >= 0; i--) {
            if (vals[i] >= largest) {
                leaders.add(0, vals[i]);
                largest = vals[i];
            }
        }
        return leaders;
    }
}
