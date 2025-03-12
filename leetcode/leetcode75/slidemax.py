from typing import List
from collections import deque


# O(n)
# https://algo.monster/problems/sliding_window_maximum
def sliding_window_maximum(nums: List[int], k: int) -> List[int]:
    # WRITE YOUR BRILLIANT CODE HERE
    r = []
    Q = deque()  # used as a stack

    # approach is to only use one loop
    # we use the monotonic queue approach
    # until the first element in the queue is less than current, we dequeue
    # then we enqueue
    # we store only indexes in Q
    # if the first index of the Q is outside the window we remove it from the queue
    # if the index is over or equal our window we simply append the maximum to our result
    # the Q maintain the state of max element this is brillant !

    for i, num in enumerate(nums):
        while Q and num > nums[Q[-1]]:
            Q.pop()

        Q.append(i)
        # remove first element if outside the window
        if Q[0] == i - k:
            Q.popleft()

        if i >= k - 1:
            r.append(nums[Q[0]])

    return r


print(sliding_window_maximum([1, 3, 2, 5, 8, 9], 3))  # [3, 5, 8, 9]
