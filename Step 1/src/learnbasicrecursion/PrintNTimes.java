package learnbasicrecursion;

public class PrintNTimes {
    public static void main(String[] args) {
        System.out.println();

        int count = 0;
        printNumber(count);
    }

    public static void printNumber(int count) {
        System.out.println(count);
        if (count == 150) {
            return;
        }
        count++;
        printNumber(count);
    }
}
