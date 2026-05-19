#ifndef LINKED_LIST_H
#define LINKED_LIST_H

struct Node {
    int value;
    struct Node *next;
};

void add_node(int);
void remove_node(int);

#endif /* LINKED_LIST_H */