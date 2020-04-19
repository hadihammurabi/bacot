import json
from flask import Flask, Response, request

from src.core.chain import Chain

chain = Chain('bacot')
web = Flask('bacot')

@web.route('/')
def index():
  return Response(json.dumps({
    'message': 'welcome to bacot blockchain engine'
  }), content_type='application/json')

@web.route('/chain', methods=['GET'])
def get_chain():
  return Response(str(chain), content_type='application/json')

@web.route('/chain', methods=['POST'])
def post_chain():
  data = request.json
  if not data:
    return Response(json.dumps({ message: 'no data' }), content_type='application/json')

  block = chain.add(data)
  return Response(str(block), content_type='application/json')
