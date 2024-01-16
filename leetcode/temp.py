def dailyTemperatures(temperatures):
    #    for i in range(len(temperatures)):
    #        s = 0
    #        for j in range(i + 1, len(temperatures)):
    #            s += 1
    #            if temperatures[j] > temperatures[i]:
    #                change[i] = s
    #                break
    """
    Approche avec une stack dans laquelle on garde les plus grands
    éléments par ordre d'apparition

    Distance d'un elem avec la next valeur + grande

    if current >

    """

    #    [0] * len(temperatures)
    #    stack = []
    #
    #    for i in range(len(temperatures)):
    #        while len(stack) > 0:
    #            if stack[len(stack) - 1] > temperatures[i]:
    #                stack = stack[:-1]
    #                stack.append(temperatures[i])
    #
    #        stack.append(temperatures[i])
    #
    #        print(stack)

    ans = [0] * len(temperatures)
    stack = []
    for i, t in enumerate(temperatures):
        while stack and temperatures[stack[-1]] < t:
            print(i, stack)
            left = stack.pop()
            ans[left] = i - left  # day difference
        stack.append(i)
    return ans


print(dailyTemperatures([73, 74, 75, 71, 69, 72, 76, 73]))  # [1,1,4,2,1,1,0,0]
