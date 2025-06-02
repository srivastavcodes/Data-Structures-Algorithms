package revision.easy;

public class LeetCodeFindMissingNumber {
    public static void main(String[] args) {

        int[] vals = {9, 6, 4, 2, 3, 5, 7, 0, 1};
        int res = missingNumber(vals);
        System.out.print(res);
    }

    private static int missingNumber(int[] vals) {
        int len = vals.length, totalSum = len * (len + 1) / 2;
        int valSum = 0;
        for (int i = 0; i < len; i++) {
            valSum = valSum + vals[i];
        }
        return totalSum - valSum;
    }
}
