package lecture2;

import java.util.ArrayList;
import java.util.Arrays;

public class AllocateBooks {
    public static void main(String[] args) {
        ArrayList<Integer> books = new ArrayList<>(Arrays.asList(12, 34, 67, 90));
        int len = books.size();
        int students = 2;
        int output = findPages(books, len, students);
        System.out.print(output);
    }

    public static int findPages(ArrayList<Integer> books, int len, int students) {
        if (students > len) return -1;
        int low = Integer.MIN_VALUE, high = 0;

        for (int i = 0; i < books.size() - 1; i++) {
            low = Math.max(low, books.get(i));
            high += books.get(i);
        }
        while (low <= high) {
            int center = (low + high) / 2;
            if (valid(books, center, students)) {
                high = center - 1;
            } else {
                low = center + 1;
            }
        }
        return low;
    }

    private static boolean valid(ArrayList<Integer> books, int center, int students) {
        int count = 1, pages = 0;
        for (Integer book : books) {
            if (book + pages <= center) {
                pages += book;
            } else {
                count++;
                pages = book;
            }
        }
        return count <= students;
    }
}