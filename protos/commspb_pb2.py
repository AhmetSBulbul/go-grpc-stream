# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/commspb.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14protos/commspb.proto\x12\x07\x63ommspb\"\x1d\n\x07\x43hannel\x12\x12\n\nchannel_id\x18\x01 \x01(\t\"=\n\x07Message\x12!\n\x07\x63hannel\x18\x01 \x01(\x0b\x32\x10.commspb.Channel\x12\x0f\n\x07message\x18\x02 \x01(\t\"\x1c\n\nMessageAck\x12\x0e\n\x06status\x18\x01 \x01(\t2{\n\x0c\x43ommsService\x12\x33\n\x0bJoinChannel\x12\x10.commspb.Channel\x1a\x10.commspb.Message0\x01\x12\x36\n\x0bSendMessage\x12\x10.commspb.Message\x1a\x13.commspb.MessageAck(\x01\x42/Z-github.com/AhmetSBulbul/go-grpc-stream/protosb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'protos.commspb_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z-github.com/AhmetSBulbul/go-grpc-stream/protos'
  _CHANNEL._serialized_start=33
  _CHANNEL._serialized_end=62
  _MESSAGE._serialized_start=64
  _MESSAGE._serialized_end=125
  _MESSAGEACK._serialized_start=127
  _MESSAGEACK._serialized_end=155
  _COMMSSERVICE._serialized_start=157
  _COMMSSERVICE._serialized_end=280
# @@protoc_insertion_point(module_scope)