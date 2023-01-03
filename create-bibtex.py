#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""
Utilise zotero to establish meta-information and create bibtext file
"""
import requests
import json

'''
Use this website to parse the curl calls <https://curlconverter.com/python/>
'''

headers = {
    'authority': 't0guvf0w17.execute-api.us-east-1.amazonaws.com',
    'accept': '*/*',
    'accept-language': 'en-US,en;q=0.9',
    'content-type': 'text/plain',
    'dnt': '1',
    'origin': 'https://zbib.org',
    'sec-ch-ua': '"Not?A_Brand";v="8", "Chromium";v="108", "Google Chrome";v="108"',
    'sec-ch-ua-mobile': '?0',
    'sec-ch-ua-platform': '"Windows"',
    'sec-fetch-dest': 'empty',
    'sec-fetch-mode': 'cors',
    'sec-fetch-site': 'cross-site',
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36',
}

data = 'https://www.nature.com/articles/s41569-021-00665-7'

response = requests.post('https://t0guvf0w17.execute-api.us-east-1.amazonaws.com/Prod/web', headers=headers, data=data, verify= False)


headers = {
    'authority': 't0guvf0w17.execute-api.us-east-1.amazonaws.com',
    'accept': '*/*',
    'accept-language': 'en-US,en;q=0.9',
    'content-type': 'application/json',
    'dnt': '1',
    'origin': 'https://zbib.org',
    'sec-ch-ua': '"Not?A_Brand";v="8", "Chromium";v="108", "Google Chrome";v="108"',
    'sec-ch-ua-mobile': '?0',
    'sec-ch-ua-platform': '"Windows"',
    'sec-fetch-dest': 'empty',
    'sec-fetch-mode': 'cors',
    'sec-fetch-site': 'cross-site',
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36',
}

params = {
    'format': 'bibtex',
}
response.content

json_data = response.json()

response = requests.post(
    'https://t0guvf0w17.execute-api.us-east-1.amazonaws.com/Prod/export',
    params=params,
    headers=headers,
    json=json_data,
    verify= False
)
