package currency

import (
	"encoding/json"
	base3 "github.com/ProtoconNet/mitum-currency/v2/base"
	"github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util"
	jsonenc "github.com/ProtoconNet/mitum2/util/encoder/json"
)

type CurrencyPolicyUpdaterFactJSONMarshaler struct {
	base.BaseFactJSONMarshaler
	Currency base3.CurrencyID     `json:"currency"`
	Policy   base3.CurrencyPolicy `json:"policy"`
}

func (fact CurrencyPolicyUpdaterFact) MarshalJSON() ([]byte, error) {
	return util.MarshalJSON(CurrencyPolicyUpdaterFactJSONMarshaler{
		BaseFactJSONMarshaler: fact.BaseFact.JSONMarshaler(),
		Currency:              fact.currency,
		Policy:                fact.policy,
	})
}

type CurrencyPolicyUpdaterFactJSONUnMarshaler struct {
	base.BaseFactJSONUnmarshaler
	Currency string          `json:"currency"`
	Policy   json.RawMessage `json:"policy"`
}

func (fact *CurrencyPolicyUpdaterFact) DecodeJSON(b []byte, enc *jsonenc.Encoder) error {
	e := util.StringErrorFunc("failed to decode json of CurrencyPolicyUpdaterFact")

	var uf CurrencyPolicyUpdaterFactJSONUnMarshaler
	if err := enc.Unmarshal(b, &uf); err != nil {
		return e(err, "")
	}

	fact.BaseFact.SetJSONUnmarshaler(uf.BaseFactJSONUnmarshaler)

	return fact.unpack(enc, uf.Currency, uf.Policy)
}

type currencyPolicyUpdaterMarshaler struct {
	base3.BaseOperationJSONMarshaler
}

func (op CurrencyPolicyUpdater) MarshalJSON() ([]byte, error) {
	return util.MarshalJSON(currencyPolicyUpdaterMarshaler{
		BaseOperationJSONMarshaler: op.BaseOperation.JSONMarshaler(),
	})
}

func (op *CurrencyPolicyUpdater) DecodeJSON(b []byte, enc *jsonenc.Encoder) error {
	e := util.StringErrorFunc("failed to decode CurrencyPolicyUpdater")

	var ubo base3.BaseNodeOperation
	if err := ubo.DecodeJSON(b, enc); err != nil {
		return e(err, "")
	}

	op.BaseNodeOperation = ubo

	return nil
}