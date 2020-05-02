class Node:

	def __init__(self, data):
		self.left = None
		self.right = None
		self.data = data

def get_vertical_sum():
	position_sum_dict = dict()
	traverse_tree_vertical_sum(root, 0, position_sum_dict)
	for index, value in position_sum_dict.items():
		print("Sum at position {} is {}".format(index, value))


def traverse_tree_vertical_sum(root, position, position_sum_dict):
	if not root:
		return
	position_sum = position_sum_dict.get(position, 0)
	position_sum += root.data
	position_sum_dict[position] = position_sum
	if root.left:
		traverse_tree_vertical_sum(root.left, position - 1, position_sum_dict)
	if root.right:
		traverse_tree_vertical_sum(root.left, position + 1, position_sum_dict)

