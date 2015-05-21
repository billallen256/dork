#!/usr/bin/env python3

# Based on poor examples from
# http://www.ncdc.noaa.gov/cdo-web/webservices/v2#data

# You'll need to get a token from NOAA and pass it as
# a parameter to this script.  See main().

from pprint import pprint
from urllib.parse import urlencode
from urllib.request import Request, urlopen

import json
import sys

def fetch_locations(token):
	headers = { 'token':token }
	params = urlencode({
			'datasetid':'GHCND',
			'locationcategoryid':'CITY',
			'limit':1000,
			'offset':1000 })

	req = Request('http://www.ncdc.noaa.gov/cdo-web/api/v2/locations?{0}'.format(params), headers=headers)
	resp = urlopen(req)
	pprint(resp.status)
	d = str(resp.read(), encoding='utf-8')
	j = json.loads(d)
	pprint(j)

def fetch_location(location, token):
	headers = { 'token':token }
	params = urlencode({
			'startdate':'2015-05-15',
			'enddate':'2015-05-18',
			'datasetid':'GHCND',
			'datatypeid':'TMAX',
			'limit':1000,
			'locationid':location})

	req = Request('http://www.ncdc.noaa.gov/cdo-web/api/v2/data?{0}'.format(params), headers=headers)
	resp = urlopen(req)
	pprint(resp.status)
	d = str(resp.read(), encoding='utf-8')
	j = json.loads(d)
	pprint(j)
	stations = set([ r['station'] for r in j['results'] ])
	pprint(stations)

def fetch_station(station, token):
	headers = { 'token':token }
	params = urlencode({
			'startdate':'2015-05-15',
			'enddate':'2015-05-18',
			'datasetid':'GHCND',
			'datatypeid':'TMAX',
			'stationid':station,
			'limit':1000})

	req = Request('http://www.ncdc.noaa.gov/cdo-web/api/v2/data?{0}'.format(params), headers=headers)
	resp = urlopen(req)
	pprint(resp.status)
	d = str(resp.read(), encoding='utf-8')
	j = json.loads(d)
	pprint(j)

if __name__ == "__main__":
	token = sys.argv[1]
	#fetch_location('CITY:US360019', token)
	fetch_station('GHCND:USW00094728', token)
