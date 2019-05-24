package ctxwait

import (
	"context"
)

type ctxCallback struct {
	done <-chan struct{}
	cb   func()
}

var inboxCh = make(chan ctxCallback)

// OnDone calls the callback in a goroutine when the context is done.
func OnDone(ctx context.Context, cb func()) {
	if ctx.Err() != nil {
		go cb()
		return
	}
	ch := ctx.Done()
	if ch == nil {
		// will never fire
		return
	}
	c := ctxCallback{ch, cb}
	select {
	case inboxCh <- c:
	default:
		go worker(c)
	}
}

// OnDoneWithCancel is like OnDone but provides a way to unregister the callback
// without triggering it.
func OnDoneWithCancel(ctx context.Context, cb func()) context.CancelFunc {
	parentDone := ctx.Done()
	ctx, cancel := context.WithCancel(ctx)
	OnDone(ctx, func() {
		cancel()
		select {
		case <-parentDone:
			cb()
		default:
			// canceled.
		}
	})
	return cancel
}

func worker(initial ctxCallback) {
	var callbacks [32]ctxCallback
	callbacks[0] = initial
	length := 1

	in := inboxCh
	for length > 0 {
		var selected int
		select {
		case <-callbacks[0].done:
			selected = 0
		case <-callbacks[1].done:
			selected = 1
		case <-callbacks[2].done:
			selected = 2
		case <-callbacks[3].done:
			selected = 3
		case <-callbacks[4].done:
			selected = 4
		case <-callbacks[5].done:
			selected = 5
		case <-callbacks[6].done:
			selected = 6
		case <-callbacks[7].done:
			selected = 7
		case <-callbacks[8].done:
			selected = 8
		case <-callbacks[9].done:
			selected = 9
		case <-callbacks[10].done:
			selected = 10
		case <-callbacks[11].done:
			selected = 11
		case <-callbacks[12].done:
			selected = 12
		case <-callbacks[13].done:
			selected = 13
		case <-callbacks[14].done:
			selected = 14
		case <-callbacks[15].done:
			selected = 15
		case <-callbacks[16].done:
			selected = 16
		case <-callbacks[17].done:
			selected = 17
		case <-callbacks[18].done:
			selected = 18
		case <-callbacks[19].done:
			selected = 19
		case <-callbacks[20].done:
			selected = 20
		case <-callbacks[21].done:
			selected = 21
		case <-callbacks[22].done:
			selected = 22
		case <-callbacks[23].done:
			selected = 23
		case <-callbacks[24].done:
			selected = 24
		case <-callbacks[25].done:
			selected = 25
		case <-callbacks[26].done:
			selected = 26
		case <-callbacks[27].done:
			selected = 27
		case <-callbacks[28].done:
			selected = 28
		case <-callbacks[29].done:
			selected = 29
		case <-callbacks[30].done:
			selected = 30
		case <-callbacks[31].done:
			selected = 31
		case cb := <-in:
			callbacks[length] = cb
			length++
			if length == len(callbacks) {
				in = nil
			}
			continue
		}
		cb := callbacks[selected].cb
		callbacks[selected] = callbacks[length-1]
		callbacks[length-1] = ctxCallback{}
		length--

		go cb()

		in = inboxCh
	}
}
