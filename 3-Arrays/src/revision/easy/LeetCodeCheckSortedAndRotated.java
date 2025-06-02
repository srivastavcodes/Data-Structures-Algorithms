package revision.easy;

public class LeetCodeCheckSortedAndRotated {
    public static void main(String[] args) {

        int[] vals = {2, 1, 3, 4};
        boolean status = checkIfSortedRotated(vals);
        System.out.print(status);
    }

    private static boolean checkIfSortedRotated(int[] vals) {
        int len = vals.length, valDrop = 0;

        for (int i = 1; i < len; i++) {
            if (vals[i] < vals[i - 1]) valDrop++;
        }
        if (valDrop == 0) return true;
        return valDrop == 1 && vals[0] >= vals[len - 1];
    }
}







