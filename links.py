import feedparser
import requests
import json

feed = feedparser.parse("https://www.inoreader.com/stream/user/1005349717/tag/save")

links = []

for entry in feed.entries:
    links.append(entry.link)
    print(entry.link)


url = 'https://t0guvf0w17.execute-api.us-east-1.amazonaws.com/Prod/web'

headers = {

  'authority': 't0guvf0w17.execute-api.us-east-1.amazonaws.com' ,
  'accept': '*/*' ,
  'accept-language': 'en-US,en;q=0.9' ,
  'content-type': 'text/plain' ,
  'dnt': '1' ,
  'origin': 'https://zbib.org' ,
  'sec-ch-ua': '"Not?A_Brand";v="8", "Chromium";v="108", "Google Chrome";v="108"' ,
  'sec-ch-ua-mobile': '?0' ,
  'sec-ch-ua-platform': '"macOS"' ,
  'sec-fetch-dest': 'empty' ,
  'sec-fetch-mode': 'cors' ,
  'sec-fetch-site': 'cross-site' ,
  'user-agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36' 
}

outinfo = []
for link in links:
    body = {'data-raw': link, "compressed": "true" }
    r = requests.post(url, data=body, headers=headers)
    if r.status_code == 200 :
      outinfo = outinfo + json.loads(r.content.decode("utf-8"))

print(outinfo)

with open('bibinfo.json', 'w') as out_file:
     json.dump(outinfo, out_file, sort_keys = True, indent = 4,
               ensure_ascii = False)


url2 = 'https://t0guvf0w17.execute-api.us-east-1.amazonaws.com/Prod/export?format=bibtex'

headers2 = {
  'authority': 't0guvf0w17.execute-api.us-east-1.amazonaws.com' ,
  'accept': '*/*' ,
   'accept-language': 'en-US,en;q=0.9' ,
   'content-type': 'application/json' ,
   'dnt': '1' ,
   'origin': 'https://zbib.org' ,
   'sec-ch-ua': '"Not?A_Brand";v="8", "Chromium";v="108", "Google Chrome";v="108"' ,
   'sec-ch-ua-mobile': '?0' ,
   'sec-ch-ua-platform': '"macOS"' ,
   'sec-fetch-dest': 'empty' ,
   'sec-fetch-mode': 'cors' ,
   'sec-fetch-site': 'cross-site' ,
   'user-agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36' 
}


outinfo = str(outinfo).replace("'", '"')

body2 = {'data-raw': json.dumps(outinfo) }

r = requests.post(url2, data=body2, headers=headers2)

r.content