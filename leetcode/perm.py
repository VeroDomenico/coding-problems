class Permutation(object):
	def permutations(input_list):
		"""
		Generate all permutations of a list.
		"""
		visited = [False for _ in range(len(input_list))]
		permutations = []
		
		def dfs(current):
			if len(current) == len(input_list):
				permutations.append(current)
			for x in range(len(input_list)):
				if not visited[x]:
					visited[x] = True
					dfs(current+[input_list[x]])
					visited[x] = False
			

		for x in range(len(input_list)):
			visited[x] = True
			dfs([input_list[x]])
			visited[x] = False
		return permutations


if __name__ == '__main__':
	print(Permutation.permutations([x for x in range(1, 4)]))
