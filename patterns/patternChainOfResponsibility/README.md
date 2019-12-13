<b>Pattern chain of responsibility</b>

This implementation has 3 main structures:<br>
DeliveryCourier - 1st handler<br>
DeliveryMail - 2nd handler<br>
DeliveryPlane - 3nd handler

All this structures are connected by one 'TypeOfDelivery' interface

Each of structures has one similar <i>method</i> - ChooseType(chosen)<br>
'ChooseType' checks entered word and returns result according to the condition

You can test this pattern in main.go

P.S.<br>
main.go location: ./main.go<br>
packet location: .patterns/patternChainOfResponsibility/* 
