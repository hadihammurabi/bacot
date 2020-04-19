import json
from .block import Block

class Chain:
  def __init__(self, name='', difficulty={'char': '0', 'times': 4}):
    self.name = name
    self.difficulty = difficulty
    self.blocks = [self.createGenesis()]

  def to_dict(self):
    return {
      'name': self.name,
      'blocks': list(map(lambda x: x.to_dict(), self.blocks)),
    }
  
  def __str__(self):
    return json.dumps(self.to_dict())

  def createGenesis(self):
    return Block(0, '', 'genesis')

  def add(self, data={}):
    index = len(self.blocks)
    prev_hash = self.blocks[index - 1].meta.hash
    block = Block(index, prev_hash, data)
    block.proof_of_work(self.difficulty)
    self.blocks.append(block)
    return block

  def peek(self):
    return self.blocks[len(self.blocks) - 1]

  def is_valid(self):
    for index, block in enumerate(self.blocks):
      before = self.blocks[index - 1] if index > 0 else None
      if block.meta.hash != block.create_hash(block.meta.nonce): return False
      if before != None and block.meta.prev_hash != before.meta.hash: return False

    return True
