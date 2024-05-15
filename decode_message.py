def decode(message_file):
    with open(message_file, 'r') as file:
        lines = file.readlines()

    # Create a dictionary to store words with their corresponding numbers
    words_dict = {}
    for line in lines:
        num, word = line.strip().split(' ')
        words_dict[int(num)] = word

    # Determine the number of rows in the pyramid
    rows = 0
    while sum(range(rows + 1)) <= len(words_dict):
        rows += 1

    # Reconstruct the pyramid and extract the message
    message = ''
    for i in range(1, rows + 1):
        message += words_dict[i] + ' '

    return message.strip()

# Example usage
message_file = 'C:\\Users\\anilc\\INT-GO\\coding_qual_input.txt'
decoded_message = decode(message_file)
print(decoded_message)
