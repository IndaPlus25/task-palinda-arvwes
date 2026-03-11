##What happens if you remove the go-command from the Seek call in the main function?##
The program will run sequentially so the same people will recieve and send the same messages every run. 

##What happens if you switch the declaration wg := new(sync.WaitGroup) to var wg sync.WaitGroup and the parameter wg *sync.WaitGroup to wg sync.WaitGroup?
using var wg sync.WaitGroup will result in the seek making a copy of the counter so it will not increment the counter in the main function resulting in a dead lock

##What happens if you remove the buffer on the channel match?
if the buffer is removed the program will deadlock at the 5fth person since it will get stuck for ever in the select claus, since there is no deafult and no available option due to an odd number of routines. 

What happens if you remove the default-case from the case-statement in the main function?
Removing the deafult will result in the program deadlocking ifd there is an even numebr of people, since main gets blocked in the SElect case when the chanel is empty

Hint: Think about the order of the instructions and what happens with arrays of different lengths.