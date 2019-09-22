/*
 * MIT License
 *
 * Copyright (c) 2019 Aniketh Girish
 * This file is part of libprobe.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package skb

type probeSkbuffHead struct {
	next, prev *ProbeSkbuff
	qlen       uint
}

// ProbeSkbuff is the implementation of sk_buff
type ProbeSkbuff struct {
	next, prev                     *ProbeSkbuff
	head, end, data, tail, payload uint
	protocol                       uint
	len, dlen, seq, endseq         uint
	refcount                       int
}

// ProbeSkbuffInit initialises a sk_buff
func ProbeSkbuffInit(size uint) *ProbeSkbuff {
	return &ProbeSkbuff{
		len: size,
	}
}

// ProbeSkbuffPeek returns nil for an empty list or a pointer to the head element
func ProbeSkbuffPeek(head *probeSkbuffHead) *ProbeSkbuff {
	skb := head.next
	if skb == nil {
		return nil
	}
	return skb
}

func ProbeSkbuffDeinit(skb *ProbeSkbuff) {
	skb = nil
}

func ProbeSkbuffPush(skb *ProbeSkbuff, len uint) uint {
	skb.data -= len
	skb.len += len
	return skb.data
}

func ProbeSkbuffReserve(skb *ProbeSkbuff, len uint) {
	skb.data += len
	skb.tail += len
}

func ProbeSkbuffResetHead(skb *ProbeSkbuff) {
	skb.head = 0
}

func ProbeSkbuffGetHead(skb *ProbeSkbuff) uint {
	return skb.head
}

// ProbeSkbuffUnlink is used to remove a SKB from the list
func ProbeSkbuffUnlink(skb *ProbeSkbuff, list *probeSkbuffHead) {
	var next, prev *ProbeSkbuff

	list.qlen--
	next = skb.next
	prev = skb.prev
	skb.next = nil
	skb.prev = nil
	next.prev = prev
	prev.next = next
}

/*
 * Probe sk buff datagram/head queue function design
 */

// ProbeSkbuffInsert is used to insert ProbeSkbuff into a list
func ProbeSkbuffInsert(newsk *ProbeSkbuff, prev *ProbeSkbuff, next *ProbeSkbuff, list *probeSkbuffHead) {
	newsk.next = next
	newsk.prev = prev
	next.prev = newsk
	prev.next = newsk
	list.qlen++
}

func ProbeSkbuffQueueInit(head *probeSkbuffHead) {
	head.next = nil
	head.prev = nil
	head.qlen = 0
}

func ProbSkbuffQueueDeinit(head *probeSkbuffHead) {
	head.next = nil
	head.prev = nil
	head.qlen = 0
}

// ProbeSkbuffDequeue is used to remove the first SKB from the queue
func ProbeSkbuffDequeue(head *probeSkbuffHead) *ProbeSkbuff {
	skb := ProbeSkbuffPeek(head)
	if skb != nil {
		ProbeSkbuffUnlink(skb, head)
	}
	return skb
}

// ProbeSkbuffQueueAdd is used to add new SKB before next SKB
func ProbeSkbuffQueueAdd(head *probeSkbuffHead, new *ProbeSkbuff, next *ProbeSkbuff) {
	ProbeSkbuffInsert(new, next.prev, next, head)
}

// ProbeSkbuffQueueAddEnd is used to add
func ProbeSkbuffQueueAddEnd(head *probeSkbuffHead, new *ProbeSkbuff) {
	ProbeSkbuffQueueAdd(head, new, head.prev)
}

// ProbeSkbuffQueueCheckEmpty return 1 if the queue is empty, 0 otherwise
func ProbeSkbuffQueueCheckEmpty(head *probeSkbuffHead) int {
	if head.next == nil {
		return 1
	} else {
		return 0
	}
}

// ProbeSkbuffQueueGetLen returns the length of the sk_buff queue
func ProbeSkbuffQueueGetLen(head *probeSkbuffHead) uint {
	return head.qlen
}
