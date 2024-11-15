package medium;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class SellAndBuyStocks {
    public static void main(String[] args) {

        List<Integer> list = Arrays.asList(7, 1, 5, 3, 6, 4);
        ArrayList<Integer> vals = new ArrayList<>(list);
        int maxProfit = maximumProfit(vals);
        System.out.print(maxProfit);
    }

    public static int maximumProfit(ArrayList<Integer> vals) {
        int maxProfit = 0, mini = vals.get(0);

        for (int i = 1; i < vals.size(); i++) {
            int cost = vals.get(i) - mini;
            maxProfit = Math.max(maxProfit, cost);
            mini = Math.min(vals.get(i), mini);
        }
        return maxProfit;
    }
}







