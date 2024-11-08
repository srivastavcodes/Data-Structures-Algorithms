package learnbasichashing;

import java.util.HashMap;
import java.util.Map;

public class HashMapStringHashing {
    public static void main(String[] args) {

        String str = "Parth Srivastav";
        String ch = "Parth";

        preHashString(str, ch);
    }

    private static void preHashString(String str, String ch) {
        Map<String, Integer> hashArr = new HashMap<>();

        for (int i = 0; i < str.length(); i++) {
            String key = str.substring(i, i + 1);
            hashArr.put(key, hashArr.getOrDefault(key, 0) + 1);
        }

        for (int i = 0; i < ch.length(); i++) {
            String key = ch.substring(i, i + 1);
            System.out.println(hashArr.get(key) + " ");
        }
    }
}







