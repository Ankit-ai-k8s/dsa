class Solution {
    public int subarraysWithKDistinct(int[] nums, int k) {
        return solve(nums, k) - solve(nums, k - 1);
    }

    int solve(int[] nums, int k) {
        int n = nums.length;
        int start = 0;
        int end = 0;
        int res = 0;
        Map<Integer, Integer> map = new HashMap<>();
        while (end < n) {
            int ed = nums[end];
            map.put(ed, map.getOrDefault(ed, 0) + 1);

            while (start < n && map.size() > k) {
                int es = nums[start];
                map.put(es, map.getOrDefault(es, 0) - 1);
                if (map.get(es) == 0)
                    map.remove(es);
                start++;
            }
            res += end - start + 1;
            end++;
        }

        return res;
    }
}