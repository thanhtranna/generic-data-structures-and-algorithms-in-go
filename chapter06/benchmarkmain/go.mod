module main

go 1.19

replace chapter06/nodequeue => ../nodequeue

replace chapter06/slicequeue => ../slicequeue

require chapter06/nodequeue v0.0.0-00010101000000-000000000000 // indirect

require chapter06/slicequeue v0.0.0-00010101000000-000000000000 // indirect