def solution(A):
    N = len(A)
    
    terms = {}
    
    for e in A:
        if e in terms:
            terms[e] += 1
        else:
            terms[e] = 1
    
    for k, v in terms.items():
        if v%2 != 0:
            return k
