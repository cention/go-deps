If your packages no longer need any of these libraries please make an attempt to
clean it up by deleting the unused submodules after verifying that it's not
used by any other packages.

version: 138

---

notes/reference on rabbitmq/amqp091-go, *ymmv:
```
git clone git@github.com:cention/go-deps.git ~/gh/go-deps
cd ~/gh/go-deps
git submodule add git@github.com:rabbitmq/amqp091-go.git src/github.com/rabbitmq/amqp091-go
```
~/gh/go-deps/.gitmodules should auto-updated, should not modify this file at all.
```
[submodule "src/github.com/rabbitmq/amqp091-go"]
	path = src/github.com/rabbitmq/amqp091-go
	url = git@github.com:rabbitmq/amqp091-go.git
```
```
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        renamed:    src/github.com/rabbitmq/amqp091-go.git -> src/github.com/rabbitmq/amqp091-go
```
- output with `renamed:` above should signal OK, git commit and push from this point.
- check from https://github.com/cention/go-deps/tree/release-5.0/src/github.com/rabbitmq amqp091-go should points to the external git when clicked
- `git rm --cached src/github.com/rabbitmq/amqp091-go.git` was used to remove git related index, `git submodule status` was also used.
