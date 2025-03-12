import pytest

from sorting import (
    custom_sort_string,
    median_sliding_window,
    find_missing_and_repeated_numbers,
    custom_sort_string_frequency,
    closest_prime_number,
    alternate_colors,
    find_provinces,
    number_of_substrings,
)


def test_custom_sort_string():
    assert custom_sort_string("cba", "abcd") == "cbad"
    assert custom_sort_string("exv", "xwvee") == "eexvw"


def test_custom_sort_string_frequency():
    assert custom_sort_string_frequency("cba", "abcd") == "cbad"
    assert custom_sort_string_frequency("exv", "xwvee") == "eexvw"
    assert custom_sort_string_frequency("bcafg", "abcd") == "bcad"


def test_median_sliding_window():
    assert median_sliding_window([1, 3, -1, -3, 5, 3, 6, 7], 3) == [1, -1, -1, 3, 5, 6]


def test_find_missing_and_repeated_numbers():
    assert find_missing_and_repeated_numbers([[9, 1, 7], [8, 9, 2], [3, 4, 6]]) == [
        9,
        5,
    ]


def test_closest_prime_number():
    assert closest_prime_number(10, 20) == [11, 13]
    assert closest_prime_number(10, 10) == [-1, -1]
    assert closest_prime_number(1, 10000) == [2, 3]


def test_alternate_colors():
    assert alternate_colors([0, 1, 0, 1, 0], 3) == 3
    assert alternate_colors([0, 1, 0, 0, 1, 0, 1], 6) == 2
    assert alternate_colors([1, 1, 0, 1], 4) == 0
    assert alternate_colors([0, 0, 0, 1], 3) == 1


def test_find_provinces():
    assert find_provinces([[1, 1, 0], [1, 1, 0], [0, 0, 1]]) == 2
    assert find_provinces([[1, 0, 0], [0, 1, 0], [0, 0, 1]]) == 3
    assert find_provinces([[1, 0, 0, 1], [0, 1, 1, 0], [0, 1, 1, 1], [1, 0, 1, 1]]) == 1


def test_number_of_substrings():
    assert number_of_substrings("abcabc") == 10
