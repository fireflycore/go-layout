# Protovalidate Cheatsheet

Reference: [buf.validate](https://buf.build/bufbuild/protovalidate/docs/main:buf.validate)

## Common Constraints

### Scalars (Numbers)
```protobuf
uint32 age = 1 [(buf.validate.field).uint32.gt = 0];
float score = 2 [(buf.validate.field).float.gte = 0.0, (buf.validate.field).float.lte = 100.0];
```

### Strings
```protobuf
string name = 1 [(buf.validate.field).string.min_len = 1];
string email = 2 [(buf.validate.field).string.email = true];
string uuid = 3 [(buf.validate.field).string.uuid = true];
string pattern = 4 [(buf.validate.field).string.pattern = "^[a-zA-Z0-9]+$"];
```

### Enums / Values
```protobuf
MyEnum status = 1 [(buf.validate.field).enum.defined_only = true];
uint32 type = 2 [(buf.validate.field).uint32 = { in: [1, 2, 3] }];
```

### Required
```protobuf
string id = 1 [(buf.validate.field).required = true];
```

### Collections (Repeated)
```protobuf
repeated string tags = 1 [(buf.validate.field).repeated.min_items = 1];
repeated string unique_tags = 2 [(buf.validate.field).repeated.unique = true];
```

### Messages
```protobuf
MyMessage msg = 1 [(buf.validate.field).cel = {
  id: "my_message_check",
  message: "field a must be greater than field b",
  expression: "this.a > this.b"
}];
```
