package Easy;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class ArrayIntersection {
    public static void main(String[] args) {

        ArrayList<Integer> arr = new ArrayList<>(Arrays.asList(1, 2, 3, 4, 5));
        ArrayList<Integer> secArr = new ArrayList<>(Arrays.asList(2, 3, 5, 7, 9));

        List<Integer> arrayIntersection = findArrayIntersection(
          arr, arr.size(), secArr, secArr.size());

        System.out.println(arrayIntersection);
    }

    public static ArrayList<Integer> findArrayIntersection(
      ArrayList<Integer> arr, int arrLen, ArrayList<Integer> secArr, int secArrLen
    ) {
        ArrayList<Integer> result = new ArrayList<>();
        int i = 0;
        int j = 0;

        while (i < arrLen && j < secArrLen) {
            if (arr.get(i) < secArr.get(j)) i++;
            else if (arr.get(i) > secArr.get(j)) j++;
            else {
                result.add(arr.get(i));
                i += 1;
                j += 1;
            }
        }
        return result;
    }
}







