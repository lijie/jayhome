## GDB

### Start

gdb program <PID>

gdb --args program <program args...>

### Checkpoint/Restart

checkpoint
> Save a snapshot of the debugged pgrograms's current execution state.

restart checkpoint
> Restore the program state that was saved as checkpoint nmber checkpoint-id.

### Reverse Execution

reverse-continue\rc

reverse-step

reverse-next

### Process record and reply

record full

record save filename

record restore filename

### Frame

frame N/up/down/info frame/info locals

### Source

### watchpoints

### catchpoints





