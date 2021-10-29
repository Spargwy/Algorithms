#include <stdio.h>
#include <memory.h>
#include <stdlib.h>
#include <math.h>

typedef struct {
    int *h;
    int sizeOfHeap;
    int sizeOfArray;
} heapStr;


//Initialize heapStr and allocating memory
heapStr * heapInit(size_t size) 
{
    heapStr *h;
    h = malloc(sizeof(heapStr *));
    h->h = malloc(sizeof(int) * (1+size));
    h->sizeOfArray = (int) size;
    h->sizeOfHeap = 0;
    return h;
}

//Memory free
void heapFree(heapStr *h) {
    free(h->h);
    free(h);
}

//Variables swap
void swap(heapStr *h, int i, int j) 
{
    int x;
    x = h->h[j];
    h->h[j] = h->h[i];
    h->h[i] = x;
}

#define LEFT(i) (2*i)
#define RIGHT(i) (LEFT(i)+1)


void heapMaxify(heapStr *h, int i) 
{
    int l, r;
    int largest = i;
    l = LEFT(i);
    r = RIGHT(i);
    if (l < h->sizeOfHeap && h->h[l] > h->h[largest]) {
        largest = l;
    }
    if (r < h->sizeOfHeap && h->h[r] > h->h[largest]) {
        largest = r;
    }
    if (largest != i) {
        swap(h, i, largest);
        heapMaxify(h, largest);
    }
}

void heapMinify(heapStr *h, int i) 
{
    int l, r;
    int smallest = i;
    l = LEFT(i);
    r = RIGHT(i);
    if(l <= h->sizeOfHeap && h->h[l] < h->h[i]) {
        smallest = l;
    }
    if(r <= h->sizeOfHeap && h->h[r] < h->h[smallest]) {
        smallest = r;
        swap(h,i, smallest);
        heapMinify(h, smallest);
    }
    if(smallest != i) {
        swap(h,i, smallest);
        heapMinify(h, smallest);
    }
}

heapStr * heapBuild(int *a, int size, int minify) 
{
    //initialize heap and memory allocation
    heapStr *h;
    h = heapInit(size);

    memcpy((h->h), a, sizeof(int) * size);

    h->sizeOfHeap = size;
    h->sizeOfArray = size;
    if (minify == 1) {
        for (int i = size/2; i >= 0; i--) {
            heapMinify(h, i);
        }
    }
    else {
        for (int i = size/2; i >= 0; i--) {
            heapMaxify(h, i);
    }
    }

    return h;
}

void heapPrintArray(heapStr *h) 
{
    for(int i = 1; i < h->sizeOfArray; i++) {
        printf("%d, ", h->h[i]);
    }
    printf("\n");
} 

heapStr *heapSort(int *array, int size) 
{   
    int arr[size];
    heapStr *h;

    h = heapBuild(array, size, 0);

    for (int i = size-1; i > 2; i--) {
        swap(h, h->h[0], h->h[i]);
        heapMaxify(h, 0);
        h->sizeOfHeap--;
    } 
    return h;
}



#undef LEFT
#undef RIGHT
