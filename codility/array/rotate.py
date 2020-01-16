def solution(A, K):
    # write your code in Python 3.6
    i = 0
    N = len(A)
    
    if N == 0:
        return A

    if K > N:
        K = K%N
    
    if K == 0:
        return A
    
    B = [0]*N
    
    while i < N:
        if i+K < N:
            B[i+K] = A[i]
        else:
            j = (i+K)%N
            B[j] = A[i]
        i += 1
        
    return B
