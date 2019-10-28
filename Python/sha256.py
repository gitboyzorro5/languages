import hashlib

def hash_string(string):

	"""
	return a sha256 hash of the given string
	"""

	return hashlib.sha256(string.encode('utf-8')).hexdigest()

print (len((hash_string("Hello"))))