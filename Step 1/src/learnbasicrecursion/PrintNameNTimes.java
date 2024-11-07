package learnbasicrecursion;

import java.util.Scanner;

public class PrintNameNTimes {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        System.out.println("Enter the amount::");
        int times = sc.nextInt();

        printNames(times);
    }

    // This is a backtracking example.
    private static void printNames(int times) {
        if (times < 1) {
            return;
        }
        printNames(times - 1);
        System.out.print("GFG" + " ");
    }
}








