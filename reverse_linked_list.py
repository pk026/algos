class Node:

    def __init__(self, data):
        self.next = None
        self.data = data


class Stack:

    def __init__(self):
        self.values = []
        self.position = -1

    def push(self, value):
        self.values.append(value)
        self.position += 1

    def pop(self):
        value = self.values[self.position]
        self.position -= 1
        return value

    def top(self):
        return self.values[self.position]


def reverse_list(root, k):
    stack = Stack()
    moving_root = root
    for i in range(k):
        print("i>>>>", i)
        if moving_root:
            stack.push(moving_root)
            moving_root = moving_root.next
    print(moving_root.data)
    print(stack.values)
    print(stack.position, "position")
    new_list = None
    if stack.position > 0:
        new_list = stack.pop()
    temp_head = new_list
    while stack.position >= 0 and temp_head:
        node = stack.pop()
        temp_head.next = node
        temp_head = node
    if moving_root:
        temp_head.next = moving_root
    return new_list

def string_couting(input_str):
    new_str = ""
    first_ele = input_str[0]
    count = 0
    for in_str in input_str:
        if first_ele == in_str:
            count += 1
        else:
            new_str += first_ele
            new_str += str(count)
            first_ele = in_str
            count = 1
    if count > 0:
        new_str += first_ele
        new_str += str(count)
    return new_str






