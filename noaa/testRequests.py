#!/usr/bin/env python

# more docs at http://www.ncdc.noaa.gov/cdo-web/webservices/v2#gettingStarted

import json
import sys
from urllib2 import Request, urlopen

def fetch(url, token):
	request = Request('http://www.ncdc.noaa.gov/cdo-web/api/v2/datasets', headers={'token':token})
	f = urlopen(request)
	j = json.loads(f.read())
	return j

if __name__ == "__main__":
	# token from http://www.ncdc.noaa.gov/cdo-web/token
	token = sys.argv[1] # to be added to request header as "token"

	print 'datasets'
	for result in fetch('http://www.ncdc.noaa.gov/cdo-web/api/v2/datasets', token)['results']:
		print result

	print 'GHCNDMS dataset'
	for result in fetch('http://www.ncdc.noaa.gov/cdo-web/api/v2/datasets/GHCNDMS', token)['results']:
		print result

	print 'datatypes'
	for result in fetch('http://www.ncdc.noaa.gov/cdo-web/api/v2/datatypes', token)['results']:
		print result

	print 'GHCNDMS data 1763'
	for result in fetch('http://www.ncdc.noaa.gov/cdo-web/api/v2/data/?datasetid=GHCNDMS&startdate=1763-01-01&enddate=1764-01-01', token)['results']:
		print result
