vowels = ["a", "e", "i", "o", "u", "A", "E", "I", "O", "U"]


def reverseVowels(s: str) -> str:
    i = 0
    j = len(s) - 1

    lmatch = False
    rmatch = False

    reverse = [""] * len(s)

    while i <= j:
        if s[i] in vowels:
            lmatch = True
        if not lmatch:
            reverse[i] = s[i]
            i += 1

        if s[j] in vowels:
            rmatch = True
        if not rmatch:
            reverse[j] = s[j]
            j -= 1

        if rmatch and lmatch:
            reverse[i], reverse[j] = s[j], s[i]
            lmatch = False
            rmatch = False
            i += 1
            j -= 1

    rev = ""
    for c in reverse:
        rev += c
    return rev


print(reverseVowels("leetcode"))  # leotcede
print(reverseVowels("hello"))  # holle
