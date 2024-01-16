def compress(chars) -> int:
    # Start pointers at the beginning of the list
    read, write, length = 0, 0, len(chars)

    # Process the entire character list
    while read < length:
        # Move read pointer to end of current character sequence
        read_next = read + 1

        while read_next < length and chars[read_next] == chars[read]:
            read_next += 1

        # Write the current character to the write pointer
        chars[write] = chars[read]
        write += 1

        # If we have a sequence longer than 1
        if read_next - read > 1:
            # Convert count to string and write each digit
            count = str(read_next - read)

            for char in count:
                chars[write] = char

                write += 1

        # Move read pointer to next new character
        read = read_next

    print(chars)
    return write


print(compress(["a", "a", "b", "b", "c", "c", "c"]))
print(compress(["a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"]))
print(compress(["a"]))
print(compress(["a", "a", "a", "a", "b", "b", "a", "a"]))
