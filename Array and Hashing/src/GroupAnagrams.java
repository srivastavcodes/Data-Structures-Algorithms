import java.util.*;

public class GroupAnagrams {
	public static void main(String[] args) {
		String[] str = {"eat", "tea", "tan", "ate", "nat", "bat"};
		List<List<String>> result = groupAnagrams(str);
		System.out.println(result);
	}

	public static List<List<String>> groupAnagrams(String[] args) {
		Map<String, List<String>> result = new HashMap<>();
		for (String str : args) {
			int[] count = new int[26];
			for (char c : str.toCharArray()) {
				count[c - 'a']++;
			}
			String key = Arrays.toString(count);
			result.putIfAbsent(key, new ArrayList<>());
			result.get(key).add(str);
		}
		return new ArrayList<>(result.values());
	}
}
