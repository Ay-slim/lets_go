
- Go routines
  These are the units that are used to run go code (this had me thinking of processes and threads)
  To make code concurrent, use the go key word before the piece of code to spin up a new go routine to run the code separately
  Just using go in front of a function within loops breaks because a default routine is created when you're running code regularly, so once the main routine finishes it's thing, it just exits like "I'm done"
  We work around this using channels which are the only way go routines can communicate with each other
  We can write to channels channel <- message or read from channels fmt.Println(<- channel)
  We can also have a loop read from a channel everytime it gets a message and do something with it
  for l := range(channel) {

  }
  Reading from a channel is a blocking activity