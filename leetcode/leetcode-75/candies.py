def kidsWithCandies(candies: list[int], extraCandies: int) -> list[bool]:
    mc = 0

    extra = [False] * len(candies)

    # first find max
    for candy in candies:
        if candy > mc:
            mc = candy


    for i in range(len(candies)):
        if candies[i] + extraCandies >= mc:
            extra[i] = True
        print(candies[i])

    return extra


print(kidsWithCandies([2, 3, 5, 1, 3], 3))
