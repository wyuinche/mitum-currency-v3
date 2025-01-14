package types // nolint: dupl, revive

import (
	bsonenc "github.com/ProtoconNet/mitum-currency/v3/digest/util/bson"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/hint"
	"go.mongodb.org/mongo-driver/bson"
)

func (cs ContractAccountStatus) MarshalBSON() ([]byte, error) {
	return bsonenc.Marshal(
		bson.M{
			"_hint":     cs.Hint().String(),
			"is_active": cs.isActive,
			"owner":     cs.owner,
		},
	)
}

type ContractAccountBSONUnmarshaler struct {
	Hint     string `bson:"_hint"`
	IsActive bool   `bson:"is_active"`
	Owner    string `bson:"owner"`
}

func (cs *ContractAccountStatus) DecodeBSON(b []byte, enc *bsonenc.Encoder) error {
	e := util.StringError("decode bson of ContractAccountStatus")

	var ucs ContractAccountBSONUnmarshaler
	if err := bsonenc.Unmarshal(b, &ucs); err != nil {
		return e.Wrap(err)
	}

	ht, err := hint.ParseHint(ucs.Hint)
	if err != nil {
		return e.Wrap(err)
	}

	return cs.unpack(enc, ht, ucs.IsActive, ucs.Owner)
}
