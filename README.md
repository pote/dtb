


```
while true; do sleep $(( $RANDOM / 100 )); git commit --allow-empty -m "$(curl http://drinking-typing-bird.herokuapp.com/)" ; done
```
