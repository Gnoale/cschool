from dataclasses import dataclass


def find_missing_and_repeated_numbers(grid):
    numbers = [elem for pair in grid for elem in pair]

    numbers.sort()

    duplicat, missing = 0, 0
    previous = numbers[0]
    i = 1

    while i < len(numbers):
        if numbers[i] == previous:
            duplicat = numbers[i]
        if numbers[i] - previous > 1:
            missing = previous + 1
        previous = numbers[i]
        i += 1

    if missing == 0:
        if numbers[0] != 1:
            missing = 1
        else:
            missing = numbers[-1] + 1

    return [duplicat, missing]


def custom_sort_string(order: str, s: str) -> str:

    word = [letter for letter in s]

    word.sort(key=lambda x: order.index(x) if x in order else len(order))

    return "".join(word)


def custom_sort_string_frequency(order: str, s: str) -> str:

    frequency = {letter: s.count(letter) for letter in s}
    ans = []
    for letter in order:
        if letter in frequency:
            ans.extend([letter] * frequency[letter])

    for letter in s:
        if letter not in order:
            ans.append(letter)

    return "".join(ans)


def binary_search(nums: list[int], target: int) -> int:
    """
        1 3 4 7 7 9 10

    where to insert 8
    left right mid nums[mid]
    0     6    3   4
    4     6    5   9
    4     4    4   7
    return 4

    where to insert 1
    left right mid nums[mid]
    0     6    3   4
    0     2    1   3
    0     0    0   1
    return 0

    where to insert 5
    left right mid nums[mid]
    0     6    3   4
    4     6    5   9
    4     4    4   7
    return 4

    """
    left = 0
    right = len(nums) - 1

    while left <= right:
        mid = (left + right) // 2
        if nums[mid] < target:
            left = mid + 1
        elif nums[mid] > target:
            right = mid - 1
        else:
            return mid

    return left


def median_sliding_window(nums: list[int], k: int) -> list[float]:

    ans = []

    left = 0
    right = k
    queue = nums[left:right]

    while left <= len(nums) - k:

        if left == 0:
            queue.sort()
        else:
            to_dequeue = nums[left - 1]
            queue.remove(to_dequeue)
            candidate = nums[right - 1]
            index_to_insert = binary_search(queue, candidate)
            if index_to_insert == len(queue):
                queue.append(candidate)
            elif candidate <= queue[index_to_insert]:
                queue.insert(index_to_insert, candidate)
            else:
                queue.insert(index_to_insert + 1, candidate)
        print(queue)

        median_index = k // 2
        if k % 2 != 0:
            ans.append(float(queue[median_index]))
        else:
            ans.append((queue[median_index] + queue[median_index - 1]) / 2)

        left += 1
        right += 1

    return ans


def is_prime(n: int) -> bool:
    if n <= 1:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False

    # on démarre à 3 car 2 est premier
    # on incrémente de 2 car on ne veut pas tester les nombres pairs
    # on s'arrête à la racine carrée de n
    for i in range(3, int(n**0.5) + 1, 2):
        if n % i == 0:
            return False
    return True


def closest_prime_number(left: int, right: int) -> list[int]:

    primes = []
    for n in range(left, right + 1):
        if is_prime(n):
            primes.append(n)

    if not primes:
        return [-1, -1]
    print(primes)

    candidates = [primes[0], primes[1]]
    i = 1

    while i < len(primes) - 1:
        if primes[i + 1] - primes[i] < candidates[1] - candidates[0]:
            candidates = [primes[i], primes[i + 1]]
        i += 1

    return candidates


def alternate_colors(colors: list[int], k: int) -> int:
    # unwrap the array
    i = 0
    while i <= k - 2:
        colors.append(colors[i])
        i += 1

    left, right = 0, 1
    ans = 0

    while right < len(colors):

        # if the current color is the same as the previous one, slide the window
        if colors[right] == colors[right - 1]:
            left = right
            right += 1

        # the window is of the correct size slide the window
        # and increment the answer
        elif right - left + 1 == k:
            print(left, right, colors[left : right + 1])
            ans += 1
            right += 1
            left += 1
        else:
            right += 1

    return ans


def find_provinces(adjacency: list[list[int]]) -> int:

    visited = [False] * len(adjacency)

    def visit_nodes(row: int, adjacency: list[list[int]], visited: list[bool]):
        visited[row] = True
        print(row, visited)

        for col in range(len(adjacency[row])):
            if adjacency[row][col] == 1 and not visited[col]:
                visit_nodes(col, adjacency, visited)

    groups = 0
    row = 0
    while row < len(adjacency):
        if not visited[row]:
            print("new group")
            visit_nodes(row, adjacency, visited)
            groups += 1
        row += 1

    return groups


def _has_all_chars(freq: list) -> bool:
    # Check if we have at least one of each character
    return all(f > 0 for f in freq)


def number_of_substrings(s: str) -> int:
    left = right = total = 0
    # Track frequency of a, b, c
    freq = [0] * 3

    while right < len(s):
        # Add character at right pointer to frequency array
        freq[ord(s[right]) - ord("a")] += 1

        # While we have all required characters
        while _has_all_chars(freq):
            # All substrings from current window to end are valid
            # Add count of valid substrings
            total += len(s) - right

            # Remove leftmost character and move left pointer
            freq[ord(s[left]) - ord("a")] -= 1
            left += 1
        print(freq, right, left)
        right += 1

    return total
