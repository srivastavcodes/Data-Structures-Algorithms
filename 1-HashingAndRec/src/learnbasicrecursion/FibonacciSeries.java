package learnbasicrecursion;

public class FibonacciSeries {
    public static void main(String[] args) {

        System.out.println(getFib(5));
    }

    private static int getFib(int n) {
        if (n <= 1) {
            return n;
        }
        return getFib(n - 1) + getFib(n - 2);
    }
}
