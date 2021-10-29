//Да простит меня бог за комменты на русском
#include <stdio.h>
#include <memory.h>
#include <stdlib.h>
#include <math.h>


typedef struct {
    int *h;
    int sizeOfHeap;
    int sizeOfArray;
} heapStr;


//Выделение памяти для структуры
heapStr * heapInit(size_t size) 
{
    heapStr *h;
    h = malloc(sizeof(heapStr *));
    h->h = malloc(sizeof(int) * (1+size));
    h->sizeOfArray = (int) size;
    h->sizeOfHeap = 0;
    return h;
}

//Освобождение памяти
void heapFree(heapStr *h) {
    free(h->h);
    free(h);
}

//обыкновенный свап переменных
void swap(int *array, int i, int j) 
{
    int x;
    x = array[j];
    array[j] = array[i];
    array[i] = x;
}

#define LEFT(i) (2*i+1)
#define RIGHT(i) (LEFT(i)+1)
#define PARENT(i) ((i-1) / 2);


//процедура, с помощью которой, узел дерева, в случае если он меньше дочернего,
//сплавляется вниз по дереву(его индекс в массиве увеличивается)
void heapMaxify(heapStr *h, int i) 
{      
    int l, r;
    //Изначально, мы предполагаем, что узел с индексом i обладает наибольшим значением
    int largest = i;

    //определяем левый и правый дочерний узел
    l = LEFT(i);
    r = RIGHT(i);

    //если левый дочерний узел больше чем корень поддерева(корень с индексом largest)
    //то мы изменяем largest на индекс левого или правого дочернего узла
    if (l < h->sizeOfHeap && h->h[l] > h->h[largest]) {
        largest = l;
    }
    if (r < h->sizeOfHeap && h->h[r] > h->h[largest]) {
        largest = r;
    }

    //если всё-таки значение изначально проверяемого узла с индексом i не является самым большим
    //то проверяемый узел и узел с большим значением меняются местами
    if (largest != i) {
        swap(h->h, i, largest);
        
        //Вызываем ещё раз функцию, чтобы в случае чего отправить элемент ниже по пирамиде
        heapMaxify(h, largest);
    }
}

//Инициализирует структуру и из входного массива создаёт невозрастающую пирамиду
heapStr * heapBuild(int *a, int size) 
{
    heapStr *h;
    //Память для структуры
    h = heapInit(size);
    //Копирование элементов из входного массива
    //в массив, который находится в структуре
    memcpy((h->h), a, sizeof(int) * size);

    h->sizeOfHeap = size;
    h->sizeOfArray = size;

    //упорядочивание пирамиды
    for (int i = size/2; i >= 0; i--) {
        heapMaxify(h, i);
    }

    return h;
}


void printArray(int *arr, int n)
{
    for (int i = 0; i < n; ++i) {
        printf("%d, ", arr[i]);
    }
    printf("\n");
}


heapStr *heapSort(heapStr *h) 
{   
    // Каждый элемент, начиная с последнего меняется местами с корневым
    // и сплавляется до подходящего ему места. 
    // Уменьшаем значение размера пирамиды, чтобы уже 
    // отобранные элементы никак не участвовали в операциях
    for (int i = h->sizeOfHeap-1; i > 0; i--) {
        swap(h->h, 0, i);
        h->sizeOfHeap--;
        heapMaxify(h, 0);
    } 
    return h;
}

int queueMax(heapStr *h) {
    return h->h[0];
}

int queueExtraMax(heapStr *h)
{
    //проверка на наличие элементов в пирамиде
    if (h->sizeOfHeap < 1) {
        printf("Size of heap must be bigger than 0\n");
        exit(1);
    }
    int max = h->h[0];

    //Присваиваем первому(пока самому большому) элементу значение
    // последнего, чтобы оно не "потерялось"
    h->h[0] = h->h[h->sizeOfHeap-1];

    // Выделение нового объёма памяти для массива,
    // которое меньше на один int. 
    // Отсекаем последний(бывший самый большой)
    int memSize = (h->sizeOfHeap) * sizeof(h->h[0]);
    h->sizeOfHeap--;
    h->sizeOfArray--;
    h->h = (int*)realloc(h->h, memSize);

    // Приводим пирамиду в порядок.
    heapMaxify(h, 0);
    return max;
}

int queueIncreaseKey(heapStr *h, int i, int key)
{
    if (key < h->h[i]) {
        printf("New key smaller than previous\n");
        exit(0);
    }
    //Определяем родительский узел
    int p = PARENT(i)
    h->h[i] = key;
    // Подобие heapMaxify, только здесь
    // узел поднимается вверх по дереву приводя его в порядок
    while (i > 0 && h->h[p] < h->h[i]) {
        swap(h->h, i, p);
        i = PARENT(i);
        p = PARENT(i)
    }
    return 0;
}

void queueInsertKey(heapStr *h, int key) 
{
    ++h->sizeOfHeap;

    //Устанавливаем наименьшее возможное число
    int bits = 8 * sizeof(int);
    h->h[h->sizeOfHeap-1] = -(1LL<<(bits-1));

    //Увеличиваем размер памяти на один int
    int memSize = (h->sizeOfHeap) * sizeof(h->h[0]);
    h->h = (int*)realloc(h->h, memSize);
    //Вызываем процедуру, чтобы впихнуть новый ключ
    // и построить дерево уже с ним
    queueIncreaseKey(h, h->sizeOfHeap-1, key);
}

#undef LEFT
#undef RIGHT
