#include "heapify_for_article.c"


int main(int argc, char **argv) {
    int a[] = {16, 4, 10, 14, 7, 9, 3, 2, 8, 1};
    int size = sizeof(a) / sizeof(a[0]);
    heapStr *h;

    h = heapBuild(a, size);

    //изменение значения ключа
    queueIncreaseKey(h, 4, 17); //[17, 16, 10, 8, 14, 9, 3, 2, 4, 1]

    //Получение и удаление максимального значения
    int max = queueMax(h);
    printf("%d\n", max);
    max = queueExtraMax(h);//17

    //Добавление ключа
    queueInsertKey(h, 100);//100, 16, 10, 8, 14, 9, 3, 2, 4, 1
    heapFree(h);
    return 0;
}