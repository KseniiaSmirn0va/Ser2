import hashlib, binascii
hashlib.pbkdf2_hmac('sha256', b'password', b'salt', 100000)
