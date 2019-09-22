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

// ProbeListAdd take the value for the new node and adds that to the linklist chain.
func ProbeListAdd(newsk *ProbeSkbuff, prev *ProbeSkbuff, next *ProbeSkbuff, list *probeSkbuffHead) {
	newsk.next = next
	newsk.prev = prev
	next.prev = newsk
	prev.next = newsk
	list.qlen++
}

// ProbeListDelete delete's a probe list node according to the node context passed.
func ProbeListDelete(skb *ProbeSkbuff, list *probeSkbuffHead) {
	var next, prev *ProbeSkbuff

	list.qlen--
	next = skb.next
	prev = skb.prev
	skb.next = nil
	skb.prev = nil
	next.prev = prev
	prev.next = next
}

// ProbeListCheckEmpty checks whether the list is empty. Returns True if the list is empty.
func ProbeListCheckEmpty(head *probeSkbuffHead) int {
	if head.next == nil {
		return 1
	}
	return 0
}
