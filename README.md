# prompt
A binary based prompt command to do all the awesome things


Ok so idea is that we do 

PROMPT_PID=$$
PS0='$(./prompt before '${PROMPT_PID}')'
PROMPT_COMMAND='_r=$?;./prompt fix;PS1=$(./prompt after '${PROMPT_PID}' "$_r")'

MUST set PS1 to the prompt string otherwise it all breaks

https://superuser.com/questions/175799/does-bash-have-a-hook-that-is-run-before-executing-a-command
http://stromberg.dnsalias.org/~strombrg/PS0-prompt/
export PS0='$(date)'
echo hi[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt >
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > aw
Sat 25 Aug 2018 23:07:04 BSTbash: aw: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > awd
Sat 25 Aug 2018 23:07:06 BSTbash: awd: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > ad
Sat 25 Aug 2018 23:07:06 BSTbash: ad: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > awd
Sat 25 Aug 2018 23:07:07 BSTbash: awd: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > awd
Sat 25 Aug 2018 23:07:08 BSTbash: awd: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > awd
Sat 25 Aug 2018 23:07:08 BSTbash: awd: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > awd
Sat 25 Aug 2018 23:07:09 BSTbash: awd: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > awd
Sat 25 Aug 2018 23:07:09 BSTbash: awd: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > aw
dSat 25 Aug 2018 23:07:10 BSTbash: aw: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > dawd
Sat 25 Aug 2018 23:07:10 BSTbash: dawd: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt >
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt >
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > awd
Sat 25 Aug 2018 23:07:12 BSTbash: awd: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > awd
Sat 25 Aug 2018 23:07:12 BSTbash: awd: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > awd
Sat 25 Aug 2018 23:07:13 BSTbash: awd: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > awd
Sat 25 Aug 2018 23:07:13 BSTbash: awd: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > awd
Sat 25 Aug 2018 23:07:14 BSTbash: awd: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > awd
Sat 25 Aug 2018 23:07:14 BSTbash: awd: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > awd
Sat 25 Aug 2018 23:07:15 BSTbash: awd: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > awd
Sat 25 Aug 2018 23:07:16 BSTbash: awd: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > aw
Sat 25 Aug 2018 23:07:17 BSTbash: aw: command not found
[bmeier@bmeier-mac ~/work/go/src/github.com/AstromechZA/prompt > export PS1='$(date)'
Sat 25 Aug 2018 23:07:28 BSTSat 25 Aug 2018 23:07:28 BST
Sat 25 Aug 2018 23:07:29 BST
Sat 25 Aug 2018 23:07:30 BST
Sat 25 Aug 2018 23:07:30 BST
Sat 25 Aug 2018 23:07:30 BST
Sat 25 Aug 2018 23:07:31 BST
Sat 25 Aug 2018 23:07:31 BST
Sat 25 Aug 2018 23:07:31 BST
Sat 25 Aug 2018 23:07:31 BSTsleep 10
Sat 25 Aug 2018 23:07:33 BSTSat 25 Aug 2018 23:07:43 BST
