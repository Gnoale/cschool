package leetcode75

def canPlaceFlowers(flowerbed: list[int], n: int) -> bool:
    i = 0

    while i < len(flowerbed):
        if i == len(flowerbed) - 1:
            if flowerbed[i] == 0 and flowerbed[i - 1] == 0:
                n -= 1
            break

        if i == 0:
            if flowerbed[i] == 0 and flowerbed[i + 1] == 0:
                n -= 1
                i += 2
                continue

        if flowerbed[i] == 0:
            if flowerbed[i + 1] == 0 and flowerbed[i - 1] == 0:
                n -= 1
                i += 2
                continue
        i += 1

    return n <= 0


print(canPlaceFlowers([1, 0, 0, 0, 0, 1], 2))  # False
print(canPlaceFlowers([0, 0, 1, 0, 1], 1))  # True
print(canPlaceFlowers([0, 0, 1, 0, 1, 0, 0], 2))  # True
print(canPlaceFlowers([0, 0, 1, 0, 0, 0], 2))  # True
print(canPlaceFlowers([0, 0, 1, 0, 0, 0, 0], 3))  # True
print(canPlaceFlowers([0, 0, 0, 0, 1], 2))  # True
print(canPlaceFlowers([1, 0], 1))  # False
print(canPlaceFlowers([0, 1, 0], 1))  # False
