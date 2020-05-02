'''
# Sample code to perform I/O:

name = input()                  # Reading input from STDIN
print('Hi, %s.' % name)         # Writing output to STDOUT

# Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
'''

# Write your code here

def time_taken(word_dict, string):
    row, col = word_dict.get(string[0], [-1, -1])
    if row == -1:
        print(string[0], "string[0]")
        return -1
    current_row = row
    current_col = col
    typing_time = 0
    for index in range(1, len(string)):
        row, col = word_dict.get(string[index], [-1, -1])
        if row == -1:
            return -1
        typing_time += abs(current_row - row) + abs(current_col - col)
        current_row = row
        current_col = col
    return typing_time

if __name__ == "__main__":
    word_dict = dict()
    row_col = input()
    row, col = row_col.split(" ")
    for i in range(int(row)):
        keyboard_in = input()
        for index, char in enumerate(keyboard_in):
            word_dict[char] = [i, index]
    string = input()
    print(time_taken(word_dict, string))