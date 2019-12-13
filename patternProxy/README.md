<b>Realisation of Proxy pattern</b>

This implementation has 2 structures: dataBase(main struct) and dataBaseNode(proxy struct)

Both structures are connected by one 'DBFunctions' interface

Both structures have two similar <i>methods</i>:<br>
SendData() - sending new data to dataBase<br>
GetData() - getting data from dataBase

You can test this pattern in main.go

P.S.<br>
main.go location: ./main.go<br>
packets location: ./patternProxy/* 
