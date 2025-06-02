package learnbasicrecursion;

public class PrintNTimes {
    public static void main(String[] args) {
        int count = 10;
        printNumber(count);
    }

    public static void printNumber(int count) {
        if (count == 0) {
            return;
        }
        printNumber(count - 1);
        System.out.print(count + " ");
    }
}
