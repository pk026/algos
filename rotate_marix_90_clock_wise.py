matrix = [[1, 2, 3], [5, 6, 7], [9, 10, 11]]

def rotate_90_clock_wise(matrix):
    row_count = len(matrix)
    column_count = len(matrix[0])
    rotated_matrix = []
    for i in range(column_count):
        row = []
        for j in range(row_count):
            row.append(None)
        rotated_matrix.append(row)
    for i, row in enumerate(matrix[::-1]):
        for j, column in enumerate(row):
            rotated_matrix[j][i] = column
    return rotated_matrix


def rotate_90_clock_wise_inplace(matrix):
    # only possible for square matrix
    row_count = len(matrix)
    column_count = len(matrix[0])
    for i in range(row_count):
        for j in range(column):
            matrix[i][j], matrix[row_count - 1 - j][column_count -1 - i] = matrix[row_count - 1 - j][column_count - 1 - i], matrix[i][j]

def rotate_90_anti_clock_wise(matrix):
    row_count = len(matrix)
    column_count = len(matrix[0])
    rotated_matrix = []
    for i in range(column_count):
        row = []
        for j in range(row_count):
            row.append(None)
        rotated_matrix.append(row)
    for i, row in enumerate(matrix):
        for j, column in enumerate(row[::-1]):
            rotated_matrix[j][i] = column
    return rotated_matrix


def rotate_180(matrix):
    row_count = len(matrix)
    column_count = len(matrix[0])
    rotated_matrix = []
    for i in range(row_count):
        row = []
        for j in range(column_count):
            row.append(None)
        rotated_matrix.append(row)
    for i, row in enumerate(matrix[::-1]):
        for j, column in enumerate(row[::-1]):
            rotated_matrix[i][j] = column
    return rotated_matrix
