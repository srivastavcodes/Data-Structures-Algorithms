package revision.easy;

public class LeetCodeRemoveDuplicates {
    public static void main(String[] args) {

        int[] vals = {0, 0, 1, 1, 1, 2, 2, 3, 3, 4};
        int result = removeDuplicates(vals);
        System.out.print(result);
    }

    private static int removeDuplicates(int[] vals) {
        int i = 0;
        for (int j = 0; j < vals.length; j++) {
            if (vals[i] != vals[j]) {
                vals[++i] = vals[j];
            }
        }
        return i + 1;
    }
}







