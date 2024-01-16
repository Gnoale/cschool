#!/usr/bin/env python

# Extract the sentence from a given text and 
# sort them by number of words
# A sentence begin with a capital and end with a dot.
# A word is a set of alphanumeric characters.
# Ignore whitespace and punctuation.

import string


def readtext(infile):
    s = ""
    with open(infile,'r') as f:
        s = f.read()

    dot = [".","!","?"]
    sentences = []
    sentence = []
    word = ""
    
    for i in range(len(s)):
        # words
        if s[i].isalnum() or s[i] == "'":
            word += s[i]
        elif word:
            sentence.append(word)
            word = ""
        # sentences
        if s[i] in dot:
            #print(sentence)
            if sentence:
                #print(sentence)
                sentences.append(sentence)
                sentence = []

    [print(s) for s in sentences]
    return sentences

def countword(s):
    wc = 0
    for e in s:
        nwc = len(e)
        if nwc > wc:
            wc = nwc

    print("Biggest sentence count {} words".format(wc))
    return wc

if __name__ == '__main__':

    infile = "sample.txt"
    s = readtext(infile)
    countword(s)
