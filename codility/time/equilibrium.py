A = [9,7,2,3,4,9,10,6,5]

s1 = 0 
s2 = 0 
R = 0 
B = []

def equilibrium(A, B, N, s1, s2, R): 
    
    # finished 
    if N==1: 
        return s1-s2

    A = A[0:N//2-1] 
    B = A[N//2:N] 
    i = N//2-1 
    j = N//2 

    # first iteration 
    if s1 == 0 and s2 == 0: 
        while i > 0: 
            s1 += A[i] 
        while j < N: 
            s2 += A[i] 
    
    elif s1 > s2:
        while
