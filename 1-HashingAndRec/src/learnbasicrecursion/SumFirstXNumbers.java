package learnbasicrecursion;

public class SumFirstXNumbers {
    public static void main(String[] args) {
        // If given 3 the sum would be 6

        System.out.println(sum(7));
    }

    // method 1
    public static long sum(long n) {
        if (n < 1) {
            return n;
        }
        return n*n*n + sum(n - 1);
    }

    // method 2
    @SuppressWarnings("unused")
    public static int sum(int n, int s) {
        if (n < 1) {
            return s;
        }
        return sum(n - 1, s + n);
    }
}







