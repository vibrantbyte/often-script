#include <stdio.h>
int main()
{
    printf("Hello world! \n"); // 教科书的写法
    printf("%ld \n",8 * sizeof(uintptr_t));
    printf("%d \n",64%64);
    printf("%d \n",(1 << 6) % 64);
    return 0;
}