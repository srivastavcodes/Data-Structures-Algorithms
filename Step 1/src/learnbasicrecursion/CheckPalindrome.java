package learnbasicrecursion;

public class CheckPalindrome {
    public static void main(String[] args) {

        String str = "madam";
        System.out.print(isPalindrome(str));
    }

    private static boolean isPalindrome(String str) {
        return verifyPalindrome(str, 0, str.length());
    }

    private static boolean verifyPalindrome(String str, int i, int len) {
        if (i >= len / 2) {
            return true;
        }
        if (str.charAt(i) != str.charAt(len - i - 1)) {
            return false;
        }
        return verifyPalindrome(str, i + 1, len);
    }
}
