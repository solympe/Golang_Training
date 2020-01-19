<b>Pattern Memenro</b>

The Memento pattern refers to behavioral patterns of the object level.

The Memento template receives and retains its internal state outside the object,<br>
so that later it can be restored to the same state. If the client subsequently needs<br>
to "roll back" the state of the original object, it passes Memento back to the original <br>
object in order to restore it.
  
This pattern has 2 structures : palette and history

The palette contains color, the history contains a stack of previous colors that were previously used.<br>
