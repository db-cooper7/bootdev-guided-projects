def get_num_words(book_contents: str):
    words = book_contents.split()
    return len(words)

def count_characters(book_contents: str):
    lower_book_contents = book_contents.lower()
    distinct_char_count = {}
    for char in lower_book_contents:
        if char not in distinct_char_count:
            distinct_char_count[char] = 1
        else:
            distinct_char_count[char] += 1
    return distinct_char_count

# Helper function to sort based on highest char count
def sort_by_frequency(distinct_char_count:dict):
    return distinct_char_count["num"]

def sort_distinct_char_count(distinct_char_count: dict):
    char_frequency_list = [{"char": char, "num": distinct_char_count[char]} for char in distinct_char_count]
    char_frequency_list.sort(key=sort_by_frequency, reverse=True)
    return char_frequency_list