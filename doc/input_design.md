Input Design
============

One design is place holder. But this prevents loop inside the computation. And
it avoids taking multiple results from the same input feed.

One way is using infeed queue design, backed by Go channel.
