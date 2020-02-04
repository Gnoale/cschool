def solution(A): 
    # write your code in Python 3.6 
    seq = [False] * (len(A)+1) 
    perm = len(A) 
    missing = None 
     
    for i in range(len(A)): 
        if A[i] < 1:
            missing = 1
            break
        try:
            if seq[A[i]] == False: 
                seq[A[i]] = True 
                perm -= 1 
                missing = A[i]+1 
        except IndexError:
            pass
        if perm == 0:
            return len(A)+1

    if missing is not None:
        for e in range(1,len(seq)): 
            if not seq[e]: 
                return e 
    return missing
