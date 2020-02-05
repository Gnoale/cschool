def solution(A): 
    # write your code in Python 3.6 
    seq = [False] * (len(A)+1) 
    perm = len(A) 
    missing = None 
     
    for i in range(len(A)): 
        if A[i] < 1:
            missing = 1
        else:
            try:
                if seq[A[i]] == False: 
                    seq[A[i]] = True 
                    perm -= 1 
            except IndexError:
                pass
        # the sequence is complete from 1 to n
        # minimum missing element is len(A)+1
        if perm == 0:
            return len(A)+1


    # the missing element is somewhere in the sequence
    for e in range(1,len(seq)): 
        if not seq[e]: 
            return e
    
    # the missing element is 1
    return missing
