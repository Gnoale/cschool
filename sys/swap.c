#include <stdio.h>
#include <string.h>

void main()
{

  void swap(void* vp1, void* vp2, int size) 
  {  
    int* buffer[size];
    
    memcpy(buffer, vp1, size);
    memcpy(vp1, vp2, size);
    
    printf("buffer[0] = %p\n", buffer[0]);
    memcpy(vp2, buffer, size);
  }

  int g1 = 20900988;
  int g2 = 1;

  printf("g1 before swap = %i\n", g1);
  printf("g2 before swap = %i\n", g2);

  printf("&g1 = %p\n", &g1);
  printf("&g2 = %p\n", &g2);
  
  swap(&g1, &g2, sizeof(int*));

  printf("-------------------------\n");

  printf("g1 after swap = %i\n", g1);
  printf("g2 after swap = %i\n", g2);

  printf("&g1 = %p\n", &g1);
  printf("&g2 = %p\n", &g2);
 



}
