def solution(X, A):
    steps = [False]*(X+1)
    steps[0] = True
    found = False
    s = False
    for i in range(len(A)):
        if A[i] <= X:
            steps[A[i]] = True
        if A[i] == X:
            s = True
        if s:
            for x in range(X+1):
                if steps[x] == False:
                    found = False
                    break
                else:
                    found = True
            if found:
                return i
    return -1
