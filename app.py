import json
from lib.chain import Chain

c = Chain('koinku', difficulty={ 'char': '1', 'times': 1 })
c.add({
  'sender': 'hadi',
  'receiver': 'hammurabi',
  'message': 'hello'
})
c.add({
  'sender': 'hammurabi',
  'receiver': 'hadi',
  'message': 'hello tooo'
})
