# Initialize protobuf types

To use the protobuf types, they have to be linked to the ```/usr/local/include```
folder, in the same way like Google's protobuf well-known-types.

```console
foo@bar:~$ mkdir /usr/local/include/roockie
foo@bar:~$ ln -s $GOPATH/src/github.com/roockie-nebula/protobuf/types /usr/local/include/roockie/protobuf
```

Now the protobuf types can be imported into your proto definition file.
```
syntax = "proto3";

package startrek;

import "roockie-nebula/protobuf/timestamp.proto";
import "roockie-nebula/protobuf/nullstring.proto";
import "roockie-nebula/protobuf/nullint.proto";

message StarfleetShip {
	string name = 1;
	roockie.protobuf.NullInt passengers = 2;
	roockie.protobuf.NullString mission = 3;
	roockie.protobuf.Timestamp departure_time = 4;
}
```
