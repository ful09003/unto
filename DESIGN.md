# WTF is this?

command hacking for fun and ~~profit~~ memory, because I have almost none..

So, the basic idea was "be able to run commands and persist not only the command, but env vars and stdout/err to a file that I could go back to later"
(the nice thing about that is, I'd then love to roll the files into something like a blog or how-to or whatever)

Funnily enough, right after I began doing this, I saw a post online about https://www.warp.dev/ which looks like it might
fill the need I have! I'll have to play with it.

Should warp not pan out, I'll be back to do more with this.. Right now, the obvious glaring problem
is that it is _not_ necessarily a drop-in replacement for a terminal.

Piping will fail, because the implementation is not really spawning a real shell, it just parses extra args and executes them.

That can be worked around ala `./unto -savedir=$(pwd)/.logs bash -c 'ping -c1 1.1.1.1 | grep "bytes from"'`

but to do what I'd love, this program would instead need to likely do some form of teeing PTY (using https://github.com/creack/pty perhaps)

Anyways, this was a fun little thought experiment and something I really do long for, but am not going to devote time to
if the alternative above (warp) ends up checking all of my boxes :)