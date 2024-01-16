def maxOperations(nums: list[int], k: int) -> int:
    # Très intuitif une fois la list sorted..
    #

    P = 0
    nums.sort()
    i = 0
    j = len(nums) - 1

    while i < j:
        if nums[i] + nums[j] == k:
            P += 1
            i += 1
            j -= 1
        elif nums[i] + nums[j] > k:
            j -= 1
        else:
            i += 1

    return P


print(maxOperations([3, 1, 3, 4, 3], 6))  # 1
print(maxOperations([1, 2, 3, 4], 5))  # 2
print(maxOperations([1, 2, 3, 4], 2))  # 0

print(maxOperations([4, 4, 1, 3, 1, 3, 2, 2, 5, 5, 1, 5, 2, 1, 2, 3, 5, 4], 2))  # 2

print(maxOperations([2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2], 3))
