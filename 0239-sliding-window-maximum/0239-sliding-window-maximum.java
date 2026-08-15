class Solution {
    public int[] maxSlidingWindow(int[] nums, int k) {

        int n = nums.length;
        // storing index for check len and max value.
        Deque<Integer> q = new ArrayDeque<>();

        int[] res = new int[n - k + 1];

    
        int idx = 0;

        for (int end = 0; end < n; end++) {

            int nm = nums[end];

            while (!q.isEmpty() && end - k >= q.peekFirst()) {
                q.pop();
            }

            while (!q.isEmpty() && nums[q.peekLast()] < nm) {
                q.removeLast();
            }

            q.add(end);

            if (end >= k - 1)
                res[idx++] = nums[q.peekFirst()];


        }

        return res;
    }
}