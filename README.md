# Pomodoro project

This is a simple pomodoro project written in Go for my personal goals :).
Feel free to use it!

# Pomodoro Technique

The *Pomodoro Technique* is a time management method developed by Francesco
Cirillo in the late 1980s. The technique uses a timer to break down work into
intervals, traditionally *25 minutes* in length, separated by short breaks.

# Usage

*There are six steps in the original technique:*

1. Decide on the task to be done.

2. Set the pomodoro timer (traditionally for 25 minutes).

3. Work on the task.

4. End work when the timer rings and take a short break (traditionally 5 to 10
   minutes).

5. If you have fewer than three pomodoros, go back to Step 2 and repeat until
   you go through all three pomodoros.

6. After three pomodoros are done, take a long break (traditionally 20 to 30
   minutes). Once the long break is finished, repeat to step 2.


*After task completion in a pomodoro, any time remaining should be devoted to
activities for instance:*

1. Review your work just completed.

2. Review the activities from a learning point of view (ex: What learning
	 objective did you accomplish? What learning outcome did you accomplish? Did you
	 fulfill your learning target, objective, or outcome for the task?)

3. Review the list of upcoming tasks for the next planned pomodoro time blocks,
	 and start reflecting on or updating those tasks.

# Installation

1. With `go get`
```
$ go get github.com/11me/pomo
```
2. Or clone the repo.
```
$ git clone https://github.com/11me/pomo

$ cd pomo

$ go install ...
```


# Pomo usage

It is assumed that pomo will be used in a status bar like i3, dwmblocks, etc.

1. Download the binary files and put them in your `$PATH`.

2. Start the `pomod` daemon.
   For example:
```
$ pomod &
```
3. Interact with `pomo` client.
```
$ pomo -start
```
4. Get the duration.
```
$ pomo -get-duration
```
5. Set a new duration.
```
$ pomo -set-duration 45m
```
