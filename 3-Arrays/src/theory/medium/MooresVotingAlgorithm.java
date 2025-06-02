package theory.medium;

public class MooresVotingAlgorithm {
    public static void main(String[] args) {

        int[] vals = {5, 7, 5, 5, 7, 7, 1, 5, 5, 5};
        int result = majorityElement(vals);
        System.out.print(result);
    }

    public static int majorityElement(int[] vals) {
        int count = 0, elmnt = -1, store = 0;

        for (int val : vals) {
            if (count == 0) {
                count = 1;
                elmnt = val;
            } else if (elmnt == val) {
                count++;
            } else count--;
        }
        for (int val : vals) if (val == elmnt) store++;
        if (store > vals.length / 2) return elmnt;
        return -1;
    }
}







