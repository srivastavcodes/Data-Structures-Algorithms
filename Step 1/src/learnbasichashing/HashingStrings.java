package learnbasichashing;

public class HashingStrings {
    public static void main(String[] args) {

        String str = "Parth Srivastav";
        String s = "s";
        int maxLen = 256;

        preHashString(str, s, maxLen);
    }

    // seems like can be used for searching.
    static void preHashString(String str, String ch, int maxLen) {
        int[] hashArr = new int[maxLen];
        for (int i = 0; i < str.length(); i++) {
            hashArr[str.charAt(i)] += 1;
        }
        for (int i = 0; i < ch.length(); i++) {
            System.out.print(hashArr[ch.charAt(i)] + " ");
        }
    }
}








