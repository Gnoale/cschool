# Bubble

```
def bubble(a):
    begin = time.time()
    N = len(a)                                                                                                                                                      
    i = j = 0
    while i < N:
        j = 0
        while j < N-i-1:
            if a[j] > a[j+1]:
                a[i], a[j] = a[j], a[i]
            j += 1
        i += 1
    end = time.time()
    print(end-begin)
```

Variant more effective

```
def bubbleplus(a):
    begin = time.time()
    N = len(a)
    i = j = 0
    while i < N-1:
        j = i+1
        while j < N:
            if a[i] > a[j]:
                a[i], a[j] = a[j], a[i]
            j += 1
        i += 1
    end = time.time()
    print(end-begin)
```

```
import time, random, copy

a = [random.randint(0,1000) for a in range(10000)]
b = copy.copy(a)
c = copy.copy(a)
b = copy.copy(a)

a == b == c == d
> True

bubble(a)
9.13589787483

bubbleplus(b)
4.56914591789

bubble(c)
8.92305612564

bubbleplus(d)
4.44120907784
```

# Insertion
Faster than mergesort on little set, article about [timesort algorithm](https://hackernoon.com/timsort-the-fastest-sorting-algorithm-youve-never-heard-of-36b28417f399)

```
def insertion(a):
    begin = time.time()
    N = len(a)
    i = j = 0
    while i < N-1:
        min = i
        if I == i:
            print(*a)
            break
        j = i+1
        while j < N:
            if a[j] < a[min]:
                min = j
            j += 1
        a[i], a[min] = a[min], a[i]
        i += 1
    end = time.time()
    print(end-begin)
```

```
insertion(a)
4.12656998634

insertion(b)
3.8756980896
```
