#include <stdio.h>

int main ()
{
    long a;
    long sum;
    /* for loop execution */
    for( a = 0; a < 10000000; a++ )
    {
        sum += a;
    }
    printf("sum: %ld\n", sum);
    return 0;
}
