package types

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/types"
)

const (
	PropertyIDPrefix = "prop"
)

type ID interface {
	String() string
	Uint64() uint64
	Bytes() []byte
	Prefix() string
	IsEqual(ID) bool
	MarshalJSON() ([]byte, error)
}

var (
	_ ID = PropertyID{}
)

type PropertyID []byte

func NewPropertyID(i uint64) PropertyID {
	return types.Uint64ToBigEndian(i)
}

func NewPropertyIDFromString(s string) (PropertyID, error) {
	if len(s) < 5 {
		return nil, fmt.Errorf("invalid property id length")
	}

	i, err := strconv.ParseUint(s[4:], 16, 64)
	if err != nil {
		return nil, err
	}

	return NewPropertyID(i), nil
}

func (id PropertyID) String() string {
	return fmt.Sprintf("%s%x", PropertyIDPrefix, id.Uint64())
}

func (id PropertyID) Uint64() uint64 {
	return binary.BigEndian.Uint64(id)
}

func (id PropertyID) Bytes() []byte {
	return id
}

func (id PropertyID) Prefix() string {
	return PropertyIDPrefix
}

func (id PropertyID) IsEqual(_id ID) bool {
	return id.String() == _id.String()
}

func (id PropertyID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

func (id *PropertyID) UnmarshalJSON(bytes []byte) error {
	var s string
	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	_id, err := NewPropertyIDFromString(s)
	if err != nil {
		return err
	}

	*id = _id

	return nil
}

var _ sort.Interface = (*IDs)(nil)

type IDs []ID

func (ids IDs) Append(id ID) IDs {
	return append(ids, id)
}

func (ids IDs) Len() int {
	return len(ids)
}

func (ids IDs) Less(x, y int) bool {
	i := strings.Compare(ids[x].Prefix(), ids[y].Prefix())
	if i < 0 {
		return true
	} else if i == 0 {
		return ids[x].Uint64() < ids[y].Uint64()
	}

	return false
}

func (ids IDs) Swap(x, y int) {
	ids[x], ids[y] = ids[y], ids[x]
}

func (ids IDs) Sort() IDs {
	sort.Slice(ids, ids.Less)
	return ids
}

func (ids IDs) Delete(x int) IDs {
	ids[x] = ids[ids.Len()-1]
	return ids[:ids.Len()-1]
}

func (ids IDs) Search(id ID) int {
	i := id.Uint64()
	index := sort.Search(len(ids), func(x int) bool {
		return ids[x].Prefix() > id.Prefix() || ids[x].Uint64() >= i
	})

	if (index == ids.Len()) ||
		(index < ids.Len() && ids[index].String() != id.String()) {
		return ids.Len()
	}

	return index
}
