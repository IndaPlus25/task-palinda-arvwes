If written answers are required, you can add them to this file. Just copy the
relevant questions from the root of the repo, preferably in
[Markdown](https://guides.github.com/features/mastering-markdown/) format :)

#### Task 1

##### Buggy Code 1
1. What is wrong:
2. How it was fixed:

##### Buggy Code 2
1. What is wrong:
2. How it was fixed:

#### Task 2

| Question | What I expected | What happened | Why I believe this happened |
|-|-|-|-|
| What happens if you do X? |  Program would still work as before | Program ended up in a deadlock | Because of reasons 🤷 |
| What happens if you switch the order of the statements `wgp.Wait()` and `close(ch)` in the end of the `main` function? | program would panic | program paniced | tried to send to closed chanel |
| What happens if you move the `close(ch)` from the `main` function and instead close the channel in the end of the function `Produce`?  | panic due to reading to closed chanel | program paniced | first producer to finish its task closed the chanel resulting in other producers sending to closed chanel |
| What happens if you remove the statement `close(ch)` completely?  | Program would still work as before | program worked as before | since nothing is waiting for the chanel to close the program moves on and finishes as normal |
| What happens if you increase the number of consumers from 2 to 4?  | the program should finish ~about twice as fast | the program finished about twice as fast | since now there were 4 threads running simultaniously consuming the 32 String |
| Can you be sure that all strings are printed before the program stops?  | No you can not be since the program does not wait for the string to be consumed | last string was not printed if random consume time was added before printing | because the program does not wait for the string to be consumed so if it reaches close before then it will stop the for range ch loop|
