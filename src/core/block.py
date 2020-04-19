import calendar
import time
import hashlib
import json

class Meta:
  def __init__(self, index, timestamp, prev_hash):
    self.index = index
    self.timestamp = timestamp
    self.prev_hash = prev_hash
    self.hash = ''
    self.nonce = 0
  
  def to_dict(self):
    return {
      'index': self.index,
      'timestamp': self.timestamp,
      'prev_hash': self.prev_hash,
      'hash': self.hash,
      'nonce': self.nonce,
    }

class Block:
  def __init__(self, index, prev_hash, data={}):
    timestamp = calendar.timegm(time.gmtime())

    self.data = data
    self.meta = Meta(index, timestamp, prev_hash)
    self.meta.hash = self.create_hash()


  def to_dict(self):
    return {
      'meta': self.meta.to_dict(),
      'data': self.data,
    }

  def __str__(self):
    return json.dumps(self.to_dict())

  def create_hash(self, nonce=0):
    h = hashlib.sha256()
    willhash = bytes(f'{self.meta.index}{self.meta.timestamp}{self.meta.prev_hash}{nonce}{self.data}'.encode('utf-8'))
    h.update(willhash)
    return h.hexdigest()

  def proof_of_work(self, difficulty):
    while(self.meta.hash[:difficulty['times']] != difficulty['char'] * difficulty['times']):
      self.meta.nonce += 1
      self.meta.hash = self.create_hash(self.meta.nonce)
