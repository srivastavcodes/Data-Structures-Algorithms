package learnbasicrecursion;

import java.util.ArrayList;
import java.util.List;

public class FindAllFactorial {
    public static void main(String[] args) {
        // Find all factorial numbers less than or equal to n. (GFG)
        // First few factorial numbers are 1, 2, 6, 24, 120
        // 1, 1*2 = 2, 2*3 = 6, 6*4 = 24, 24*5 = 120
        // 6 is given. output = 1, 2, 6

        long num = 100L;
        List<Long> factList = factorialNumbers(num);
        System.out.println(factList);
    }

    static ArrayList<Long> factorialNumbers(long n) {
        ArrayList<Long> factorials = new ArrayList<>();
        getAllFactorials(n, 1L, 1L, factorials);
        return factorials;
    }

    static void getAllFactorials(long n, Long i, Long currentFact, ArrayList<Long> factorials) {
        if (currentFact > n) {
            return;
        }
        factorials.add(currentFact);
        i++;
        getAllFactorials(n, i, currentFact * i, factorials);
    }
}








