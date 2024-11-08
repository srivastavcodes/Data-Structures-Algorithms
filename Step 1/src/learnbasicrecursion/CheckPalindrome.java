package learnbasicrecursion;

public class CheckPalindrome {
    public static void main(String[] args) {
        // Valid Palindrome (LeetCode)

        String str = "0P";
        System.out.print(isPalindrome(str));
    }

    private static boolean isPalindrome(String str) {
        String regStr = str.replaceAll("[^a-zA-Z0-9]", "");
        return verifyPalindrome(regStr.toLowerCase(), 0, regStr.length());
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
