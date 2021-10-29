#include <stdio.h>
#include <stdlib.h>


int elementInArray(int element, int array[], int size)
{   
    int elementInArray = 0;
    for (int i = 0; i < size; i++) {
        if (array[i] == element) {
            elementInArray = 1;
            break;
        }
    }
    return elementInArray;
}

void arrayRandomization(int* assistants, int size) 
{   
    int usedIndexes[size];
    int i = 0;
    int randArray[size];
    while (i < size) {
        int randIndex = rand()%size+1;
        if (elementInArray(randIndex, usedIndexes, size) == 0) {
            randArray[i] = assistants[randIndex-1];
            usedIndexes[i] = randIndex;
            i++;
        }
    }
}

int hireAssistant(int *assistants, int size)
{   
    int bits = 8 * sizeof(int);
    int best = -(1LL<<(bits-1));
    for (int i = 0; i < size; ++i) {
        if (assistants[i] > best) {
            best = assistants[i];
        }
    }
    return best;
}


int main(int argc, char **argv) {
    int size = 11;
    int assistants[11] = {1, 2, 4, 9, 5, 8, 100, 86, 79, 985, 1000};
    int best;
    int i = 0;

    // Рандомизация, чтобы не беспокоится о том, в каком порядке нам подаётся массив
    arrayRandomization(assistants, size);
    best = hireAssistant(assistants, size);
    printf("%d\n", best);
    return 0;
}

