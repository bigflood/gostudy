# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ping.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='ping.proto',
  package='pb',
  syntax='proto3',
  serialized_pb=_b('\n\nping.proto\x12\x02pb\"\x16\n\x07PingReq\x12\x0b\n\x03msg\x18\x01 \x01(\t\"2\n\x07PingRes\x12\x0b\n\x03msg\x18\x01 \x01(\t\x12\x0b\n\x03log\x18\x02 \x01(\t\x12\r\n\x05\x65rror\x18\x03 \x01(\t2*\n\x04Ping\x12\"\n\x04Ping\x12\x0b.pb.PingReq\x1a\x0b.pb.PingRes\"\x00\x62\x06proto3')
)




_PINGREQ = _descriptor.Descriptor(
  name='PingReq',
  full_name='pb.PingReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='msg', full_name='pb.PingReq.msg', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=18,
  serialized_end=40,
)


_PINGRES = _descriptor.Descriptor(
  name='PingRes',
  full_name='pb.PingRes',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='msg', full_name='pb.PingRes.msg', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='log', full_name='pb.PingRes.log', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error', full_name='pb.PingRes.error', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=42,
  serialized_end=92,
)

DESCRIPTOR.message_types_by_name['PingReq'] = _PINGREQ
DESCRIPTOR.message_types_by_name['PingRes'] = _PINGRES
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

PingReq = _reflection.GeneratedProtocolMessageType('PingReq', (_message.Message,), dict(
  DESCRIPTOR = _PINGREQ,
  __module__ = 'ping_pb2'
  # @@protoc_insertion_point(class_scope:pb.PingReq)
  ))
_sym_db.RegisterMessage(PingReq)

PingRes = _reflection.GeneratedProtocolMessageType('PingRes', (_message.Message,), dict(
  DESCRIPTOR = _PINGRES,
  __module__ = 'ping_pb2'
  # @@protoc_insertion_point(class_scope:pb.PingRes)
  ))
_sym_db.RegisterMessage(PingRes)



_PING = _descriptor.ServiceDescriptor(
  name='Ping',
  full_name='pb.Ping',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=94,
  serialized_end=136,
  methods=[
  _descriptor.MethodDescriptor(
    name='Ping',
    full_name='pb.Ping.Ping',
    index=0,
    containing_service=None,
    input_type=_PINGREQ,
    output_type=_PINGRES,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_PING)

DESCRIPTOR.services_by_name['Ping'] = _PING

# @@protoc_insertion_point(module_scope)
