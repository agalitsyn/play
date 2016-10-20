Git aggresively caches submodules, so at first we need a cleanup
```
git submodule deinit -f <name>
rm -rf <path>
git rm --cached <path>
rm -rf  .git/modules/<name>
```

Then add new module which tracks branch:
```
git submodule add -b master <url> <path>
```
And then commit changes.

On clean repo:

To add module run:
```
git submodule update --init
```

To update from remote run:
```
git submodule update --remote
```
And then commit changes.