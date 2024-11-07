package learnbasicrecursion;

import java.util.Scanner;

public class PrintNameNTimes {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        System.out.println("Enter your Name::");
        String name = sc.nextLine();

        System.out.println("Enter the amount::");
        int times = sc.nextInt();

        printNames(name, times);
    }

    // This is a backtracking example.
    private static void printNames(String name, int times) {
        if (times < 1) {
            return;
        }
        printNames(name, times - 1);
        System.out.println(name + times);
    }
}








