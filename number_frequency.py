numbers = [1, 3, 56, 78, 45, 68, 78, 1, 3, 56, 91]

def count_frequency(numbers):
    frequency = {}
    for num in numbers:
        if num in frequency:
            frequency[num] += 1
        else:
            frequency[num] =1
            
    return frequency

frequency = count_frequency(numbers)

for number, count in frequency.items():
    print(f"{number}:{count}")