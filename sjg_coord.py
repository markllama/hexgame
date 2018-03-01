#!/usr/bin/env python

from __future__ import print_function
import sys

# rx = hx + ox
# ry = hy + oy
# rz = (hy + oy) - (hx - ox)

if __name__ == "__main__":
    hx = int(sys.argv[1])
    hy = int(sys.argv[2])

    nx = hx - hy + 7
    ny = hx - 1
    #print('{{ \'hx\': {0}, \'hy\': {1} }} # hz: {2}'.format(hx, hy, hy - hx))
    #print('{{ \'hx\': {0}, \'hy\': {1} }} # hz: {2}'.format(nx, ny, ny - nx))
    print('{{ \'hx\': {0}, \'hy\': {1} }}'.format(nx, ny))
