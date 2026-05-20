#ifndef LINKED_LIST_H
#define LINKED_LIST_H

#include <stdbool.h>

typedef struct Node {
    int value;
    struct Node *next;
} Node;

extern struct Node *head;

bool add_node(int val);
bool remove_node(int idx);
void print_list(void);

#endif /* LINKED_LIST_H */