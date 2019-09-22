# sk_buff

The socket buffer, or "SKB", is the most fundamental data structure in the Linux networking code. Every packet sent or received is handled using this data structure.

![sk_buff representation](http://www.embeddedlinux.org.cn/linux_net/0596002556/images/understandlni_0201.jpg)

The above image is a simple representation of the `sk_buff` data structure coupled with the `sk_buff_head` structure. The SKBs are stored on a list whose head is implemented by `struct sk_buff_head`. Next we have queue management functions that which are done using the `sk_buff_head` which has pointers pointing to the first and last `sk_buff` in the list.

## Resources
* Read this for an overview on SKB. [https://wiki.linuxfoundation.org/networking/sk_buff](https://wiki.linuxfoundation.org/networking/sk_buff)

* Read this to understand how the SKB is a doubly linked list and a queue implemented together. [http://vger.kernel.org/~davem/skb_list.html](http://vger.kernel.org/~davem/skb_list.html)
* Read this to get an overview on the various fields in `struct sk_buff`. [http://vger.kernel.org/~davem/skb.html](http://vger.kernel.org/~davem/skb.html)