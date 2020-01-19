<b>Pattern Observer</b>

The Observer pattern refers to behavioral patterns at the object level.

The Observer pattern defines a one-to-many relationship between objects so that when the state of one object changes, all objects that depend on it are notified and updated automatically.

The main participants in the pattern are Publishers (newsPortal) and Subscribers (subscribers).

NewsPortal can add and delete new subscribers.<br>
All subscribers can be notified of a new event on the news portal<br>
If subscriber receive the message he will talks about it (prints in the console)