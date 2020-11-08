import ctypes as c
import json

# Load Topic library from Go
LibTopic = c.cdll.LoadLibrary("./libtopic.so")


def __c_str(string):
    """
    :param string: Python String
    :return C String:
    """
    return c.c_char_p(string.encode('utf-8'))


def __go_get_topic(c_string):
    """
    :param c_string: C String
    :return: C String
    """
    # Prepare to call GetTopic func from Go
    LibTopic.GetTopic.restype = c.c_char_p
    LibTopic.GetTopic.argtypes = [c.c_char_p]

    # Return result of C String
    retval = LibTopic.GetTopic(__c_str(c_string))
    return retval


def __append_topic_result(parent_dict, child_dict):
    """
    :param parent_dict: Dictionay
    :param child_dict: Dictionary
    :return: Dictionary
    """

    for key in child_dict:

        # Get value from child dict
        value = float(child_dict[key])

        # Gather the highest value for that topic for each word.
        if key in parent_dict:
            parent_value = float(parent_dict[key])
            if parent_value < value:
                parent_dict[key] = value
        else:
            parent_dict[key] = value

    # Return dictionary with highest values
    return parent_dict


def __call_x_get_topic(word, n=100):
    """
    :param word: single word
    :param n: it's time in the loop
    :return: a combined dictionary
    """
    result_dict = {}
    for _ in range(n):
        result_topic = __go_get_topic(word)
        topic_dict = json.loads(result_topic)
        result_dict = __append_topic_result(result_dict, topic_dict)

    return result_dict


def filter_text(c_string):
    """
    :param c_string: C String
    :return Json String
    """

    # Split text to array
    words = c_string.split()

    # Combine of result all word
    result_dict = {}

    # Use for loop call go_get_topic for each word
    for word in words:

        # Call 100x
        combined_dict = __call_x_get_topic(word, 100)
        result_dict = __append_topic_result(result_dict, combined_dict)

    # Summary and return
    retval = json.dumps(result_dict).encode('utf-8')
    return retval
