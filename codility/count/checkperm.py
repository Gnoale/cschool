def solution(A):
    # write your code in Python 3.6
    seq = [False] * (len(A)+1)
    perm = len(A)
    
    for i in range(len(A)):
        try:
            if seq[A[i]] == False:
                seq[A[i]] = True
                perm -= 1
        except IndexError:
            return 0
    
        if perm == 0:
            return 1
    return 0    
    
