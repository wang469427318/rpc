```protoc --go_out=plugins=grpc:../services Prod.proto```

```protoc --grpc-gateway_out=logtostderr=true:../services Prod.proto```



***CA证书生成***

1.openssl genrsa -out ca.key 2048

2.openssl req -new -x509 -days 3650 -key ca.key -out ca.pem

```
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) []:cn
State or Province Name (full name) []:beijing
Locality Name (eg, city) []:bejing
Organization Name (eg, company) []:andy
Organizational Unit Name (eg, section) []:andy
Common Name (eg, fully qualified host name) []:localhost
Email Address []:
```

生成服务端证书

1.openssl genrsa -out server.key 2048

2.openssl req -new -key server.key -out server.csr

```
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) []:cn
State or Province Name (full name) []:beijing
Locality Name (eg, city) []:bejing
Organization Name (eg, company) []:andy
Organizational Unit Name (eg, section) []:andy
Common Name (eg, fully qualified host name) []:localhost
Email Address []:

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:
```

CA证书验签Server证书

3.openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem

生成客户端证书

1.openssl ecparam -genkey -name secp384r1 -out client.key

2.openssl req -new -key client.key -out client.csr

```
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) []:cn
State or Province Name (full name) []:bejing
Locality Name (eg, city) []:beijing
Organization Name (eg, company) []:andy
Organizational Unit Name (eg, section) []:andy
Common Name (eg, fully qualified host name) []:localhost
Email Address []:

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:
```

CA证书验签Client证书

3.openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem