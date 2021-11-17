# creating rabbit/rabbyte meme with every possible word which ends with bit using golang and some other stuff.

### if you are using macOS or linux, you can find all the words on your system at this exact path: /usr/share/dict/words
```
wc -l /usr/share/dict/words
OUT: 235886 /usr/share/dict/words
```
### will show you how many words are on that file, the next thing we want to do is find the words which ends with word "bit"
```
grep "bit$" /usr/share/dict/words | wc -l 
OUT: 46
```
### means you got 46 words which end with "bit". now lets search them on google and create our memes :D


![This is an image](https://img.devrant.com/devrant/rant/r_1794865_ucb6w.jpg)