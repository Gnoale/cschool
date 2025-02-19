package leetcode75

import math


def gcdOfStrings(str1: str, str2: str) -> str:
    """
    ABABAB
    ABAB
    GCD = AB
    """

    i = 0
    j = 0

    gcd = ""

    while i < len(str1) and j < len(str2):
        if math.gcd(len(str1), len(str2)) == len(gcd):
            break

        if str1[i] == str2[j]:
            gcd += str1[i]
            i += 1
            j += 1
        else:
            j += 1

    if len(gcd) == 0:
        return ""

    # check if we concatenate the string "q" times
    # that we really get the original string
    q = len(str1) / len(gcd)
    check = ""
    for i in range(0, int(q)):
        check += gcd
    if check != str1:
        return ""


    q = len(str2) / len(gcd)
    check = ""
    for i in range(0, int(q)):
        check += gcd
    if check != str2:
        return ""

    return gcd


print(gcdOfStrings("ABABAB", "ABAB"))  # AB
print(gcdOfStrings("ABABABAB", "ABAB"))  # ABAB
print(gcdOfStrings("ABCABC", "ABC"))  # ABC
print(gcdOfStrings("ABCDEF", "ABC"))  #
print(gcdOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"))  # TAUXX
print(gcdOfStrings("AA", "A"))  # A
print(gcdOfStrings("LEET", "CODE"))  #
