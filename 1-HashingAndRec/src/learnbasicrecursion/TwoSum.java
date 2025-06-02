package learnbasicrecursion;

import java.util.Arrays;

public class TwoSum {
    public static void main(String[] args) {
        int[] nums = {11, 7, 15, 2};
        int target = 9;
        System.out.println(Arrays.toString(twoSum(nums, target)));
    }

    static int[] twoSum(int[] nums, int target) {
        return findTwoSum(nums, target, 0, 1);
    }

    private static int[] findTwoSum(int[] nums, int target, int i, int j) {
        if (i >= nums.length - 1) {
            return new int[]{-1, -1}; // No solution found
        }
        if (j >= nums.length) {
            return findTwoSum(nums, target, i + 1, i + 2);
        }
        if (nums[i] + nums[j] == target) {
            return new int[]{i, j};
        }
        return findTwoSum(nums, target, i, j + 1);
    }
}








