#!/usr/bin/env python

# more docs at http://www.ncdc.noaa.gov/cdo-web/webservices/v2#gettingStarted

import sys
from urllib2 import Request

if __name__ == "__main__":
	# token from http://www.ncdc.noaa.gov/cdo-web/token
	token = sys.argv[1] # to be added to request header as "token"

http://www.ncdc.noaa.gov/cdo-web/api/v2/datasets

http://www.ncdc.noaa.gov/cdo-web/api/v2/datasets/GHCNDMS

http://www.ncdc.noaa.gov/cdo-web/api/v2/datatypes

http://www.ncdc.noaa.gov/cdo-web/api/v2/data/?datasetid=GHCNDMS&startdate=1763-01-01&enddate=1764-01-01"
