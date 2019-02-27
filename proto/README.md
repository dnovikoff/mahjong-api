## General conventions


### Instance
All `int64` fields, named `instance` or ending with suffix `_instance` are tiles instances.

Instance values starts from value of `1` and ends with value of
`136`. Value `0` used for `undefined` (not set).

Four sequenced numbers indicates same tile. Ex. 1,2,3,4 is for 1 Man, 5,6,7,8 is for 2 Man, etc.

Tiles order is the following:

- `123456789` `Man` (numbers 1-36)
- `123456789` `Pin` (numbers 37-72)
- `123456789` `Sou` (numbers 73-108)
- `East`, `South`, `West`, `East` (numbers 109-124)
- `White`, `Green`, `Red` (numbers 125-136). Take note, that this values starts from `White` (not `Red`, as stated in some manuals)

### Mask
All `int64` fields, ending with suffix `_mask` are tile bit-masks.

Each bit represents tile type (not instance). This is used Ex. to show which tiles are allowed to be droped.
The client should match the correcponding instances to the bits in mask. The value of `0` indicates no tiles. The value of `2^34-1` indicates all tiles.

The bit values are the fololowing:

- `123456789` `Man` (bits 0-8)
- `123456789` `Pin` (bits 9-17)
- `123456789` `Sou` (bits 18-26)
- `East`, `South`, `West`, `East` (bits 27-30)
- `White`, `Green`, `Red` (bits 31-33)

# Player index

Player indexes are `int64` fields ends with suffix `_index`.
Ex.: `who_index`, `from_index`, `dealer_index`, `client_index`.

Possible values are `1`-`4`.
They indicate indexes in arrays like `hands` or `changes` (starting from `1`).
Zero `0` used to indicate empty value.