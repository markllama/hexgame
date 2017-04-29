#!/usr/bin/env python

import sys
import lxml.etree as etree
import json

if __name__ == "__main__":

    tree = etree.parse(sys.argv[1])
    root = tree.getroot()

    m = {}

    m['game'] = tree.xpath("/map/@name")[0]
    m['copyright'] = tree.xpath("/map/copyright")[0].text
    m['shape'] = "rectangle"
    #m['author'] = root.xpath("/map/author")[0].text

    esize = tree.xpath("/map/size/vector")[0]
    size = {"hx": esize.get("hx"), "hy": esize.get("hy")}
    eorigin = tree.xpath("/map/origin/vector")[0]
    origin = {"hx": eorigin.get("hx"), "hy": eorigin.get("hy")}
    m['size'] = size
    m['origin'] = origin
    m['terrains'] = []

    eterrains = tree.xpath("/map/terrains")[0]
 #   print eterrains
    for eterrain in eterrains:
        terrain = {}
        terrain['type'] = eterrain.get('type')
        terrain['name'] = eterrain.get('name')
        terrain['locations'] = []

        elocations = list(eterrain)
#        print "there are %d locations" % len(elocations)
        for eloc in elocations:
            evector = list(eloc)
            if len(evector) > 0:
                vec = {"hx": evector[0].get("hx"), "hy": evector[0].get("hy")}
                terrain['locations'].append(vec)
        #terrain = {"hx": eterrain.get("hx"), "hy": eterrain.get("hy")}
        m['terrains'].append(terrain)
    print json.dumps(m, indent=2)
