package revision.medium;

import java.util.ArrayList;

public class GeeksForGeeksFindUnion {
    public static void main(String[] args) {

        int[] vals1 = {-7, 8};
        int[] vals2 = {-8, -3, 8};
        ArrayList<Integer> union = findUnion(vals1, vals2);
        System.out.print(union);
    }

    private static ArrayList<Integer> findUnion(int[] vals1, int[] vals2) {
        ArrayList<Integer> resList = new ArrayList<>();
        int i = 0, j = 0;

        while (i < vals1.length && j < vals2.length) {
            int len = resList.size();
            if (vals1[i] <= vals2[j]) {
                if (len == 0 || resList.get(len - 1) != vals1[i]) {
                    resList.add(vals1[i]);
                }
                i += 1;
            } else {
                if (len == 0 || resList.get(len - 1) != vals2[j]) {
                    resList.add(vals2[j]);
                }
                j += 1;
            }
        }
        while (i < vals1.length) {
            int len = resList.size();
            if (len == 0 || resList.get(len - 1) != vals1[i]) {
                resList.add(vals1[i]);
            }
            i += 1;
        }
        while (j < vals2.length) {
            int len = resList.size();
            if (len == 0 || resList.get(len - 1) != vals2[j]) {
                resList.add(vals2[j]);
            }
            j += 1;
        }
        return resList;
    }
}







