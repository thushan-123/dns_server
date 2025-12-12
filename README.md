# Simple DNS Server

## Test DNS using Dns Lookup
```
    dig @127.0.0.1 -p 1234 google.com

```

## Dig Dns Lookup output

```
    ; <<>> DiG 9.18.39-0ubuntu0.24.04.2-Ubuntu <<>> @127.0.0.1 -p 1234 google.com
    ; (1 server found)
    ;; global options: +cmd
    ;; Got answer:
    ;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 53875
    ;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0

    ;; QUESTION SECTION:
    ;google.com.			IN	A

    ;; ANSWER SECTION:
    google.com.		60	IN	A	127.0.0.1

    ;; Query time: 0 msec
    ;; SERVER: 127.0.0.1#1234(127.0.0.1) (UDP)
    ;; WHEN: Fri Dec 12 06:32:46 +0530 2025
    ;; MSG SIZE  rcvd: 54


```