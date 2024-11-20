package theory.medium;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class DutchNationalFlagAlgoImpl {
    public static void main(String[] args) {

        List<Integer> list = Arrays.asList(2, 2, 2, 2, 0, 0, 1, 0);
        ArrayList<Integer> arrList = new ArrayList<>(list);
        sortArray(arrList, arrList.size());
    }

    public static void sortArray(ArrayList<Integer> arrList, int size) {
        int low = 0, mid = 0, high = size - 1;

        while (mid <= high) {
            if (arrList.get(mid) == 0) swapElements(arrList, low++, mid++);
            else if (arrList.get(mid) == 1) mid += 1;
            else swapElements(arrList, mid, high--);
        }
        System.out.println(arrList);
    }

    private static void swapElements(ArrayList<Integer> arrList, int left, int right) {
        int hv = arrList.get(left);
        arrList.set(left, arrList.get(right));
        arrList.set(right, hv);
    }
}







