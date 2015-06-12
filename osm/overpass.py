#!/usr/bin/env python3

# test queries at http://overpass-turbo.eu/

from pprint import pprint

import overpy # https://github.com/DinoTools/python-overpy

if __name__ == "__main__":
	api = overpy.Overpass()
	result = api.query('''
node
	["name"~"."]
	(39.0,-105.0,39.1,-104.9);
out body;
''')

	for node in result.nodes:
		pprint(node.tags)
