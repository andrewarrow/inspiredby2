import requests
import os
import uuid

PIKA = os.getenv("PIKA")

url = 'https://api.pika.art/generate'

headers = {
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:129.0) Gecko/20100101 Firefox/129.0',
    'Accept': '*/*',
    'Accept-Language': 'en-US,en;q=0.5',
    'Accept-Encoding': 'gzip, deflate, br, zstd',
    'Referer': 'https://pika.art/',
    'Authorization': "Bearer {}".format(PIKA),
    'Content-Type': 'multipart/form-data; boundary=---------------------------266926460920144731353527800262',
    'Origin': 'https://pika.art',
    'Connection': 'keep-alive',
    'Sec-Fetch-Dest': 'empty',
    'Sec-Fetch-Mode': 'cors',
    'Sec-Fetch-Site': 'same-site',
    'Priority': 'u=4',
    'Pragma': 'no-cache',
    'Cache-Control': 'no-cache',
    'TE': 'trailers',
}

data = (
    b'-----------------------------266926460920144731353527800262\r\n'
    b'Content-Disposition: form-data; name="styleId"\r\n\r\n\r\n'
    b'-----------------------------266926460920144731353527800262\r\n'
    b'Content-Disposition: form-data; name="promptText"\r\n\r\ncollaborating with athletic trainers to analyze data\r\n'
    b'-----------------------------266926460920144731353527800262\r\n'
    b'Content-Disposition: form-data; name="options"\r\n\r\n{"frameRate":24,"parameters":{"guidanceScale":12,"motion":1},"camera":{"zoom":null,"pan":null,"tilt":null,"rotate":null},"extend":false}\r\n'
    b'-----------------------------266926460920144731353527800262\r\n'
    b'Content-Disposition: form-data; name="userId"\r\n\r\n8d46d0a2-7e9a-46d8-9796-f0458ccb4171\r\n'
    b'-----------------------------266926460920144731353527800262--\r\n'
)

response = requests.post(url, headers=headers, data=data)

print(response.text)



