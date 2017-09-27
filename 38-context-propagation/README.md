Strage error:

```
$ for i in {1..10}; do go run main.go; done

0
1
0
0
2
1
1
3
4
2
2
3
4
3
5
5
6
4
6
5
6
7
7
8
8
7
9
9
8
10
2017/09/27 15:36:48 canceling
2017/09/27 15:36:48 close gen: context canceled
2017/09/27 15:36:48 close gen: context canceled
2017/09/27 15:36:48 close receiver: context canceled
2017/09/27 15:36:48 close gen: context canceled
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc42001609c)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/sema.go:56 +0x39
sync.(*WaitGroup).Wait(0xc420016090)
	/usr/local/Cellar/go/1.9/libexec/src/sync/waitgroup.go:131 +0x72
main.main()
	/Users/a.galitsyn/src/github.com/agalitsyn/play/38-context-propagation/main.go:29 +0x2fe

goroutine 9 [chan receive]:
main.merge.func1(0xc420070120)
	/Users/a.galitsyn/src/github.com/agalitsyn/play/38-context-propagation/main.go:57 +0x90
created by main.merge
	/Users/a.galitsyn/src/github.com/agalitsyn/play/38-context-propagation/main.go:69 +0x11a

goroutine 10 [chan receive]:
main.merge.func1(0xc420070180)
	/Users/a.galitsyn/src/github.com/agalitsyn/play/38-context-propagation/main.go:57 +0x90
created by main.merge
	/Users/a.galitsyn/src/github.com/agalitsyn/play/38-context-propagation/main.go:69 +0x11a

goroutine 11 [semacquire]:
sync.runtime_Semacquire(0xc4200160ac)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/sema.go:56 +0x39
sync.(*WaitGroup).Wait(0xc4200160a0)
	/usr/local/Cellar/go/1.9/libexec/src/sync/waitgroup.go:131 +0x72
main.merge.func2(0xc420016090, 0xc4200160a0, 0xc4200701e0)
	/Users/a.galitsyn/src/github.com/agalitsyn/play/38-context-propagation/main.go:75 +0x51
created by main.merge
	/Users/a.galitsyn/src/github.com/agalitsyn/play/38-context-propagation/main.go:72 +0x16e
exit status 2
```