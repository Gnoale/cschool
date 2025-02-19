def removeStars(s: str) -> str:
    rstar = ""
    i = 0
    while len(s) > 0:
        """
        Remove char left to the star as well as the star itself

        """
        # pop the latest char
        c = s[len(s) - 1]
        s = s[:-1]

        # we store the original string chars  from right to left in a new one
        # each time we find a * we increment a counter
        if c == "*":
            i += 1

        if i == 0:
            rstar += c

        # while the star counter is not empty, just decrease the char and decrement it
        if c != "*" and i > 0:
            i -= 1

    star = ""
    for i in range(len(rstar) - 1, -1, -1):
        star += rstar[i]

    return star


print(removeStars("leet**cod*e"))  # lecoe
print(removeStars("erase*****"))  #
