def solution(A):
    N = len(A)+1
    s = 0
    
    expected = [n for n in range(N+1)]
    
    for e in expected:
        s+=e
    
    s1 = 0
    for e in A:
        s1+=e
    
    return s-s1
