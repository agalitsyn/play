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


with `GOTRACEBACK=crash`

```
2017/09/27 21:07:03 canceling
2017/09/27 21:07:03 close gen: context canceled
2017/09/27 21:07:03 close gen: context canceled
2017/09/27 21:07:03 close gen: context canceled
2017/09/27 21:07:03 close receiver: context canceled
fatal error: all goroutines are asleep - deadlock!

runtime stack:
runtime.throw(0x10d59c6, 0x25)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/panic.go:605 +0x95 fp=0xc420039e18 sp=0xc420039df8 pc=0x1027635
runtime.checkdead()
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:3800 +0x300 fp=0xc420039e68 sp=0xc420039e18 pc=0x1032b00
runtime.mput(0xc42002e380)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:4139 +0x54 fp=0xc420039e78 sp=0xc420039e68 pc=0x1033fb4
runtime.stopm()
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:1668 +0xba fp=0xc420039ea0 sp=0xc420039e78 pc=0x102ceda
runtime.findrunnable(0xc420021300, 0x0)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:2125 +0x4d2 fp=0xc420039f38 sp=0xc420039ea0 pc=0x102e0e2
runtime.schedule()
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:2245 +0x12c fp=0xc420039f80 sp=0xc420039f38 pc=0x102eb9c
runtime.park_m(0xc420000600)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:2308 +0xb6 fp=0xc420039fb8 sp=0xc420039f80 pc=0x102eeb6
runtime.mcall(0x0)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/asm_amd64.s:286 +0x5b fp=0xc420039fc8 sp=0xc420039fb8 pc=0x1050b1b

goroutine 1 [semacquire]:
runtime.gopark(0x10d7830, 0x115c4a0, 0x10d11b1, 0xa, 0xc420052019, 0x4)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:277 +0x12c fp=0xc42003fdb0 sp=0xc42003fd80 pc=0x102919c
runtime.goparkunlock(0x115c4a0, 0x10d11b1, 0xa, 0xc420050019, 0x4)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:283 +0x5e fp=0xc42003fdf0 sp=0xc42003fdb0 pc=0x102928e
runtime.semacquire1(0xc42001609c, 0x0, 0x1)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/sema.go:144 +0x1d4 fp=0xc42003fe60 sp=0xc42003fdf0 pc=0x1039854
sync.runtime_Semacquire(0xc42001609c)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/sema.go:56 +0x39 fp=0xc42003fe88 sp=0xc42003fe60 pc=0x1039479
sync.(*WaitGroup).Wait(0xc420016090)
	/usr/local/Cellar/go/1.9/libexec/src/sync/waitgroup.go:131 +0x72 fp=0xc42003feb0 sp=0xc42003fe88 pc=0x1062b02
main.main()
	/Users/a.galitsyn/src/github.com/agalitsyn/play/38-context-propagation/main.go:30 +0x2fe fp=0xc42003ff80 sp=0xc42003feb0 pc=0x109ca0e
runtime.main()
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:185 +0x20d fp=0xc42003ffe0 sp=0xc42003ff80 pc=0x1028d0d
runtime.goexit()
	/usr/local/Cellar/go/1.9/libexec/src/runtime/asm_amd64.s:2337 +0x1 fp=0xc42003ffe8 sp=0xc42003ffe0 pc=0x1053181

goroutine 2 [force gc (idle)]:
runtime.gopark(0x10d7830, 0x1150570, 0x10d1b6b, 0xf, 0x14, 0x1)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:277 +0x12c fp=0xc42002a768 sp=0xc42002a738 pc=0x102919c
runtime.goparkunlock(0x1150570, 0x10d1b6b, 0xf, 0xc420000114, 0x1)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:283 +0x5e fp=0xc42002a7a8 sp=0xc42002a768 pc=0x102928e
runtime.forcegchelper()
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:235 +0xcc fp=0xc42002a7e0 sp=0xc42002a7a8 pc=0x1028fbc
runtime.goexit()
	/usr/local/Cellar/go/1.9/libexec/src/runtime/asm_amd64.s:2337 +0x1 fp=0xc42002a7e8 sp=0xc42002a7e0 pc=0x1053181
created by runtime.init.4
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:224 +0x35

goroutine 3 [GC sweep wait]:
runtime.gopark(0x10d7830, 0x1150660, 0x10d175d, 0xd, 0x101c814, 0x1)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:277 +0x12c fp=0xc42002af60 sp=0xc42002af30 pc=0x102919c
runtime.goparkunlock(0x1150660, 0x10d175d, 0xd, 0x14, 0x1)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:283 +0x5e fp=0xc42002afa0 sp=0xc42002af60 pc=0x102928e
runtime.bgsweep(0xc42001e070)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/mgcsweep.go:52 +0xa3 fp=0xc42002afd8 sp=0xc42002afa0 pc=0x101c8f3
runtime.goexit()
	/usr/local/Cellar/go/1.9/libexec/src/runtime/asm_amd64.s:2337 +0x1 fp=0xc42002afe0 sp=0xc42002afd8 pc=0x1053181
created by runtime.gcenable
	/usr/local/Cellar/go/1.9/libexec/src/runtime/mgc.go:216 +0x58

goroutine 4 [finalizer wait]:
runtime.gopark(0x10d7830, 0x116e8e0, 0x10d1993, 0xe, 0x14, 0x1)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:277 +0x12c fp=0xc42002b700 sp=0xc42002b6d0 pc=0x102919c
runtime.goparkunlock(0x116e8e0, 0x10d1993, 0xe, 0x14, 0x1)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:283 +0x5e fp=0xc42002b740 sp=0xc42002b700 pc=0x102928e
runtime.runfinq()
	/usr/local/Cellar/go/1.9/libexec/src/runtime/mfinal.go:175 +0xb8 fp=0xc42002b7e0 sp=0xc42002b740 pc=0x1013968
runtime.goexit()
	/usr/local/Cellar/go/1.9/libexec/src/runtime/asm_amd64.s:2337 +0x1 fp=0xc42002b7e8 sp=0xc42002b7e0 pc=0x1053181
created by runtime.createfing
	/usr/local/Cellar/go/1.9/libexec/src/runtime/mfinal.go:156 +0x62

goroutine 9 [chan receive]:
runtime.gopark(0x10d7830, 0xc420070178, 0x10d15cd, 0xc, 0xc420040f17, 0x3)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:277 +0x12c fp=0xc420040dc0 sp=0xc420040d90 pc=0x102919c
runtime.goparkunlock(0xc420070178, 0x10d15cd, 0xc, 0x17, 0x3)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:283 +0x5e fp=0xc420040e00 sp=0xc420040dc0 pc=0x102928e
runtime.chanrecv(0xc420070120, 0xc420040f10, 0xc42002de01, 0xc4200500d0)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/chan.go:506 +0x304 fp=0xc420040eb0 sp=0xc420040e00 pc=0x10046e4
runtime.chanrecv2(0xc420070120, 0xc420040f10, 0x0)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/chan.go:393 +0x2b fp=0xc420040ee0 sp=0xc420040eb0 pc=0x10043cb
main.merge.func1(0xc420070120)
	/Users/a.galitsyn/src/github.com/agalitsyn/play/38-context-propagation/main.go:58 +0x90 fp=0xc420040fd8 sp=0xc420040ee0 pc=0x109cf80
runtime.goexit()
	/usr/local/Cellar/go/1.9/libexec/src/runtime/asm_amd64.s:2337 +0x1 fp=0xc420040fe0 sp=0xc420040fd8 pc=0x1053181
created by main.merge
	/Users/a.galitsyn/src/github.com/agalitsyn/play/38-context-propagation/main.go:70 +0x11a

goroutine 10 [chan receive]:
runtime.gopark(0x10d7830, 0xc4200701d8, 0x10d15cd, 0xc, 0xc420041f17, 0x3)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:277 +0x12c fp=0xc420041dc0 sp=0xc420041d90 pc=0x102919c
runtime.goparkunlock(0xc4200701d8, 0x10d15cd, 0xc, 0x17, 0x3)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:283 +0x5e fp=0xc420041e00 sp=0xc420041dc0 pc=0x102928e
runtime.chanrecv(0xc420070180, 0xc420041f10, 0xc420026601, 0xc4200500d0)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/chan.go:506 +0x304 fp=0xc420041eb0 sp=0xc420041e00 pc=0x10046e4
runtime.chanrecv2(0xc420070180, 0xc420041f10, 0x0)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/chan.go:393 +0x2b fp=0xc420041ee0 sp=0xc420041eb0 pc=0x10043cb
main.merge.func1(0xc420070180)
	/Users/a.galitsyn/src/github.com/agalitsyn/play/38-context-propagation/main.go:58 +0x90 fp=0xc420041fd8 sp=0xc420041ee0 pc=0x109cf80
runtime.goexit()
	/usr/local/Cellar/go/1.9/libexec/src/runtime/asm_amd64.s:2337 +0x1 fp=0xc420041fe0 sp=0xc420041fd8 pc=0x1053181
created by main.merge
	/Users/a.galitsyn/src/github.com/agalitsyn/play/38-context-propagation/main.go:70 +0x11a

goroutine 11 [semacquire]:
runtime.gopark(0x10d7830, 0x115c520, 0x10d11b1, 0xa, 0xc420052019, 0x4)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:277 +0x12c fp=0xc420026ea0 sp=0xc420026e70 pc=0x102919c
runtime.goparkunlock(0x115c520, 0x10d11b1, 0xa, 0xc420001919, 0x4)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/proc.go:283 +0x5e fp=0xc420026ee0 sp=0xc420026ea0 pc=0x102928e
runtime.semacquire1(0xc4200160ac, 0xc400000000, 0x1)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/sema.go:144 +0x1d4 fp=0xc420026f50 sp=0xc420026ee0 pc=0x1039854
sync.runtime_Semacquire(0xc4200160ac)
	/usr/local/Cellar/go/1.9/libexec/src/runtime/sema.go:56 +0x39 fp=0xc420026f78 sp=0xc420026f50 pc=0x1039479
sync.(*WaitGroup).Wait(0xc4200160a0)
	/usr/local/Cellar/go/1.9/libexec/src/sync/waitgroup.go:131 +0x72 fp=0xc420026fa0 sp=0xc420026f78 pc=0x1062b02
main.merge.func2(0xc420016090, 0xc4200160a0, 0xc4200701e0)
	/Users/a.galitsyn/src/github.com/agalitsyn/play/38-context-propagation/main.go:76 +0x51 fp=0xc420026fc8 sp=0xc420026fa0 pc=0x109d181
runtime.goexit()
	/usr/local/Cellar/go/1.9/libexec/src/runtime/asm_amd64.s:2337 +0x1 fp=0xc420026fd0 sp=0xc420026fc8 pc=0x1053181
created by main.merge
	/Users/a.galitsyn/src/github.com/agalitsyn/play/38-context-propagation/main.go:73 +0x16e
```