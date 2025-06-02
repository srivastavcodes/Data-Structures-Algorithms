package revision.medium;

class LeetCodeMajorityElement {
    public static void main(String[] args) {

        int[] vals = {3, 2, 3};
        int res = majorityElement(vals);
        System.out.print(res);
    }

    private static int majorityElement(int[] vals) {
        int repetition = 0, reqElement = 0;
        for (int i = 0; i < vals.length; i++) {
            if (repetition == 0) {
                repetition = 1;
                reqElement = vals[i];
            } else if (vals[i] == reqElement) {
                repetition += 1;
            } else {
                repetition -= 1;
            }
        }
        return reqElement;
    }
}







