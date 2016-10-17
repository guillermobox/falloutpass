# falloutpass

*Helps you to hack Fallout 3 terminals*

This go application helps you to hack Fallout 3 terminals. Just provide
the available words, and the application will suggest the best word for
each stage, and will guide you depending on the answers from the game.

An example of usage, if you use the file [testwords.txt](testwords.txt) provided with
this code. You can load them:

```
$ go run falloutpass.go testwords.txt

With 16 words left, I recommend you to chose HEARING

What word and score did you play (empty to see info)?
```

Now instead of playing the word `HEARING` and get a result from the game, you can
press enter and see some info:

```
    Word list [MENTION HUNTING LOOTING CRAZIES SOLDIER SHOWING STAMINA CONTENT SPARING SPOTTED STATUES SERVING HURTING FERTILE SYSTEMS HEARING]
    Optimal maxwords and strokeout are HEARING HEARING

    Distribution of surviving words
    MENTION  -   8   5   2   -   -   -   1
    HUNTING  -   5   2   5   2   -   1   1
    LOOTING  -   3   5   4   3   -   -   1
    CRAZIES  -  10   4   1   -   -   -   1
    SOLDIER  -   7   8   -   -   -   -   1
    SHOWING  -   6   2   4   3   -   -   1
    STAMINA  -   5   5   4   1   -   -   1
    CONTENT  -  10   3   2   -   -   -   1
    SPARING  -   4   4   3   3   1   -   1
    SPOTTED  1   8   5   1   -   -   -   1
    STATUES  -   9   2   4   -   -   -   1
    SERVING  -   5   2   4   4   -   -   1
    HURTING  -   5   3   3   3   -   1   1
    FERTILE  -   9   3   3   -   -   -   1
    SYSTEMS  1  11   2   1   -   -   -   1
    HEARING  2   3   3   3   3   1   -   1
```

The distribution of surviving words is: how many words will be left in the
list of words, if the choose to play a given word and the
game answers us with the score for that word. For example, if we play `MENTION`
the game can provide us with scores 1, 2, 3 and 7. For score 1, we have to
deal in the next step with 8 words, for score 2 we have to deal in the next
step with 5 words. For score 3 only 2 words. And, of course, if the score is 7,
the list of words is reduced to one, `MENTION`, because it has a score of 7/7.

The optimal word selected here is `HEARING` because it will give you only 3
words in the worse scenario.

Now we can keep playing:

```
What word and score did you play (empty to see info)? HEARING 2

With 3 words left, I recommend you to chose LOOTING

What word and score did you play (empty to see info)? 

    Word list [LOOTING SHOWING STAMINA]
    Optimal maxwords and strokeout are LOOTING LOOTING

    Distribution of surviving words
    LOOTING  -   -   1   -   1   -   -   1 
    SHOWING  -   -   -   1   1   -   -   1 
    STAMINA  -   -   1   1   -   -   -   1 
```

Only three words left after playing `HEARING`, which is the expected behaviour.
If, instead, we played something stupid like I did in the actual game:

```
What word and score did you play (empty to see info)? CONTENT 1

With 10 words left, I recommend you to chose HEARING

What word and score did you play (empty to see info)? 

    Word list [CRAZIES SOLDIER SHOWING STAMINA SPARING SPOTTED STATUES SERVING FERTILE HEARING]
    Optimal maxwords and strokeout are HEARING HEARING

    Distribution of surviving words
    CRAZIES  -   4   4   1   -   -   -   1 
    SOLDIER  -   2   7   -   -   -   -   1 
    SHOWING  -   3   2   2   2   -   -   1 
    STAMINA  -   2   2   4   1   -   -   1 
    SPARING  -   1   4   -   3   1   -   1 
    SPOTTED  1   4   3   1   -   -   -   1 
    STATUES  -   4   2   3   -   -   -   1 
    SERVING  -   3   1   2   3   -   -   1 
    FERTILE  -   7   1   1   -   -   -   1 
    HEARING  1   2   2   2   1   1   -   1 
```

I played `CONTENT`, and the score was 1, which reduced the list to a length of 10,
way worse than 3. After that point, falloutpass still recommends you to play
`HEARING`, but of course I played other thing: `SPARING`.

I couldn't open the stupid terminal.

But I had lockpicks, which is nice.
