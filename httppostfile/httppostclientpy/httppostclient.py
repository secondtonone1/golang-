import requests
if __name__ == '__main__':
	files = {"file": open("file.txt", "rb")}
	r = requests.post("http://localhost:8089/upload", files=files)
	print(r.text)