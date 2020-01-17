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

```
import time, random, copy

a = [random.randint(0,1000) for a in range(1000)]
b = copy.copy(a)
c = copy.copy(a)
d = copy.copy(a)

a == b == c == d
> True

bubble(a)
9.13589787483

bubble(c)
8.92305612564

```

# Insertion
Faster than mergesort on little set, article about [timesort algorithm](https://hackernoon.com/timsort-the-fastest-sorting-algorithm-youve-never-heard-of-36b28417f399)
Shift every element to the right while element j-1 > tmp.
Fill the hole with tmp then.

```
def insertion(a):
    begin = time.time()
    N = len(a)
    i = 1
    while i < N:
      j = i
      tmp = a[i]
      while j > 0 and a[j-1] > tmp:
        a[j] = a[j-1]
        j -= 1
      a[j] = tmp
    i += 1
    end = time.time()
    print(end-begin)
```


# Selection

```
def selection(a):
    begin = time.time()
    N = len(a)
    i = j = 0
    while i < N-1:
        min = i
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


# Snippet to find a numeric palindrom

``̀`
def is_numeric_palindrome(n):
  arr = []
  while(n > 0):
    arr.append(n % 10)
    n = n // 10

  print(arr)

  num_len = len(arr)
  for i in range(num_len // 2):
    if arr[i] != arr[num_len - i - 1]:
      return False

  return True
```





