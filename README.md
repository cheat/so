so
==
`so` is a companion binary to [vim-so][]. It queries the Stack Overflow API, and makes
code snippets  available to `vim-so`.


Installing
----------
Download the newest release from [Releases][] and manually install it on your
system `PATH`.


FAQ
---
### Why is this executable necessary? ###
`vim-so` must query the Stack Overflow API for code snippets, yet `vim`
provides no native mechanism for making HTTP requests. This executable bridges
that gap.

### Why is this executable not bundled with vim-so? ###
This executable is approximately 6Mb in size. In order to support multiple
platforms (currently six in total), `vim-so` would need to contain a compiled
executable for each, and I do not wish to release a ~40Mb `vim` plugin. 

### Why was this not written in Python? ###
I have had significant difficulty in the past maintaining Python software that
is distributed among a wide user-base. I'm sidestepping that problem by
distributing an executable.

### Can I use this to "manually" search Stack Overflow on the command-line? ###
Not in a convenient way. The executable only outputs JSON.

### Why not? ###
Initially, I intended for `so` to support that use-case. However, I backed away
from that goal upon realizing two things:

1. It is difficult to display a Stack Overflow thread on the command line with
	 a user-experience that is equivalent to the browser.  The browser provides a
	 superior experience.

2. Signficantly: I've found that search engines do a better job of returning
	 relevant Stack Overflow threads than Stack Overflow's own API. If the API
	 doesn't return an adequate response, your preferred search engine will
	 likely do a better job.

With that being said, if you're interested in doing this anyway, you may wish
to check out the [how2][] project.


### So is this tool useless? ###
I don't think so. It's imperfect, but seems to get the job done about 50% of
the time. And when it works, it feels a bit magical.

### What else should I know? ###
You shouldn't waste your time asking Stack Overflow the same question more than
once. After you've gotten an answer from Stack Overflow - either via `vim-so`
or your web-browser - you should cache it locally using [cheat][] and
[vim-cheat][]. <!-- TODO: link to blog post explaining tool ecosystem -->


[Releases]:  https://github.com/cheat/so/releases
[cheat]:     https://github.com/cheat/cheat
[how2]:      https://github.com/santinic/how2
[vim-cheat]: https://github.com/cheat/vim-cheat
[vim-so]:    https://github.com/cheat/vim-so/releases
