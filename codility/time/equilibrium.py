#!/usr/bin/env python


def equilibrium(A, N, s1, s2): 
    
    B = A[0:N//2] 
    C = A[N//2:N]

    N = len(B)
    M = len(C)

    i = N-1
    j = 0
    print(B, C)
    
    # first iteration 
    if s1 == 0 and s2 == 0: 
        for b in B:    
            s1 += b
        for c in C:
            s2 += c

    minimum = abs(s1-s2)
    print(s1, s2, minimum)

    if abs(s1) > abs(s2):
        while i >= 0:
            if abs(s1-s2) < minimum:
                minimum = abs(s1-s2)
            s1 -= B[i]
            s2 += B[i]
            i -= 1
        
    elif abs(s1) < abs(s2):
        while j < M:
            if abs(s1-s2) < minimum:
                minimum = abs(s1-s2)
            s1 += C[j]
            s2 -= C[j]
            j += 1

    elif s1 > s2:
        while j < M:
            if abs(s1-s2) < minimum:
                minimum = abs(s1-s2)
            s1 += C[j]
            s2 -= C[j]
            j += 1

    elif s1 < s2:
        while i >= 0:
            if abs(s1-s2) < minimum:
                minimum = abs(s1-s2)
            s1 -= B[i]
            s2 += B[i]
            i -= 1

    elif s1 == s2:
        return 0

    return minimum
    

A = [9,7,2,3,4,9,10,6,5]
#A = [3,1,2,4,3]
# return 60 expected 20
#A = [-10, -20, -30, -40, 100]
N = len(A)
s1 = 0 
s2 = 0 

r = equilibrium(A, N, s1, s2)

print(r)

