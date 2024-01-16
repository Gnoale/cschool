#!/usr/bin/env python

A = [2,4,6,8,9,10,12,18,22,24,39,39]
B = [0,1,4,5,7,8,9,11,13,14,15,18,20,22,24,24,39]

# Find the intersections

# Linear search

def linears(A,B):
    i = 0
    j = 0
    while i < len(A) and j < len(B):
        if A[i] == B[j]:
            print("intersection :", A[i], B[j], i, j)
            i += 1
            j += 1
        elif A[i] < B[j]:
            i += 1
        else:
            j += 1

# Recursive hashtable search 

def bsection(A,d):

    if len(A) > 1:
        L = A[:len(A)//2]
        R = A[len(A)//2:]
        bsection(L,d)
        bsection(R,d)

    if len(A) == 1:
        if A[0] in d:
            d[A[0]] += 1
        else:
            d[A[0]] = 1
        print("Leaf {}".format(A))
        return A

# merge sort

def merge_sort(A):

    if len(A) == 1:
        print("Leaf {}".format(A))
        return A

    L = A[:len(A)//2]
    R = A[len(A)//2:]
    print("Left tree {}, Right Tree {}".format(L,R))
    # At one point, L and R length will be == 1
    L = merge_sort(L)
    R = merge_sort(R)
    print("Merge {} and {}".format(L,R))
    return merge(L,R)

def merge(L,R):
    i = 0
    j = 0
    A = [] 
    while i < len(L) and j < len(R):
        if L[i] < R[j]:
            A.append(L[i])
            i += 1
        else:
            A.append(R[j])
            j += 1
    while j < len(R):
        A.append(R[j])
        j += 1
    while i < len(L):
        A.append(L[i])
        i += 1

    print("Merged array {}".format(A))
    print()
    return A


if __name__ == '__main__':
    import random
    linears(A,B)
#    A = [random.randint(0,20) for a in range(20)]
#    A = [8, 9, 3, 6, 11, 14, 5, 8, 4, 14, 4, 0, 2, 1, 12, 0, 10, 18, 5, 5]
#    A = [8,1,0,4,2]
#    merge_sort(A)
#    print(A)
#    print(B)
#    merge(A,B)   
#    d = {}
#    e = {}
#    import random
#    A = [random.randint(500,1500) for a in range(1000)]
#    B = [random.randint(1300,2300) for a in range(1000)]
#    bsection(A,d)
#    bsection(B,e)
#    print(d)
#    print(e)
#
#    for k, v in d.items():
#        if k in B:
#            print("Value {} found {} times in A, and {} times in B".format(k,v,e[k]))

    
