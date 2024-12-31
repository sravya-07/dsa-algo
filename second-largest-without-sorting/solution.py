def second_largest_smallest(list1):
    largest, smallest, second_largest, second_smallest = list1[0], float('inf'), list1[0], float('inf')
    for item in list1:
        if item > largest:
            second_largest = largest
            largest = item
        elif item > second_largest and item != largest:
            second_largest = item
        
        if item < smallest:
            second_smallest = smallest
            smallest = item
        elif item < second_smallest and item != smallest:
            second_smallest = item
    return second_largest, second_smallest

list1 = [1, 2, 3, 4, 5, 8, 8, 9, 1, 10, 11]
max1 = list1[0]
min1 = list1[0]

for item in list1:
    if item > max1:
        max1 = item
    if item < min1:
        min1 = item

max2 = min1
min2 = max1
for item in list1:
    if item > max2 and item != max1:
        max2 = item
    if item < min2 and item != min1:
        min2 = item 

print("Max2 = {}, Min2 = {}".format(max2, min2))

print(second_largest_smallest(list1))