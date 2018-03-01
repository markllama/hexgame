#!/usr/bin/env python

import sys, json

if __name__ == "__main__":
    filename = sys.argv[1]
    f = open(filename)

    map = json.load(f)
    
