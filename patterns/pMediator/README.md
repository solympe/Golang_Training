Realisation of <b>pattern Mediator</b>

This implementation has 2 interfaces: Airport(mediator) and Aeroflot(colleagues)

Airport structure is a mediator<br>

Aeroflot contains 2 structures: helicopter and plane<br>

Colleagues can notify airport about delaying of flight departure<br>
Airport can manipulate this information and delay departure time


You can test this pattern in main.go

P.S.<br>
main.go location: ./main.go<br>
packets location: ./patterns/pMediator/*