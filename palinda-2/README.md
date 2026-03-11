If written answers are required, you can add them to this file. Just copy the
relevant questions from the root of the repo, preferably in
[Markdown](https://guides.github.com/features/mastering-markdown/) format :)

#### Task 1

##### Buggy Code 1
1. What is wrong:
does not work due to chanel never being able to let go of the value
causing the program to freeze on the ch <- "Hello world!" line
2. How it was fixed:
creating a buffered chanel fixes it since the ch can let go of the value and place it in the buffer (as long as the buffer is not full) and move on and print


##### Buggy Code 2
1. What is wrong: this program does not count the final value because main closes before the print function has time to print. 
2. How it was fixed:
Fixed by cereating a wait group so main waits for the go print thread to finish.
#### Task 2

| Question | What I expected | What happened | Why I believe this happened |
|-|-|-|-|
| What happens if you do X? |  Program would still work as before | Program ended up in a deadlock | Because of reasons 🤷 |
| What happens if you switch the order of the statements `wgp.Wait()` and `close(ch)` in the end of the `main` function? | program would panic | program paniced | tried to send to closed chanel |
| What happens if you move the `close(ch)` from the `main` function and instead close the channel in the end of the function `Produce`?  | panic due to reading to closed chanel | program paniced | first producer to finish its task closed the chanel resulting in other producers sending to closed chanel |
| What happens if you remove the statement `close(ch)` completely?  | Program would still work as before | program worked as before | since nothing is waiting for the chanel to close the program moves on and finishes as normal |
| What happens if you increase the number of consumers from 2 to 4?  | the program should finish ~about twice as fast | the program finished about twice as fast | since now there were 4 threads running simultaniously consuming the 32 String |
| Can you be sure that all strings are printed before the program stops?  | No you can not be since the program does not wait for the string to be consumed | last string was not printed if random consume time was added before printing | because the program does not wait for the string to be consumed so if it reaches close before then it will stop the for range ch loop|
