from stats import get_num_words
from stats import count_characters
from stats import sort_distinct_char_count
import sys

def get_book_text(book_filepath: str):
    with open(book_filepath) as book:
        book_contents = book.read()
        return book_contents

def print_report(book_path: str, num_words: int, chars_sorted_list: dict):
    print("============ BOOKBOT ============")
    print(f"Analyzing book found at {book_path}...")
    print("----------- Word Count ----------")
    print(f"Found {num_words} total words")
    print("--------- Character Count -------")
    for item in chars_sorted_list:
        if not item["char"].isalpha():
            continue
        print(f"{item['char']}: {item['num']}")

    print("============= END ===============")


def main():
    if len(sys.argv) < 2:
        print("Usage: python3 main.py <path_to_book>")
        sys.exit(1)
    contents = get_book_text(book_filepath=sys.argv[1])
    num_words = get_num_words(book_contents=contents)
    counted_characters = sort_distinct_char_count(count_characters(contents))
    print_report(book_path=sys.argv[1], num_words=num_words, chars_sorted_list=counted_characters)

if __name__ == "__main__":
    main()