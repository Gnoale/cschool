#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int jump(int X, int Y, int D) 
{
  if ( (Y - X) % D == 0)
  {
    return (Y - X)/D;
  }
  else
  {
    return (Y - X)/D + 1;
  }
}


int main() 
{
   int steps = jump(8, 150, 9);
   printf("Steps = %d\n", steps);

}





