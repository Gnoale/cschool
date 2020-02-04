def solution(X, A):
    
    steps = [False] * (X+1)
    steps[0] = True
    passage = X
    
    for i in range(len(A)):
        if steps[A[i]] == False:
            passage -= 1
            steps[A[i]] = True
        if passage == 0:
            return i
    
    return -1

