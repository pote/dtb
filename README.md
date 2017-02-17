# Drinking Typing Bird

![I don't think pepole look at tooltips anymore](https://media.giphy.com/media/l41lUJ1YoZB1lHVPG/giphy-facebook_s.jpg)

As we contemplate the meaningless existance that we lead fuelling a capitalist system based on on our labour we sometimes wish there was a way to make it look like we are actually earning our salary.

In an amazing display of selflessness two developers decided to sacrifice about 50 entire minutes to this righteous cause, this is the result.


### Usage

We've deployed this app to Heroku, because we really wanted to spend the lease possible amount of effort.

```bash
$ curl http://drinking-typing-bird.herokuapp.com
```

### REAL usage

Jump into whatever repository you're working on and run this command, it will sleep randomly and create an empty commit using a phrase from [whatthecommit.com](whatthecommit.com) as the commit message. 

```
while true; do sleep $(( $RANDOM / 100 )); git commit --allow-empty -m "$(curl http://drinking-typing-bird.herokuapp.com/)" ; done
```
### "I REGRET NOTHING" usage

Add a `git push origin master` to the one-liner above, if you dare.


## Contribution guidelines

We accept only commit messages created drinking typing bird,  or some other method that takes even less effort.
