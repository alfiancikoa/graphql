# GraphQL Basic

Doc of graphql <a href="https://graphql.org"><b>here</b></a> 

## Apa itu GraphQL ?

<p align="justify"><b>GraphQL API</b> dikembangkan oleh Facebook, dengan tujuan untuk mendapatkan sebuah API yang lebih fleksibel dan efisien. Pendekatan ini dapat menyelesaikan permasalahan yang ada sekarang dan juga permasalahan yang dihadapi pengembang pada saat mengembangkan API dengan menggunakan pendekatan REST API. Meskipun demikian, mengembangkan sebuah API dengan pendekatan GraphQL API juga menambah daftar tantangan baru yang harus diselesaikan oleh para developer. <br><br>
GraphQL API menyediakan kemudahan dalam pengambilan data dan proses pengembangan yang lebih cepat. Meskipun begitu, GraphQL bukan merupakan pengganti dari REST API. Masing-masing pendekatan ini memiliki kelebihan dan kekurangannya masing-masing sesuai dengan kasus atau kebutuhan kita pada saat mengembangkan sistem. 
<br><br>
more info <a href="https://www.dicoding.com/blog/graphql-api-vs-rest-api-apa-bedanya"><b>here</b></a> 
<br><br>
(Sumber: https://www.dicoding.com/blog/graphql-api-vs-rest-api-apa-bedanya)
</p>

## Command To Use

### Set up Project
 
- Membuat file workspace

```
mkdir graphql-server
cd graphql-server
go mod init github.com/[username]/graphql-server
```

- Kemudian buat file ```tools.go``` dan tambahkan gqlgen sebagai dependensi alat untuk modul.

```
//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
)
```

- Download library gqlgen dari 99design

```
$ go get github.com/99designs/gqlgen
```

- 

Doc gqlgen library <a href="https://gqlgen.com/getting-started/"><b>here</b></a><br>
Github gqlgen Library <a href="https://github.com/99designs/gqlgen/"><b>here</b></a>

