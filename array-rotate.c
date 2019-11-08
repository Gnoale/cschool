#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int * rotate(int A[], int N, int K)
{
    int *B = malloc (sizeof(int)*N);

    for (int i=0; i<N; i++)
      {
        if (i+K < N)
          {
            B[i+K] = A[i];
          }
        else
          {
            int diff = i+K - N;
            B[diff] = A[i];
          }
      }
    return B;
}

void main()
{
  int tab[7] = {1,98,100,5,6,87,0};
  int * result = rotate(tab,7,3);
  for (int i=0; i<7; i++)
    {
      printf("result[%d] = %d\n", i, result[i]);
    }
}
