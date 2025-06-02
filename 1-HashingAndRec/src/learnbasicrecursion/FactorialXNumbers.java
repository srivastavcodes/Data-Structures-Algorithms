package learnbasicrecursion;

public class FactorialXNumbers {
    public static void main(String[] args) {
        // If given 3 the sum would be 6

        System.out.println(fact(5));
    }

    // method 1
    public static int fact(int n) {
        if (n < 1) {
            return 1;
        }
        return n * fact(n - 1);
    }
}
