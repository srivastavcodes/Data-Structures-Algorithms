package theory.hard;

import java.util.ArrayList;
import java.util.List;

class PascalTriangle {
    public static void main(String[] args) {

        int r = 6;
        List<List<Integer>> res = generate(r);
        System.out.println(res);
    }

    @SuppressWarnings("unused")
    private static int getElement(int r, int c) {
        int res = 1;
        for (int i = 0; i < c - 1; i++) {
            res = res * ((r - 1) - i);
            res = res / (i + 1);
        }
        return res;
    }

    private static List<Integer> generateRow(int r) {
        List<Integer> resList = new ArrayList<>();
        int res = 1;
        resList.add(res);
        for (int i = 1; i < r; i++) {
            res = res * (r - i);
            res = res / i;
            resList.add(res);
        }
        return resList;
    }

    public static List<List<Integer>> generate(int rows) {
        List<List<Integer>> resList = new ArrayList<>();
        for (int i = 1; i <= rows; i++) {
            resList.add(generateRow(i));
        }
        return resList;
    }
}









