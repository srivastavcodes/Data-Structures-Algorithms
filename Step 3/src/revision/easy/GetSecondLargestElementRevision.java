package revision.easy;

public class GetSecondLargestElementRevision {
    public static void main(String[] args) {

        int[] vals = {12, 35, 1, 10, 34, 1};
        int secLargest = getSecondLargest(vals);
        System.out.print(secLargest);
    }

    private static int getSecondLargest(int[] vals) {
        int secLargest = 0, largest = vals[0];

        for (int i = 0; i < vals.length; i++) {
            if (vals[i] > largest) {
                secLargest = largest;
                largest = vals[i];
            } else if (vals[i] < largest && vals[i] > secLargest) {
                secLargest = vals[i];
            }
        }
        return secLargest;
    }
}
